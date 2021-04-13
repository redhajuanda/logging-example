package main

import (
	"fmt"
	"logging-example/infra/logger"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	log := logger.New()
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Error("error without trace id")

		ctx := r.Context()
		log.With(ctx).Error("error with context 1")
		log.With(ctx).Info("info")
		log.With(ctx).Error("error with context 2")
		log.With(ctx).Debug("debug with context")
		log.With(ctx).WithProcess(func() { fmt.Println("Run process") }).Error("error with process")

		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
