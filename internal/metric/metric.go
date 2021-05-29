// Package metric contains http server for prometheus /metrics handler
package metric

import (
	"context"
	"net/http"
	"time"

	"github.com/kaatinga/plantbook/pkg/logging"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	writeTimeout time.Duration = 15 * time.Second
	readTimeout  time.Duration = 15 * time.Second
	idleTimeout  time.Duration = 15 * time.Second
	stopTimeout  time.Duration = 15 * time.Second
)

type Server struct {
	addr string
	srv  *http.Server
}

func NewServer(address string) *Server {
	r := mux.NewRouter().StrictSlash(true)
	r.Path("/metrics").Handler(promhttp.Handler())
	srv := &http.Server{
		Addr: address,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: writeTimeout,
		ReadTimeout:  writeTimeout,
		IdleTimeout:  writeTimeout,
		Handler:      r,
	}
	return &Server{addr: address, srv: srv}
}

func (s *Server) Run(ctx context.Context) {
	log := logging.FromContext(ctx)
	go func() {
		log.Infof("Starting listening for metrics on %s", s.addr)
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	<-ctx.Done()
	log.Infof("got signal to shutdown, timeout %v", stopTimeout)
	ctxstop, cancelstop := context.WithTimeout(context.Background(), stopTimeout)
	defer cancelstop()
	err := s.srv.Shutdown(ctxstop)
	if err != nil {
		log.Errorf("srv.Shutdown error, %s", err)
	}
}
