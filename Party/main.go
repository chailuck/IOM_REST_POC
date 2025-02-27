package main

import (
	"log"

	"party/internal/model"
	"party/pkg/postgresql"

	imservice "party/pkg/service"

	iomlog "gitlab.com/ft25/iom/framework/logger"
	iomcore "gitlab.com/ft25/iom/framework/rabbitmq"
	iomservice "gitlab.com/ft25/iom/framework/service"

	"party/internal/controller"
	"party/internal/repository"

	"party/internal/config"
)

// @title IOM.PartyManagement API
// @version 5.0
// @description TMF632 Party Management API Implementation
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email it-iom-dev@truecorp.co.th
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /partyManagement/v1
func main() {
	// Initialize context
	var ctx iomcore.Context
	ctx.Logger, _ = iomlog.New("party")
	ctx.RabbitMQClient = iomcore.New()

	// Initialize configurations
	cfg := config.New()

	// Initialize database connection
	db, err := postgresql.NewPostgresConnection(cfg.PostgresConfig)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	// Auto-migrate schemas
	err = db.AutoMigrate(
		&model.Party{},
		&model.PartyAddress{},
		&model.PartyAddressExt{},
		&model.PartyAttributes{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database schemas: %v", err)
	}

	// Initialize repository
	partyRepo := repository.NewPartyRepository(db)

	// Initialize service
	partyService := controller.NewPartyService(partyRepo)

	// Initialize rest service controller
	imservice.NewRestServiceController(partyService)

	// Initialize route services
	routeServices, err := iomservice.LoadService(ctx, "party.Partymanagement", []interface{}{imservice.RestService{}})
	if err != nil {
		log.Fatalf("Failed to load service: %v", err)
	}

	// Initialize REST service
	/*
		srv := iomservice.NewRestService(
			iomservice.WithAuthen(&ctx),
			iomservice.WithHealthcheck(),
			iomservice.WithLogger(nil),
			iomservice.WithRoute(routeServices),
			iomservice.WithSwagger("/swagger/*", swagger.HandlerDefault),
		)
	*/
	srv := iomservice.NewRestService(
		iomservice.WithRoute(routeServices),
		iomservice.WithLogger(nil),
	)

	// Start the service
	log.Println("Starting Party Management Service...")
	srv.Start()
}
