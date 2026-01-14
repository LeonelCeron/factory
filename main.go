package main

import (
	"factory/configuration"
	"factory/repository"
	"log"
)

func main() {
	config := configuration.Configuration{
		Engine: "SqlServer",
		Host:   "3306",
	}

	repo, err := repository.New(config)

	if err != nil {
		log.Fatal(err)
	}

	err = repo.Save("hello world")

	if err != nil {
		log.Fatal(err)
	}

	data := repo.Find(1)
	log.Print(data)
}
