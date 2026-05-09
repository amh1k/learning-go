package gracefulshutdown

import (
	"context"
	"net/http"
	"os"
	"time"
)

const (
	k8sDefaultTerminationGracePeriod = 30 * time.Second
)

type (
	HTTPServer interface {
		ListenAndServe() error
		Shutdown(ctx context.Context) error
	}
	Server struct {
		shutdown <- chan os.Signal
		delegate HTTPServer
		timeout time.Duration

	}
	ServerOption func(server *Server)
)

func WithShutdownSignal(shutdownSignal <- chan os.Signal)ServerOption{
	return func(server *Server) {
		server.shutdown = shutdownSignal
	}
}
func WithTimeout(timeout time.Duration) ServerOption {
	return func(server *Server) {
		server.timeout = timeout
	}
}

func NewServer(server HTTPServer, options ...ServerOption) * Server {
	s := &Server{
		delegate: server,
		timeout:  k8sDefaultTerminationGracePeriod,
		shutdown: newInterruptSignalChannel(),
	}

	for _, option := range options {
		option(s)
	}

	return s
}
func (s *Server) ListenAndServe(ctx context.Context) error {
	select {
	case err := <-s.delegateListenAndServe():
		return err
	case <-ctx.Done():
		return s.shutdownDelegate(ctx)
	case <-s.shutdown:
		return s.shutdownDelegate(ctx)
	}
}
func (s *Server) delegateListenAndServe() chan error {
	listenErr := make(chan error)
	go func() {
		if err := s.delegate.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			listenErr <- err
		}
	}()
	return listenErr
}

func (s *Server) shutdownDelegate(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	if err := s.delegate.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		return err
	}
	return ctx.Err()
}


