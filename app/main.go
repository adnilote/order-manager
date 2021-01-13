package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/adnilote/order-manager/app/api"
	"github.com/adnilote/order-manager/app/config"
	"github.com/adnilote/order-manager/app/logger"
	"github.com/adnilote/order-manager/app/store"
	"github.com/adnilote/order-manager/app/usecases"
	"github.com/sirupsen/logrus"
)

var (
	configFileName = flag.String("config", "config.yaml", "Config file name")
	logLevel       = flag.String("l", "debug", "Minimum log level")
)

func main() {
	flag.Parse()

	err := config.LoadConfig(*configFileName)
	if err != nil {
		log.Fatal("Load config failed: ", err)
	}

	err = logger.Init(*logLevel, config.Config.SentryDSN)
	if err != nil {
		log.Fatal("Init logger failed: ", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	store, err := store.NewStore(ctx)
	if err != nil {
		log.Fatal("init store failed: ", err)
	}

	us, err := usecases.Init(ctx, store)
	if err != nil {
		log.Fatal("init usecases failed: ", err)
	}

	service := api.NewService(ctx, us)
	err = service.Start()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		cancel()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	defer func() {
		logrus.Info("Gracefull shutdown")
		cancel()
		service.Close()
		us.Close()
	}()

	select {
	case <-ctx.Done():
		break
	case <-quit:
		logrus.Info("Gracefull shutdown1")
		break
	}

}
