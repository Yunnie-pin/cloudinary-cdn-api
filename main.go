package main

import (
	"claudinary-cdn-api/config"
	"claudinary-cdn-api/controllers"
	"claudinary-cdn-api/migration"
	"claudinary-cdn-api/repository"
	"claudinary-cdn-api/router"
	"flag"
	"log"
)

func main() {
	loadConfig, _ := config.LoadConfig()

	db, err := config.StartDB(&loadConfig)

	if err != nil {
		log.Fatalf("🧊 Tidak bisa terhubung ke database.")
	}

	// Definisikan flag --migrate
	input := flag.Bool("migrate", false, "Run migration scripts")

	// Parsing flag
	flag.Parse()
	if *input {
		//Migration
		err = migration.MigrationDB(db)
		if err != nil {
			log.Fatal("🧊 Tidak bisa melakukan migrasi database. ", err)
		}
		return
	}

	//init repository
	bucketRepository := repository.NewBucketRepository(db)

	//init controller
	bucketController := controllers.NewBucketController(bucketRepository)
	//init router
	r := router.NewRouter(
		bucketController,
	)

	log.Println("🧊 ENV: ", loadConfig.ENV)
	log.Println("🧊 Menjalankan di port :", loadConfig.PORT)

	r.Run(":" + loadConfig.PORT)
}
