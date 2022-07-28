package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KenethSandoval/doc-md/internal/config"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/app"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/mongodb"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	cfg := config.New()

	mongoDB := mongodb.New(cfg)

	if err := mongoDB.Connect(); err != nil {
		log.Fatal("Error connect mongo db")
	}

	e := app.New(cfg, mongoDB)

	go func() {
		if err := e.Start(":" + cfg.Port); err != nil && err != http.ErrServerClosed {
			log.Fatal(err.Error())
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	defer mongoDB.Close()

	fmt.Println("💥 shutdown server ...")

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
