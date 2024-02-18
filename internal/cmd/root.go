package cmd

import (
	"github.com/j3yzz/quickbite/internal/cmd/httpserver"
	"github.com/spf13/cobra"
	"log"
)

func Execute() {
	cmd := &cobra.Command{
		Use:   "QuickBite",
		Short: "Fast and Fresh Food Delivery",
	}

	httpserver.Register(cmd)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
