package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config структура для хранения конфигурационных данных
type Config struct {
	AppEnv       string
	ServerAdress string
	PgUsername   string
	PgPassword   string
	PgHost       string
	PgPort       int
	PgDb         string
	RandomSecret string
}

// GlobalConfig содержит текущую конфигурацию
var GlobalConfig Config

// InitConfig инициализирует конфигурацию приложения
func InitConfig() {
	// Загружаем .env файл, если он существует
	if _, err := os.Stat(".env"); err == nil {
		if loadErr := godotenv.Load(".env"); loadErr != nil {
			log.Printf("Ошибка загрузки .env файла: %v", loadErr)
		}
	} else {
		log.Println(".env файл отсутствует, используются только переменные окружения и значения по умолчанию")
	}

	// Установка конфигурации
	GlobalConfig = Config{
		AppEnv:       resolveEnv("APP_ENV", "prod"),
		ServerAdress: resolveEnv("SERVER_ADDRESS", "0.0.0.0:8080"),
		PgUsername:   resolveEnv("POSTGRES_USERNAME", "root"),
		PgPassword:   resolveEnv("POSTGRES_PASSWORD", "qwerty"),
		PgHost:       resolveEnv("POSTGRES_HOST", "localhost"),
		PgPort:       resolveEnvAsInt("POSTGRES_PORT", 5432),
		PgDb:         resolveEnv("POSTGRES_DATABASE", ""),
		RandomSecret: resolveEnv("RANDOM_SECRET", "qwerty"),
	}
}

// resolveEnv получает строковое значение из переменной окружения, затем из .env или использует значение по умолчанию
func resolveEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	// Если переменная не определена в окружении, читаем из .env (godotenv делает это автоматически)
	return getEnv(key, defaultValue)
}

// resolveEnvAsInt получает числовое значение из переменной окружения, затем из .env или использует значение по умолчанию
func resolveEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	// Если переменная не определена в окружении, читаем из .env (godotenv делает это автоматически)
	return getEnvAsInt(key, defaultValue)
}

// getEnv получает строковое значение из .env или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt получает числовое значение из .env или возвращает значение по умолчанию
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

// GetConfig возвращает текущую конфигурацию
func GetConfig() Config {
	return GlobalConfig
}
