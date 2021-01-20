package db

import (
	"fmt"
	"github.com/swexbe/govulcantv/internal/db/commands"
	"github.com/swexbe/govulcantv/internal/db/queries"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// TODO: Rename to 'init' when this package is used (so this is called when imported rather than manually).
func Init() {
	username := os.Getenv("db_user")
	password := os.Getenv("db_password")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				Colorful:      true,
				LogLevel:      logger.Info,
			},
		),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to database, err: %s", err))
	}

	commands.Init(conn)
	queries.Init(conn)

	setupDB(conn)

	log.Println("Initialized database connection")
}

