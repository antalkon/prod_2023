package pggorm

import (
	"fmt"
	"log"

	"github.com/antalkon/prod_2023/internal/config"
	"github.com/antalkon/prod_2023/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB инициализирует подключение к базе данных и выполняет миграцию моделей
func InitDB() error {
	// Получаем конфигурацию через config.GetConfig()
	cfg := config.GetConfig()

	// Формируем строку подключения
	dsn := createDSN(&cfg) // Передаем указатель на cfg
	log.Printf("Connecting to database at %s", cfg.PgHost)

	// Подключаемся к базе данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true, // Кэширование подготовленных запросов
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	DB = db
	log.Println("Connected to database successfully.")

	// Выполняем миграцию моделей
	if err := migrateModels(); err != nil {
		return fmt.Errorf("failed to migrate models: %w", err)
	}

	log.Println("Database migration completed successfully.")
	return nil
}

// createDSN создает строку подключения к базе данных на основе конфигурации
func createDSN(cfg *config.Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.PgHost,
		cfg.PgUsername,
		cfg.PgPassword,
		cfg.PgDb,
		cfg.PgPort,
	)
}

// migrateModels выполняет миграцию всех моделей
func migrateModels() error {
	modelsToMigrate := []interface{}{
		&models.User{},
	}

	for _, model := range modelsToMigrate {
		log.Printf("Migrating model: %T", model)
		if err := DB.AutoMigrate(model); err != nil {
			log.Printf("Failed to migrate model: %T, error: %v", model, err)
			return fmt.Errorf("failed to migrate model: %T, error: %w", model, err)
		}
	}

	log.Println("All models migrated successfully.")
	return nil
}
