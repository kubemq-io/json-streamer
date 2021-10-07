package root

import (
	"context"
	"github/kubemq-io/json-streamer/config"
	"github/kubemq-io/json-streamer/services"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var cfg *config.Config
var rootCmd = &cobra.Command{
	Use: "player",
}

func Execute(args []string) error {
	cfg := config.NewConfig()
	rootCmd.PersistentFlags().StringVarP(&cfg.Address, "address", "a", "kubemq-cluster-grpc.kubemq:50000", "kubemq server address")
	rootCmd.PersistentFlags().StringVarP(&cfg.Queue, "queue", "q", "songs", "kubemq queue name")
	rootCmd.PersistentFlags().StringVarP(&cfg.Table, "table", "t", "songs", "sql table name to insert")
	rootCmd.PersistentFlags().IntVarP(&cfg.Interval, "interval", "i", 1, "new song play interval")
	_ = rootCmd.PersistentFlags().Parse(args)
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	if err := cfg.Validate(); err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	service := services.NewService()
	if err := service.Init(ctx, cfg); err != nil {
		return err
	}
	if err := service.Start(ctx); err != nil {
		return err
	}
	var gracefulShutdown = make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM)
	signal.Notify(gracefulShutdown, syscall.SIGINT)
	signal.Notify(gracefulShutdown, syscall.SIGQUIT)

	<-gracefulShutdown
	return nil
}
