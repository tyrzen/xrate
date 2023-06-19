// Package subscription provides functionality to manage subscriptions.
package subscription

import (
	"errors"
	"fmt"
	"net/mail"
)

// ErrEmailAlreadyExists is an error indicating that the email address already exists in the database.
var ErrEmailAlreadyExists = errors.New("email address is already in the database")

// Email represents an email address.
type Email struct {
	Address *mail.Address
}

// EmailRepository is an interface for managing email subscriptions.
type EmailRepository interface {
	Add(Email) error
	GetAll() ([]Email, error)
}

// RateGetter is an interface for retrieving a rate.
type RateGetter interface {
	Get() (float64, error)
}

// EmailSender is an interface for sending emails.
type EmailSender interface {
	Send(Email, float64) error
}

// Service represents a service that manages email subscriptions and sends emails.
type Service struct {
	repo EmailRepository
	rate RateGetter
	mail EmailSender
}

// NewService creates a new Service instance with the provided dependencies.
func NewService(repo EmailRepository, rate RateGetter, mail EmailSender) *Service {
	return &Service{
		repo: repo,
		rate: rate,
		mail: mail,
	}
}

// Subscribe adds a new email subscription to the repository.
func (svc *Service) Subscribe(email Email) error {
	if err := svc.repo.Add(email); err != nil {
		return fmt.Errorf("adding email subscription: %w", err)
	}

	return nil
}

// SendEmails sends emails to all subscribers using the current rate.
func (svc *Service) SendEmails() error {
	rate, err := svc.rate.Get()
	if err != nil {
		return err
	}

	emails, err := svc.repo.GetAll()
	if err != nil {
		return err
	}

	var errArr []error

	for _, email := range emails {
		if err := svc.mail.Send(email, rate); err != nil {
			errArr = append(errArr, err)
		}
	}

	if errArr != nil {
		return errors.Join(errArr...)
	}

	return nil
}