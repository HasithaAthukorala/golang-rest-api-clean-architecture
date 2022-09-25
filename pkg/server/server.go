package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"golang-rest-api-clean-architecture/pkg/api"
	services "golang-rest-api-clean-architecture/pkg/external-services"
	"net/http"
)

type Options struct {
	EnableAuth bool
	Port       int
}

type RouterFunc func(*mux.Router)

type Server struct {
	opt    *Options
	logger *zap.SugaredLogger
	router *mux.Router
}

type Routes struct {
	Prefix string
	Routes []api.Route
}

func HandleWithClientSet(clientSet services.ClientSet, handle api.HandlerFunc) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		handle(clientSet, w, req)
	}
}

func New(clientSet services.ClientSet, routes Routes, opt Options, logger *zap.SugaredLogger) *Server {
	logger = logger.Named("http-server")
	r := mux.NewRouter().StrictSlash(true)
	r.Use(ResponseHeadersMiddleware(map[string]string{
		"Content-Type": "application/json",
	}))
	secured := r.Name("secured").Subrouter()
	unsecured := r.Name("unsecured").Subrouter()

	if opt.EnableAuth {
		//secured.Use() TODO: handle authentication here
	}

	addRouteFunc := func(r *mux.Router, route api.Route, prefix string) {
		newRoute := r.NewRoute()
		newRoute.PathPrefix(prefix)
		if len(route.Path) > 0 {
			newRoute.Path(route.Path)
		}
		if len(route.Method) > 0 {
			newRoute.Methods(route.Method)
		}
		newRoute.HandlerFunc(HandleWithClientSet(clientSet, route.HandlerFunc))
	}
	for _, route := range routes.Routes {
		if route.Insecure {
			addRouteFunc(unsecured, route, routes.Prefix)
		} else {
			addRouteFunc(secured, route, routes.Prefix)
		}
	}
	srv := &Server{
		opt:    &opt,
		logger: logger,
		router: r,
	}
	return srv
}

func (s *Server) Run(stopCh <-chan struct{}) error {
	addr := fmt.Sprintf(":%d", s.opt.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: s.router,
	}
	go func() {
		s.logger.Infof("serving on %s", addr)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.Errorf("listenAndServe error: %v", err)
		}
	}()
	<-stopCh
	s.logger.Info("shutting down server...")
	return srv.Shutdown(context.Background())
}
