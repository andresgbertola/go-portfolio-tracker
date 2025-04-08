package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"go-portfolio-tracker/internal/infrastructure/repository"
	"go-portfolio-tracker/internal/interface/controller"
	"go-portfolio-tracker/internal/usecase/command"
	"go-portfolio-tracker/internal/usecase/query"
)

func main() {
	repo := &repository.AssetRepositoryMock{}

	// Use cases
	getAllAssetsQuery := query.NewGetAllAssetsQuery(repo)
	createNewAssetCommand := command.NewCreateNewAssetCommand(repo)

	// HTTP Controller
	handler := controller.NewAssetController(getAllAssetsQuery, createNewAssetCommand)

	// Router setup
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Routes
	r.Get("/api/assets", handler.GetAllAssets)
	r.Post("/api/assets", handler.CreateNewAsset)

	// Start server
	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
