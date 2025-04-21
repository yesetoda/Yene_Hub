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

// @title           Hub API
// @version         1.0
// @description     This is the API documentation for the Hub backend.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.email  support@hub.a2sv.org
// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT
// @host      yene-hub-ls0y.onrender.com
// @BasePath  /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
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
	recentActionRepo := postgres.NewRecentActionRepository(db) // should implement repository.RecentActionRepository
	voteRepo := postgres.NewVoteRepository(db)                 // should implement repository.VoteRepository
	trackRepo := postgres.NewTrackRepository(db)               // should implement repository.TrackRepository
	stippendRepo := postgres.NewStipendRepository(db)          // should implement repository.StipendRepository
	submissionRepo := postgres.NewSubmissionRepository(db)     // should implement repository.SubmissionRepository
	superToGroupRepo := postgres.NewSuperToGroupRepository(db) // should implement repository.SuperToGroupRepository
	problemRepo := postgres.NewProblemRepository(db)           // should implement repository.ProblemRepository
	sessionRepo := postgres.NewSessionRepository(db)           // should implement repository.SessionRepository

	// Initialize use case
	userUseCase := usecases.NewUserUseCase(userRepo)
	roleUseCase := usecases.NewRoleUseCase(roleRepo)
	groupUseCase := usecases.NewGroupUseCase(groupRepo)
	countryUseCase := usecases.NewCountryUseCase(countryRepo)
	bulkRegistrationUseCase := usecases.NewBulkRegistrationUseCase(userRepo, roleRepo, groupRepo, countryRepo)
	superGroupUseCase := usecases.NewSuperGroupUseCase(superGroupRepo)
	recentActionUseCase := usecases.NewRecentActionUsecase(recentActionRepo)
	voteUseCase := usecases.NewVoteUsecase(voteRepo)
	trackUseCase := usecases.NewTrackUsecase(trackRepo)
	stippendUseCase := usecases.NewStipendUsecase(stippendRepo)
	submissionUseCase := usecases.NewSubmissionUsecase(submissionRepo)
	superToGroupUseCase := usecases.NewSuperToGroupUsecase(superToGroupRepo)
	problemUseCase := usecases.NewProblemUsecase(problemRepo)
	sessionUsecase := usecases.NewSessionUsecase(sessionRepo)
	// Setup router
	router := deliveryHttp.SetupRouter(
		*userUseCase,
		*roleUseCase,
		groupUseCase,
		countryUseCase,
		bulkRegistrationUseCase,
		superGroupUseCase,
		*recentActionUseCase,
		*voteUseCase,
		*trackUseCase,
		*superToGroupUseCase,
		*submissionUseCase,
		*stippendUseCase,
		problemUseCase,
		*sessionUsecase,
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
