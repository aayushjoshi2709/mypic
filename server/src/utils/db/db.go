package db

import (
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var db *mongo.Database

func Init() {
	db = getDB()
}

func GetConn() *mongo.Database {
	if db == nil {
		db = getDB()
	}
	return db
}

func getDB() *mongo.Database {
	uri := os.Getenv("MONGO_DB_URL")

	if uri == "" {
		log.Fatal("MONGO_DB_URL environment variable is not set")
	}

	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	maxPoolSize, err := strconv.Atoi(os.Getenv("MONGO_DB_MAX_POOL_SIZE"))

	if err != nil {
		log.Println("MONGO_DB_MAX_POOL_SIZE environment variable is not set or invalid, defaulting to 10")
		maxPoolSize = 10
	}

	minPoolSize, err := strconv.Atoi(os.Getenv("MONGO_DB_MIN_POOL_SIZE"))

	if err != nil {
		log.Println("MONGO_DB_MIN_POOL_SIZE environment variable is not set or invalid, defaulting to 5")
		minPoolSize = 5
	}

	dbName := os.Getenv("MONGO_DB_NAME")

	if dbName == "" {
		log.Println("MONGO_DB_NAME environment variable is not set, defaulting to 'mypic'")
		dbName = "mypic"
	}

	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverApi).
		SetCompressors([]string{"snappy", "zlib", "zstd"}).
		SetMaxPoolSize(uint64(maxPoolSize)).
		SetMinPoolSize(uint64(minPoolSize)).
		SetMaxConnIdleTime(30 * time.Second)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	} else {
		log.Print("Successfully connected to MongoDB")
	}
	return client.Database(dbName)
}
