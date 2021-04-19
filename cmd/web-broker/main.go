package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vlslav/web-broker/internal/app/endpoint"

	"github.com/vlslav/web-broker/internal/app/service"
	"github.com/vlslav/web-broker/internal/pkg/repository"
	clientPkg "github.com/vlslav/web-broker/internal/pkg/some-client"
)

func main() {
	log.Print("Starting the app")

	port := flag.String("port", "8000", "Port")
	storageName := flag.String("storage", "storage.json", "data storage")
	shutdownTimeout := flag.Int64("shutdown_timeout", 3, "shutdown timeout")

	client := clientPkg.New()
	//repo := repository.NewFileRepo(*storageName)
	repo, err := repository.New(repository.FileRepoType, *storageName)
	if err != nil {
		log.Fatalf("something went wrong: %v", err)
	}
	svc := service.New(repo, client)

	serv := http.Server{
		Addr:    net.JoinHostPort("", *port),
		Handler: endpoint.RegisterPublicHTTP(svc),
	}
	go func() {
		if err := serv.ListenAndServe(); err != nil {
			log.Fatalf("listen and serve err: %v", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	log.Print("Started app")

	sig := <-interrupt

	log.Printf("Sig: %v, stopping app", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*shutdownTimeout)*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Printf("shutdown err: %v", err)
	}
}
