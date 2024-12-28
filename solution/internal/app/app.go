package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/antalkon/prod_2023/internal/config"
	"github.com/antalkon/prod_2023/internal/transport/rest/router"
	"github.com/antalkon/prod_2023/pkg/logs"
	"github.com/antalkon/prod_2023/pkg/pggorm"
)

func Run() error {
	startTime := time.Now()

	// Инициализация модулей
	if err := initModules(); err != nil {
		panic("Ошибка инициализации модулей: " + err.Error())
	}

	// Запуск REST API сервера
	if err := startRESTServer(); err != nil {
		return err
	}

	// Логирование успешного завершения запуска
	logger := logs.GetLogger()
	logger.Infof("Приложение запущено успешно за %s", time.Since(startTime))
	return nil
}

func initModules() error {
	// 1. Инициализация конфигурации
	config.InitConfig()

	// 2. Инициализация логгера
	cfg := config.GetConfig()
	logs.InitLogger(cfg.AppEnv)

	// 3. Инициализация базы данных
	if err := pggorm.InitDB(); err != nil {
		return err
	}

	// Логирование завершения инициализации
	logger := logs.GetLogger()
	logger.Infof("Инициализация модулей завершена. Среда: %s", cfg.AppEnv)
	return nil
}

func startRESTServer() error {
	cfg := config.GetConfig()
	logger := logs.GetLogger()

	// Создаем Echo сервер через функцию маршрутов
	e := router.NewEchoServer()

	// Канал для обработки сигналов завершения
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Запуск сервера в отдельной горутине
	go func() {
		logger.Infof("Запуск REST API сервера на %s", cfg.ServerAdress)
		if err := e.Start(cfg.ServerAdress); err != nil {
			if err != http.ErrServerClosed {
				logger.Fatalf("Ошибка запуска сервера: %v", err)
			}
		}
	}()

	// Ожидание сигнала завершения
	<-stopChan
	logger.Info("Получен сигнал завершения. Остановка сервера...")

	// Завершение работы сервера
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Errorf("Ошибка при остановке сервера: %v", err)
		return err
	}
	logger.Info("Сервер успешно остановлен.")
	return nil
}
