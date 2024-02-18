package httpserver

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "run server",
			Run: func(_ *cobra.Command, _ []string) {
				fmt.Println("Hello server")
			},
		},
	)
}
