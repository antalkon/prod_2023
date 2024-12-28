package main

import (
	"log"

	"github.com/antalkon/prod_2023/internal/app"
)

func main() {
	// Запуск приложения
	if err := app.Run(); err != nil {
		log.Fatalf("Ошибка запуска приложения: %v", err)
	}
}
