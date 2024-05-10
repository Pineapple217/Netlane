package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/Pineapple217/Netlane/pkg/config"
	"github.com/Pineapple217/Netlane/pkg/database"
	"github.com/Pineapple217/Netlane/pkg/handler"
	"github.com/Pineapple217/Netlane/pkg/server"
	"github.com/Pineapple217/Netlane/pkg/static"
	"github.com/Pineapple217/Netlane/pkg/util"
)

const banner = `
███    ██ ███████ ████████ ██       █████  ███    ██ ███████ 
████   ██ ██         ██    ██      ██   ██ ████   ██ ██      
██ ██  ██ █████      ██    ██      ███████ ██ ██  ██ █████   
██  ██ ██ ██         ██    ██      ██   ██ ██  ██ ██ ██      
██   ████ ███████    ██    ███████ ██   ██ ██   ████ ███████ 
> https://github.com/Pineapple217/Netlane
===- v 0.0.0 -==============================================

`

func main() {
	if len(os.Args) < 2 {
		serve()
		return
	}

	switch os.Args[1] {
	case "serve":
		serve()
	case "config":
		configCmd()
	default:
		fmt.Println("Unknown command:", os.Args[2])
		os.Exit(1)
	}

}

func serve() {
	fmt.Print(banner)

	cnf, err := config.Load()
	util.MaybeDie(err, "Failed to load configs")

	rr := static.HashPublicFS()

	dbc := database.NewDatabaseClient(cnf.Database)
	defer dbc.Close()

	h := handler.NewHandler(dbc)

	server := server.NewServer(cnf.Server)
	server.RegisterRoutes(h)
	server.ApplyMiddleware(rr)
	server.Start()
	defer server.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("Received an interrupt signal, exiting...")
}

func configCmd() {
	switch os.Args[2] {
	case "default":
		fmt.Print(config.GetDefault())
	default:
		fmt.Println("Unknown command:", os.Args[2])
		os.Exit(1)
	}
}
