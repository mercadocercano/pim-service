package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBClient encapsula la conexión a MongoDB
type MongoDBClient struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// NewMongoDBClient crea una nueva conexión a MongoDB
func NewMongoDBClient() (*MongoDBClient, error) {
	// Obtener configuración desde variables de entorno
	host := getEnvOrDefault("MONGO_HOST", "localhost")
	port := getEnvOrDefault("MONGO_PORT", "27017")
	user := getEnvOrDefault("MONGO_USER", "admin")
	password := getEnvOrDefault("MONGO_PASSWORD", "admin123")
	database := getEnvOrDefault("MONGO_DATABASE", "pim_marketplace")

	// Construir URI de conexión
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
		user, password, host, port, database)

	log.Printf("🔌 Conectando a MongoDB: %s:%s/%s", host, port, database)

	// Configurar opciones del cliente
	clientOptions := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(10).
		SetMinPoolSize(2).
		SetMaxConnIdleTime(30 * time.Second).
		SetServerSelectionTimeout(5 * time.Second).
		SetConnectTimeout(10 * time.Second)

	// Crear contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conectar al cliente
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error conectando a MongoDB: %w", err)
	}

	// Verificar la conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error haciendo ping a MongoDB: %w", err)
	}

	log.Printf("✅ Conexión exitosa a MongoDB: %s", database)

	// Obtener referencia a la base de datos
	db := client.Database(database)

	return &MongoDBClient{
		Client:   client,
		Database: db,
	}, nil
}

// Close cierra la conexión a MongoDB
func (mc *MongoDBClient) Close() error {
	if mc.Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := mc.Client.Disconnect(ctx)
		if err != nil {
			return fmt.Errorf("error cerrando conexión MongoDB: %w", err)
		}

		log.Println("🔌 Conexión MongoDB cerrada")
	}
	return nil
}

// HealthCheck verifica el estado de la conexión
func (mc *MongoDBClient) HealthCheck(ctx context.Context) error {
	if mc.Client == nil {
		return fmt.Errorf("cliente MongoDB no inicializado")
	}

	err := mc.Client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("MongoDB no responde: %w", err)
	}

	return nil
}

// GetCollection obtiene una colección específica
func (mc *MongoDBClient) GetCollection(name string) *mongo.Collection {
	return mc.Database.Collection(name)
}

// getEnvOrDefault obtiene una variable de entorno o retorna un valor por defecto
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
