package httpserver

import (
	"github.com/j3yzz/quickbite/internal/config"
	"github.com/j3yzz/quickbite/internal/delivery/httpserver"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "run server",
			Run: func(_ *cobra.Command, _ []string) {
				serverConfig := config.Server{Port: "8080"}

				s := httpserver.New(serverConfig)
				s.Provide()
			},
		},
	)
}
