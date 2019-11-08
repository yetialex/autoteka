package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/yetialex/autoteka/internal/web/handlers"

	"github.com/gorilla/mux"
)

func newRouter(ctx context.Context) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/autos/{id}", handlers.FindAuto(ctx)).Methods("GET")
	r.HandleFunc("/autos", handlers.CreateAuto(ctx)).Methods("POST")
	r.HandleFunc("/autos", handlers.UpdateAuto(ctx)).Methods("PUT")
	r.HandleFunc("/autos/{id}", handlers.DeleteAuto(ctx)).Methods("DELETE")
	return r
}

func Start(ctx context.Context) *http.Server {
	r := newRouter(ctx)
	listenAddr := ":8080"
	envPort := os.Getenv("listen_port")
	if envPort != "" {
		listenAddr = fmt.Sprintf(":%s", envPort)
	}
	log.Println("listening on", listenAddr)
	srv := &http.Server{
		Addr:         listenAddr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	return srv
}

func GracefulShutdown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Println("graceful shutdown error: ", err)
	}
}
