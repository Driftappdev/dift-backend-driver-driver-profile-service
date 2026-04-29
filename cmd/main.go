package main

import (
	"fmt"
	"net/http"

	"dift_backend_driver/driver-profile-service/config"
	"dift_backend_driver/driver-profile-service/internal/handler"
	"dift_backend_driver/driver-profile-service/internal/route"
	"dift_backend_driver/driver-profile-service/internal/service"
	"github.com/driftappdev/libpackage/gologger"
)

func main() {
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		panic(err)
	}
	logger := gologger.Default()
	svc := service.NewProfileService()
	h := handler.NewProfileHandler(svc)
	mux := http.NewServeMux()
	route.Register(mux, h)
	addr := fmt.Sprintf(":%d", cfg.Server.HTTPPort)
	logger.Info("driver-profile-service started", gologger.F("port", cfg.Server.HTTPPort))
	if err := http.ListenAndServe(addr, mux); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
