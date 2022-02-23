package database

import (
	"fmt"
	"log"
	"os"

	"github.com/nemonicgod/terraforms-api/config"
	"github.com/nemonicgod/terraforms-api/infra/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// _ "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
)

// var ctx = context.Background()

// Connect kind of explanatory
func Connect(c config.Reader) *gorm.DB {
	// env := c.GetString(config.Environment)
	host := c.GetString(config.PGHost)
	port := c.GetString(config.PGPort)
	user := c.GetString(config.PGUser)
	pass := c.GetString(config.PGPass)
	dbname := c.GetString(config.PGData)

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, pass)
	// dbMigrationURL := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, pass, host, port, dbname)

	//Connect to DB via GORM
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: false})
	if err != nil {
		panic(err)
	}

	// Check OS role
	role := os.Getenv("ROLE")
	if role == "" {
		panic(fmt.Errorf("DB Conn: exited, no role detected"))
	}

	if role == "api" {
		// Parcel
		// -------------------------------------------------------------------------------
		if err = db.Debug().AutoMigrate(&models.Parcel{}); err != nil {
			log.Fatalf("cannot migrate [parcel] table: %v", err)
		}
		// -------------------------------------------------------------------------------

		// m_files, err := ioutil.ReadDir(os.Getenv("MIGRATIONS_PATH"))
		// if err != nil {
		// 	log.Printf("Error reading migration indexes from directory: %v", err)
		// }

		// max_migration_num := 0

		// for _, f := range m_files {
		// 	if !f.IsDir() {
		// 		s := strings.Split(f.Name(), "_")
		// 		index, err := strconv.Atoi(s[0])

		// 		if err != nil {
		// 			log.Printf("Invalid Migration Index; Not converted to int: %v", err)
		// 		}
		// 		if index > max_migration_num {
		// 			max_migration_num = index
		// 		}

		// 	}
		// }

		// log.Printf("Max Migration Version Detected: %v", max_migration_num)

		// // Run seeds
		// if c.GetString("MIGRATE") == "no" {
		// 	return db
		// }

		// // Read migrations from /home/mattes/migrations and connect to a local postgres database.
		// m, err := migrate.New(`file://`+os.Getenv("MIGRATIONS_PATH"), dbMigrationURL)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// //Skip migrations if current-migration is up to date
		// current_migration, dirty_flag, err := m.Version()
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// // Only drop tables if local environment
		// if (env == config.Local) || (dirty_flag) {
		// 	log.Printf("Local environment or migration dirty flag detected. Dropping Tables and re-migrating....")
		// 	// Drop everything in the database ...
		// 	if err := m.Drop(); err != nil {
		// 		log.Fatal(err)
		// 	}

		// 	// Restart migration schema if tables are dropped.
		// 	m, err = migrate.New(`file://`+os.Getenv("MIGRATIONS_PATH"), dbMigrationURL)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	current_migration = 0
		// }

		// // Parcel
		// // -------------------------------------------------------------------------------
		// if err = db.Debug().AutoMigrate(&models.Parcel{}); err != nil {
		// 	log.Fatalf("cannot migrate [parcel] table: %v", err)
		// }
		// // -------------------------------------------------------------------------------

		// log.Printf("Current Migration %v", current_migration)
		// log.Printf("Max Migration %v", uint(max_migration_num))

		// //Don't migrate up if current version is up-to-date (attempting to do so will create migration error and lock further migrations)
		// if current_migration >= uint(max_migration_num) {
		// 	log.Printf("Current Migration is up to date.")
		// 	return db
		// }

		// log.Printf("Updated migrations detected. Migrating database up...")
		// // Migrate all the way up ...
		// if err := m.Up(); err != nil {
		// 	log.Fatalf("Error during Up migration: %v", err)
		// }
		// log.Printf("Up migration completed.")
	}

	return db
}
