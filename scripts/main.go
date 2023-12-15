package main

import (
	"blocklist/internal/config"
	"blocklist/internal/infra/database/mongodb"
	"blocklist/internal/utils"
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PhoneNumber struct {
	Number string `bson:"number"`
}

func generateRandomPhoneNumber() string {
	return fmt.Sprintf("%010d", rand.Intn(1e10))
}

func main() {
	// Initialize the config
	config.LoadEnv()
	utils.LogInfo("env loaded")

	// Initialize the database
	databaseConfig := config.NewDatabaseConfig()
	utils.LogInfo("database config loaded")

	// Initialise the connection to the database
	conn := mongodb.GetInstance(databaseConfig)
	utils.LogInfo("database connection initialised")

	defer func() {
		if err := conn.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Get a reference to the phone_numbers collection
	collection := conn.Database(databaseConfig.Dbname).Collection("blocklist")

	// Create an index on the "number" field
	indexModel := mongo.IndexModel{
		Keys:    map[string]interface{}{"number": 1},
		Options: options.Index().SetUnique(true),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal(err)
	}

	// Ingest 50,000 random phone numbers using goroutines and waitgrp 
	fmt.Println(time.Now())
	wg := sync.WaitGroup{}
	wg.Add(50000)

	for i := 0; i < 50000; i++ {
		go IngestToDb(collection, ctx, &wg)	
	}

	wg.Wait()
	fmt.Println(time.Now())

	fmt.Println("Ingestion complete.")
}

func IngestToDb(collection *mongo.Collection, ctx context.Context, wg *sync.WaitGroup){
	defer wg.Done()
		phoneNumber :=  generateRandomPhoneNumber()
		_, err := collection.InsertOne(ctx, PhoneNumber{Number: phoneNumber})
		if err != nil {
			log.Printf("Error inserting phone number %s: %v\n", phoneNumber, err)
		}
}
