package main

import (
	"api-service/internal/config"
	"api-service/internal/database"
	"api-service/internal/repository"
	"api-service/internal/services"
	"api-service/internal/worker"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.InitConfig()

	db := database.InitDatabase(cfg)
	repo := repository.NewSiteRepositoryGORM(db)
	checkService := services.NewCheckerService(repo)
	consumer, err := worker.NewConsumer(cfg.MqUrl, checkService)
	if err != nil {
		log.Fatalf("failed to create consumer: %v", err)
	}

	concurrency := 5
	log.Printf("Starting consumer with concurrency: %d", concurrency)
	consumer.Start(concurrency)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down worker...")
	consumer.Shutdown()
	log.Println("Worker shut down gracefully")
}
