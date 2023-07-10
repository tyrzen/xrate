// Package subscription provides functionality to manage subscriptions.
package subs

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/GenesisEducationKyiv/main-project-delveper/internal/event"
)

const defaultTimeout = 5 * time.Second

const (
	currencyBTC = "BTC"
	currencyUAH = "UAH"
)

var (
	// ErrEmailAlreadyExists is an error indicating that the email address already exists in the database.
	ErrEmailAlreadyExists = errors.New("email address exists")

	// ErrMissingEmail is an error indicating that the email address is missing.
	ErrMissingEmail = errors.New("missing email")
)

// Subscriber represents an entity that subscribes to emails.
type Subscriber struct {
	Address *mail.Address
	Topic   Topic
}

func NewSubscriber(address *mail.Address, topic Topic) *Subscriber {
	return &Subscriber{Address: address, Topic: topic}
}

// Topic represents a topic that subscribes to emails.
type Topic struct {
	BaseCurrency  string
	QuoteCurrency string
}

func NewTopic(base, quote string) Topic {
	return Topic{BaseCurrency: base, QuoteCurrency: quote}
}

// Message represents an email message.
type Message struct {
	From    *mail.Address
	To      *mail.Address
	Subject string
	Body    string
}

func NewMessage(subject, body string, to *mail.Address) Message {
	return Message{
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

//go:generate moq -out=../../test/mock/email_repository.go -pkg=mock . SubscriberRepository

// SubscriberRepository is an interface for managing email subscriptions.
type SubscriberRepository interface {
	Add(context.Context, Subscriber) error
	List(context.Context) ([]Subscriber, error)
}

//go:generate moq -out=../../test/mock/email_sender.go -pkg=mock . EmailSender

// EmailSender is an interface for sending emails.
type EmailSender interface {
	Send(Message) error
}

// Service represents a service that manages email subscriptions and sends emails.
type Service struct {
	repo SubscriberRepository
	mail EmailSender
	bus  *event.Bus
}

// NewService creates a new Service instance with the provided dependencies.
func NewService(repo SubscriberRepository, mail EmailSender) *Service {
	return &Service{
		repo: repo,
		mail: mail,
	}
}

// Subscribe adds a new email subscription to the repository.
func (svc *Service) Subscribe(ctx context.Context, sub Subscriber) error {
	// TODO: Fetch a topic from query params in further iteration.
	sub.Topic = NewTopic(currencyBTC, currencyUAH)
	if err := svc.repo.Add(ctx, sub); err != nil {
		return fmt.Errorf("adding subscription: %w", err)
	}

	return nil
}

// SendEmails sends emails to all subscribers using the current rate.
func (svc *Service) SendEmails(ctx context.Context) error {
	// TODO: Improve the logic of retrieving exchange rate from event bus.
	var val float64

	subscribers, err := svc.repo.List(ctx)
	if err != nil {
		return err
	}

	var errArr []error

	for _, sub := range subscribers {
		// TODO: Improve building message body.
		subject := fmt.Sprintf("%s exchange rate at %s", sub.Topic, time.Now().Format(time.Stamp))
		body := fmt.Sprintf("Current exhange rate: %f", val)

		msg := NewMessage(
			subject,
			body,
			sub.Address,
		)

		if err := svc.mail.Send(msg); err != nil {
			errArr = append(errArr, err)
		}
	}

	if errArr != nil {
		return errors.Join(errArr...)
	}

	return nil
}