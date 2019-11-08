package main

import (
	"context"
	"github.com/yetialex/autoteka/internal/db"
	"github.com/yetialex/autoteka/internal/metadata"
	"github.com/yetialex/autoteka/internal/web"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	serviceName = "autoteka"
)

func main() {
	//time.Sleep(10 * time.Second)
	ctx := context.Background()
	conn := db.Connect()
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	if conn == nil {
		log.Println("no db connection, exiting")
		return
	}
	ctx = metadata.SetContextValues(ctx, serviceName, conn)

	srv := web.Start(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	web.GracefulShutdown(srv)
}
