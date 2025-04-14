package server

import (
	"gocarch/config"
	"gocarch/internal/handler"
	"os"
	"os/signal"
	"syscall"

	"github.com/devrian/golb/log"
	"github.com/devrian/golb/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
)

var (
	serveHttp = &cobra.Command{
		Use:   "serve-http",
		Short: "HTTP service",
		Long:  "Serve through HTTP",
		RunE:  run,
	}
)

func ServeHTTP() *cobra.Command {
	serveHttp.Flags().StringP("config", "c", "", "Config Path, both relative or absolute. i.e: /usr/local/bin/config/files")
	return serveHttp
}

func run(cmd *cobra.Command, args []string) error {
	clog, err := cmd.Flags().GetString("config")
	if err != nil {
		log.WithError(err).Fatalf("Failed to fetching config: %v", err)
	}

	cfg := &config.MainConfig{}
	config.ReadModuleConfig(cfg, "main", clog)

	tracer, closer := tracing.InitFromEnv(cfg.Service.Name)
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	container := newContainer(&opts{
		Config: cfg,
	})

	server := handler.New(&handler.Options{
		Config: container.Config,
	})

	go server.Run()

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Infoln("Exiting gracefully...")
	case err := <-server.ListenError():
		log.Errorln("Error starting web server, exiting gracefully:", err)
	}

	return nil
}
