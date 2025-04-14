package handler

import (
	"gocarch/config"

	"github.com/devrian/golb/httputil"
	"github.com/devrian/golb/log"
	"github.com/devrian/golb/router"
)

type (
	Options struct {
		Config config.MainConfig
	}

	Handler struct {
		options     *Options
		listenErrCh chan error
		router      *router.HttpRouter
	}
)

func New(o *Options) *Handler {
	handler := &Handler{options: o}
	return handler
}

func (h *Handler) Run() {
	log.Infof("API Listening on %s", h.options.Config.Server.Port)
	h.listenErrCh <- httputil.Serve(
		h.options.Config.Server.Port,
		h.router,
		h.options.Config.Server.GracefulTimeout,
		h.options.Config.Server.ReadTimeout,
		h.options.Config.Server.WriteTimeout,
	)
}

func (h *Handler) ListenError() <-chan error {
	return h.listenErrCh
}
