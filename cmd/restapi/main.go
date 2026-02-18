package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pelletier/go-toml/v2"

	"github.com/ixmael/99minutos/internal/adapters/handlers/shipment_handlers"
	"github.com/ixmael/99minutos/internal/core/services/shipmentservice"
	"github.com/ixmael/99minutos/internal/infrastructure/postgres/shipmentrepository"
	"github.com/ixmael/99minutos/internal/infrastructure/postgres/shipmentstatusrepository"
	"github.com/ixmael/99minutos/internal/infrastructure/zaplogger"
)

// ApplicationConfig represents the configuration for the application.
type ApplicationConfig struct {
	Environment string `toml:"environment"`
	RestAPI     struct {
		Port int `toml:"port"`
	} `toml:"restapi"`
	Repository struct {
		PostgresURL string `toml:"postgres"`
	} `toml:"repository"`
}

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

	zaplogger, err := zaplogger.NewZapLogger(cnfg.Environment)
	if err != nil {
		log.Println("error on initializing the logger")
		os.Exit(1)
		return
	}

	shipmentpostgresrepository, err := shipmentrepository.NewPostgresShipmentRepository(cnfg.Repository.PostgresURL)
	if err != nil {
		log.Println("error on initializing the shipment repository")
		os.Exit(1)
		return
	}

	shipmentstatuspostgresrepository, err := shipmentstatusrepository.NewPostgresShipmentStatusRepository(cnfg.Repository.PostgresURL)
	if err != nil {
		log.Println("error on initializing the shipment repository")
		os.Exit(1)
		return
	}

	shipmentService, err := shipmentservice.NewShipmentService(zaplogger, shipmentpostgresrepository, shipmentstatuspostgresrepository)
	if err != nil {
		log.Println("error on setup the shipment service")
		os.Exit(1)
		return
	}

	h := shipment_handlers.NewHTTPShipmentHandler(shipmentService)

	http.HandleFunc("POST /shipments", h.Register)
	http.HandleFunc("GET /shipments/{shipment_id}", h.ListShipmentDetails)

	port := fmt.Sprintf(":%d", cnfg.RestAPI.Port)
	log.Fatal(http.ListenAndServe(port, nil))
}
