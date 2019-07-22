package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/pnutmath/gappa-tappa/backend/internal/common"
	"github.com/pnutmath/gappa-tappa/backend/internal/config"
)

func main() {
	fmt.Println("Gappa Tappa v0.1: Started")
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = config.DefaultHost
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := common.Connect(ctx, host)
	if err != nil {
		fmt.Printf("Unable to connect DB %v", err)
	}
	db := client.Database("gappatappa")

	router := mux.NewRouter()

	common.SetupRoute(router, db)

	http.ListenAndServe(":8080", router)
}
