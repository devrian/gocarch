package cmd

import (
	"gocarch/cmd/server"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	root = &cobra.Command{
		Use:   "serve-http",
		Short: "HTTP service",
		Long:  "Serve through HTTP",
	}
)

func Execute() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// You could register your command here
	root.AddCommand(server.ServeHTTP())

	if err := root.Execute(); err != nil {
		log.Fatalln("Error: \n", err.Error())
		os.Exit(-1)
	}
}
