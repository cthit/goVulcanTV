package db

import (
	"encoding/json"
	"fmt"
	"github.com/swexbe/govulcantv/internal/db/commands"
	"github.com/swexbe/govulcantv/internal/db/models"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"github.com/swexbe/govulcantv/internal/db/common"
	"os"
)

func setupDB(db *gorm.DB) {
	createTables(db)
	loadDefaults(db)
}

func createTables(db *gorm.DB) {
	createTable(db, models.PageContent{})
}

func createTable(db *gorm.DB, model common.NamedStruct) {
	err := db.AutoMigrate(model)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to migrate table %s", model.StructName()))
	} else {
		log.Println(fmt.Sprintf("Successfully migrated table %s", model.StructName()))
	}
}

type DefaultJson struct {
	PageContents []models.PageContent `json:"pageContents"`
}

func loadDefaults(db *gorm.DB) {
	defaultsFile, err := os.Open("internal/db/defaults.json")
	if err != nil {
		log.Fatalf("Failed to load db defaults, err: %s", err)
	}

	defer defaultsFile.Close()

	byteVal, err := ioutil.ReadAll(defaultsFile)
	if err != nil {
		log.Fatalf("Unable to read db defaults, err: %s", err)
	}

	var defaultDb DefaultJson
	err = json.Unmarshal(byteVal, &defaultDb)
	if err != nil {
		log.Fatalf("Unable to parse json due to err: %s", err)
	}

	for _, pageContent := range defaultDb.PageContents {
		err := commands.InsertPageContent(&pageContent)
		if err != nil {
			log.Printf("Failed to insert default page content %+v to database, err %s\n", pageContent, err)
		}
	}
}