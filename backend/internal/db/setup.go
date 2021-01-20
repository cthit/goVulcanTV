package db

import (
	"encoding/json"
	"fmt"
	"github.com/swexbe/govulcantv/internal/db/common"
	"github.com/swexbe/govulcantv/internal/db/models"
	"github.com/swexbe/govulcantv/internal/process"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
)

func setupDB(db *gorm.DB) {
	resetDB(db)
	createTables(db)
	loadDefaults(db)
}

func resetDB(db *gorm.DB) {
	reset := os.Getenv("reset_db")

	if reset == "true" {
		log.Printf("==== RESETTING DB ====")
		db.Exec("DROP SCHEMA public CASCADE")
		db.Exec("CREATE SCHEMA public")
	}
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
		err := process.CreatePageContent(&pageContent)
		if err != nil {
			log.Printf("Failed to insert default page content %+v to database, err %s\n", pageContent, err)
		}
	}
}