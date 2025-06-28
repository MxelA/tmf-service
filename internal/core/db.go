package core

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type DatabaseMongo struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func NewDatabaseMongo(l *Logger) *DatabaseMongo {
	_ = godotenv.Load()

	host, ok := os.LookupEnv("MONGO_DB_HOST")
	if !ok {
		log.Fatalf(".env property MONGO_DB_HOST not found")
	}

	port, ok := os.LookupEnv("MONGO_DB_PORT")
	if !ok {
		log.Fatalf(".env property MONGO_DB_PORT not found")
	}

	dbName, ok := os.LookupEnv("MONGO_DB_DATABASE")
	if !ok {
		log.Fatalf(".env property MONGO_DB_DATABASE not found")
	}

	dbUser, ok := os.LookupEnv("MONGO_DB_USERNAME")
	if !ok {
		log.Fatalf(".env property MONGO_DB_USERNAME not found")
	}

	dbPassword, ok := os.LookupEnv("MONGO_DB_PASSWORD")
	if !ok {
		log.Fatalf(".env property MONGO_DB_PASSWORD not found")
	}

	var mongoURL = "mongodb://" + dbUser + ":" + dbPassword + "@" + host + ":" + port + "/" + dbName + "?authMechanism=SCRAM-SHA-256&authSource=admin"

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(dbName)

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	l.GetCore().Info("Ping MongoDB success!")

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	l.GetCore().Info("Connected to MongoDB!", "host", host, "port", port, "database", dbName)

	return &DatabaseMongo{
		Client: client,
		Db:     db,
	}
}

func (db *DatabaseMongo) GetCore() *DatabaseMongo {
	return db
}

type DatabaseNeo4j struct {
	core *neo4j.DriverWithContext
}

func NewDatabaseNeo4j(l *Logger) *DatabaseNeo4j {
	_ = godotenv.Load()

	host, ok := os.LookupEnv("NEO4J_DB_HOST")
	if !ok {
		log.Fatalf(".env property NEO4J_DB_HOST not found")
	}

	port, ok := os.LookupEnv("NEO4J_DB_PORT")
	if !ok {
		log.Fatalf(".env property NEO4J_DB_PORT not found")
	}

	username, ok := os.LookupEnv("NEO4J_DB_USERNAME")
	if !ok {
		log.Fatalf(".env property NEO4J_DB_USERNAME not found")
	}

	password, ok := os.LookupEnv("NEO4J_DB_PASSWORD")
	if !ok {
		log.Fatalf(".env property NEO4J_DB_PASSWORD not found")
	}

	ctx := context.Background()

	dbUri := host + ":" + port
	dbUser := username
	dbPassword := password

	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))

	if err != nil {
		log.Fatalf("Failed to connect to Neo4j database", err)
	}

	defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		log.Fatalf("Failed to Verify Connectivity to Neo4j database", err)
	}

	l.core.Info("Connection established to Neo4j database")

	return &DatabaseNeo4j{core: &driver}
}

func (db *DatabaseNeo4j) GetCore() *neo4j.DriverWithContext {
	return db.core
}
