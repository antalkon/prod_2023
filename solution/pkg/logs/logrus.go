package logs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Логгер доступен глобально
var logger = logrus.New()

// customTextFormatter добавляет цветное форматирование и префикс LOGRUS в dev режиме
type customTextFormatter struct {
	ForceColors      bool
	FullTimestamp    bool
	TimestampFormat  string
	DisableTimestamp bool
}

func (f *customTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := map[logrus.Level]string{
		logrus.DebugLevel: "\033[36m", // Cyan
		logrus.InfoLevel:  "\033[32m", // Green
		logrus.WarnLevel:  "\033[33m", // Yellow
		logrus.ErrorLevel: "\033[31m", // Red
		logrus.FatalLevel: "\033[35m", // Magenta
		logrus.PanicLevel: "\033[41m", // Background Red
	}

	reset := "\033[0m"
	color, exists := levelColor[entry.Level]
	if !exists {
		color = reset
	}

	timestamp := ""
	if !f.DisableTimestamp {
		timestamp = fmt.Sprintf("[%s] ", entry.Time.Format(f.TimestampFormat))
	}

	logMessage := fmt.Sprintf(
		"LOGRUS: %s%s%-8s%s %s\n",
		color,
		timestamp,
		entry.Level.String(),
		reset,
		entry.Message,
	)

	return []byte(logMessage), nil
}

// InitLogger настраивает логгер в зависимости от окружения
func InitLogger(environment string) {
	// Настройка формата вывода
	if environment == "prod" {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	} else {
		// Настройка кастомного текстового форматирования в режиме dev
		logger.SetFormatter(&customTextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
		})
	}

	// Установка уровня логирования
	switch environment {
	case "dev":
		logger.SetLevel(logrus.DebugLevel) // Показать все логи, включая Debug
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "prod":
		logger.SetLevel(logrus.InfoLevel)
	default:
		logger.SetLevel(logrus.WarnLevel)
	}

	// Путь к файлу логов
	logDir := "data/logs"
	logFile := filepath.Join(logDir, "app.log")

	// Создать папку для логов, если её нет
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			logger.Fatalf("Ошибка при создании директории логов: %v", err)
		}
	}

	// Настройка ротации файлов
	fileWriter := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    200, // Максимальный размер файла в МБ
		MaxAge:     7,   // Максимальное количество дней хранения файла
		MaxBackups: 3,   // Максимальное количество резервных копий
		Compress:   true,
	}

	// Настройка вывода
	if environment == "prod" {
		// Только файл в продакшене
		logger.SetOutput(io.MultiWriter(fileWriter))
	} else {
		// Консоль и файл в dev/debug
		logger.SetOutput(io.MultiWriter(os.Stdout, fileWriter))
	}
}

// GetLogger возвращает глобальный логгер
func GetLogger() *logrus.Logger {
	return logger
}
