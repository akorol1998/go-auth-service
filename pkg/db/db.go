package db

import (
	"fmt"
	"log"

	"github.com/akorol1998/go-auth-service/pkg/config"
	"github.com/akorol1998/go-auth-service/pkg/models"
	"github.com/akorol1998/go-auth-service/pkg/utils"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Handler struct {
	DB *gorm.DB
}

func Init(c config.Config) Handler {
	var db *gorm.DB

	err := godotenv.Load("./pkg/config/envs/dev.env")
	if err != nil {
		log.Println("Failed to read .env file")
	}

	// connectionString := username + ":" + password + "@tcp" + "(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.PostgresHost, c.PostgresUser, c.PostgresPassword, c.PostgresDb, c.PostgresPort)

	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open initialize session to Database")
	}
	log.Printf("Successfully connected - host:port:databaseName %s:%s:%s", c.PostgresHost, c.PostgresPort, c.PostgresDb)
	// _, err = db.DB()

	// Read about this. What is the idle pool and open connections pool
	// genDB.SetConnMaxIdleTime()

	if err != nil {
		log.Fatal("Error happened while connecting to Database instance")
	}
	log.Println("Innit function of domain package is ran!")

	//Printing queries
	db.Logger.LogMode(logger.Info)

	// This command is pretty obvious))
	db.Migrator().DropTable(&models.UserRoles{}, "role_permissions", &models.Role{}, &models.Permission{}, &models.User{}, &models.Order{})

	// Setting up custom many2many
	// Important: must be run before creational Migrations
	db.SetupJoinTable(&models.User{}, "Roles", &models.UserRoles{})
	db.SetupJoinTable(&models.Role{}, "Users", &models.UserRoles{})

	db.AutoMigrate(&models.User{}, &models.Order{}, &models.Role{})
	return Handler{DB: db}
}

func InitialFixture(h Handler) error {
	desc := "Starting role for further roles/permissions management." +
		"Has access to all CRUD operations in the system."
	res := h.DB.Create(&models.Role{
		Name:        "SuperAdmin",
		Description: &desc,
		Permissions: models.InitialPermissions,
		Users:       []models.User{{Name: "SuperUser", Email: "super_user@gmail.com", Password: utils.HashPassword("1234admin")}},
	})
	return res.Error
}
