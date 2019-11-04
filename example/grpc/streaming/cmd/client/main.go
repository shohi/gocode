package main

import (
	"log"

	"github.com/shohi/gocode/example/grpc/streaming/client"
	"github.com/spf13/cobra"
)

var (
	conf client.Config

	rootCmd = &cobra.Command{
		Use:   "client",
		Short: "rpc client",
		RunE:  runRoot,
	}
)

func main() {
	Execute()
}

// Execute is entrance
func Execute() {
	setupFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}

// runRoot implements the main command with no subcommands
func runRoot(_cmd *cobra.Command, _args []string) error {
	client.Run(conf)
	return nil
}

// setupFlags sets flags for comand line
func setupFlags(cmd *cobra.Command) {
	fs := cmd.Flags()

	fs.StringVar(&conf.Addr, "addr", ":9002", "rpc server address")
	fs.StringVar(&conf.Service, "service", "",
		`rpc services, separated by comma, e.g. "list,route,record`)
}
