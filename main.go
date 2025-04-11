package main

import (
	"fmt"
	"log"
	"os"

	deliveryHttp "a2sv.org/hub/Delivery/http"
	"a2sv.org/hub/Delivery/http/handlers"
	"a2sv.org/hub/Repository/postgres"
	"a2sv.org/hub/infrastructure"
	"a2sv.org/hub/usecases"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using environment variables")
	}

	// Initialize database connection
	db, err := infrastructure.NewDBConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established successfully")

	// Initialize repository
	userRepo := postgres.NewUserRepository(db)
	handlers.UserRepo = userRepo // Set the global UserRepo variable
	roleRepo := postgres.NewRoleRepository(db)
	groupRepo := postgres.NewGroupRepository(db)
	countryRepo := postgres.NewCountryRepository(db)
	superGroupRepo := postgres.NewSuperGroupRepository(db)

	// Initialize use case
	userUseCase := usecases.NewUserUseCase(userRepo)
	roleUseCase := usecases.NewRoleUseCase(roleRepo)
	groupUseCase := usecases.NewGroupUseCase(groupRepo)
	countryUseCase := usecases.NewCountryUseCase(countryRepo)
	bulkRegistrationUseCase := usecases.NewBulkRegistrationUseCase(userRepo, roleRepo, groupRepo, countryRepo)
	superGroupUseCase := usecases.NewSuperGroupUseCase(superGroupRepo)

	// Setup router
	router := deliveryHttp.SetupRouter(
		*userUseCase,
		*roleUseCase,
		groupUseCase,
		countryUseCase,
		bulkRegistrationUseCase,
		superGroupUseCase,
	)

	// Print all registered routes for debugging
	log.Println("Routes registered successfully")

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	router.Run(":" + port)
	fmt.Println("this is used to test the backend.yeneineh.main docker publish workflow 2")
}
