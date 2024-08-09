package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ImranZahoor/blog-api/internal/controller"
	"github.com/ImranZahoor/blog-api/internal/repository"
	"github.com/ImranZahoor/blog-api/internal/router"
	"github.com/ImranZahoor/blog-api/internal/service"
	"github.com/ImranZahoor/blog-api/pkg/storage"
)

const (
	PORT = 5000
)

func main() {
	// Initalize In Memory Storage
	memoryStorage := storage.NewInMemoryStorage()
	// Initalize Database Storage
	dbStorage, err := storage.NewMySQLStorageInit()

	if err != nil {
		fmt.Println("datbase Initalization error ")
		panic(err)
	}
	// Initalize file storage
	fileStorage, err := storage.NewFileStorage("category.txt")

	if err != nil {
		fmt.Println("file storage Initalization error")
		panic(err)
	}
	// Initalize Repository
	repo := repository.NewRepository(memoryStorage, dbStorage, fileStorage)
	//Initalize Service
	service := service.NewService(repo)
	// Initalize controller/handlers
	controller := controller.NewController(service)
	// setup routes
	server := router.NewServer(controller)
	server.RegisterHandlers()
	// get router and start server
	router := server.GetRouter()
	if err := http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(PORT)), router); err != nil {
		fmt.Printf("server startup failed %v", err)
		os.Exit(0)
	}
}
