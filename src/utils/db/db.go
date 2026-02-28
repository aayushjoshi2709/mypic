package db

import (
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

func Init() {
	client = getClient()
}

func New() *mongo.Client {
	if client == nil {
		client = getClient()
	}
	return client
}

func getClient() *mongo.Client {
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
	return client
}
