package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ixmael/99minutos/internal/adapters/workers"
	"github.com/pelletier/go-toml/v2"
)

func main() {
	configFilePathFlag := flag.String("config", ".env.toml", "The configuration file path")

	flag.Parse()

	configData, err := os.ReadFile(*configFilePathFlag)
	if err != nil {
		log.Println("error on reading the config file")
		os.Exit(1)
		return
	}

	var cnfg ApplicationConfig
	err = toml.Unmarshal(configData, &cnfg)
	if err != nil {
		log.Println("error on parsing the config file")
		os.Exit(1)
		return
	}

	services, err := SetupServices(&cnfg)
	if err != nil {
		log.Println("error on initializing the services")
		os.Exit(1)
		return
	}
	defer services.Stop()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// var wg sync.WaitGroup
	//
	worker, err := workers.NewShipmentEventWorker(services.Logger, services.ShipmentService)
	if err != nil {
		log.Println("error on initializing the worker")
		os.Exit(1)
		return
	}

	api := SetupAPI(&cnfg, services)

	go func() {
		services.Logger.Info(fmt.Sprintf("API is running on :%d", cnfg.RestAPI.Port))
		if err := api.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
			os.Exit(1)
		}
	}()

	go func() {
		services.Logger.Info("Worker started")
		worker.Start(ctx)
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := api.Shutdown(shutdownCtx); err != nil {
		log.Printf("Shutdown error: %v", err)
		os.Exit(0)
	}

	os.Exit(0)

	/*
		// Start the RestAPI
		wg.Add(1)
		go func() {
			defer wg.Done()
			services.Logger.Info(fmt.Sprintf("API is running on %s", port))
			if err := api.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Printf("Server error: %v", err)
				os.Exit(1)
			}
		}()

		<-ctx.Done()

		// sigChan := make(chan os.Signal, 1)
		// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		// <-sigChan
		fmt.Println("Shutdown signal received")
		log.Println("Shutdown signal received")
		zaplogger.Info("Shutdown signal received")

		// Graceful shutdown with timeout
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := api.Shutdown(shutdownCtx); err != nil {
			log.Printf("Shutdown error: %v", err)
			os.Exit(0)
		}

		wg.Wait()

		log.Println("Server stopped")
		os.Exit(0)
	*/
}
