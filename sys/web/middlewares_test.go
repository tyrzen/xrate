package web_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GenesisEducationKyiv/main-project-delveper/sys/logger"
	"github.com/GenesisEducationKyiv/main-project-delveper/sys/web"
	"github.com/stretchr/testify/require"
)

func TestChainMiddlewares(t *testing.T) {
	key := &struct{}{}

	mwFn := func(n int) web.Middleware {
		return func(next web.Handler) web.Handler {
			return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
				val := fmt.Sprintf("%v%v", ctx.Value(&struct{}{}), n)
				ctx = context.WithValue(ctx, key, val)

				return next(ctx, rw, req)
			}
		}
	}

	mws := []web.Middleware{
		mwFn(1),
		mwFn(2),
		mwFn(3),
		mwFn(4),
	}

	handler := func(ctx context.Context, rw http.ResponseWriter, _ *http.Request) error {
		rw.Write([]byte(fmt.Sprint(ctx.Value(key))))
		return nil
	}

	chainedHandler := web.ChainMiddlewares(handler, mws...)

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	err := chainedHandler(context.WithValue(context.Background(), key, ""), rw, req)
	require.NoError(t, err)

	require.Equal(t, "1234", rw.Body.String())
}

func TestMiddlewares(t *testing.T) {
	log := logger.New(logger.LevelDebug)
	defer log.Sync()

	t.Run("WithJSON", func(t *testing.T) {
		rw := httptest.NewRecorder()
		h := func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			rw.Write([]byte("test"))
			return nil
		}
		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)

		mw := web.WithJSON(h)
		err := mw(context.Background(), rw, req)

		require.NoError(t, err)
		require.Equal(t, "application/json; charset=UTF-8", rw.Header().Get("Content-Type"))
	})

	t.Run("WithLogRequest", func(t *testing.T) {
		rw := httptest.NewRecorder()
		h := func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			rw.Write([]byte("test"))
			return nil
		}
		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)

		mw := web.WithLogRequest(log)(h)
		err := mw(context.Background(), rw, req)

		require.NoError(t, err)
		require.Equal(t, "test", rw.Body.String())
	})

	t.Run("WithRecover", func(t *testing.T) {
		rw := httptest.NewRecorder()
		h := func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			panic("test")
		}
		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)

		mw := web.WithRecover(log)(h)
		err := mw(context.Background(), rw, req)

		require.NoError(t, err)
	})

	t.Run("WithCORS", func(t *testing.T) {
		rw := httptest.NewRecorder()
		h := func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			rw.Write([]byte(`{"test":"test"}`))
			return nil
		}
		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)

		mw := web.WithCORS("*")(h)
		err := mw(context.Background(), rw, req)

		require.NoError(t, err)
		require.Equal(t, `{"test":"test"}`, rw.Body.String())
	})

	t.Run("WithErrors", func(t *testing.T) {
		rw := httptest.NewRecorder()
		h := func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			return errors.New("test error")
		}
		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)

		mw := web.WithErrors(log)(h)
		err := mw(context.Background(), rw, req)

		require.NoError(t, err)
		require.JSONEq(t, `{"error":"Internal Server Error"}`, rw.Body.String())
	})
}
