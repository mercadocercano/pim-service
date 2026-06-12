package config

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hornosg/go-shared/infrastructure/env"
	goshpostgres "github.com/hornosg/go-shared/infrastructure/postgres"
)

// DatabaseConfig contiene la configuración de la base de datos
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewDatabaseConfig crea una nueva configuración de base de datos desde variables de entorno
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     env.Get("DB_HOST", "localhost"),
		Port:     env.Get("DB_PORT", "5432"),
		User:     env.Get("DB_USER", "postgres"),
		Password: env.Get("DB_PASSWORD", "postgres"),
		DBName:   env.Get("DB_NAME", "pim"),
		SSLMode:  env.Get("DB_SSL_MODE", "disable"),
	}
}

// GetDSN retorna la cadena de conexión para PostgreSQL
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

// Connect establece la conexión con la base de datos
func (c *DatabaseConfig) Connect() (*sql.DB, error) {
	db, err := goshpostgres.Connect(goshpostgres.Config{
		Host:     c.Host,
		Port:     c.Port,
		User:     c.User,
		Password: c.Password,
		DBName:   c.DBName,
		SSLMode:  c.SSLMode,
	})
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %v", err)
	}

	goshpostgres.StartPoolMonitor(context.Background(), db, goshpostgres.MonitorOptions{Service: "pim-service", DBName: c.DBName})

	return db, nil
}
