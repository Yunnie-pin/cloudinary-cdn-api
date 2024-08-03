package main

import (
	"claudinary-cdn-api/config"
	"claudinary-cdn-api/router"
	"log"
)

func main() {
	loadConfig, _ := config.LoadConfig()

	db, err := config.StartDB(&loadConfig)

	if err != nil {
		log.Fatalf("🧊 Tidak bisa terhubung ke database.")
	}

	if db != nil {
		log.Println("🧊🧊🧊 Berhasil terhubung ke database!")
	}

	r := router.NewRouter()

	log.Println("🧊 ENV: ", loadConfig.ENV)
	log.Println("🧊 Menjalankan di port :", loadConfig.PORT)

	r.Run(":" + loadConfig.PORT)
}
