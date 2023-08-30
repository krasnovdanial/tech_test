package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"dynamic_segmentation/config"
	"dynamic_segmentation/internal/db"
	"dynamic_segmentation/internal/server"
	_ "github.com/lib/pq"
)

func main() {
	// Путь до файла конфигурации
	configPath := "config/config.yml"

	// Инициализация конфигурации
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err) // Программа завершится, если не удастся загрузить конфигурацию
	}

	// Подключение к базе данных
	sqlDB, err := connectToDB(cfg)
	if err != nil {
		log.Fatal(err) // Программа завершится, если не удастся подключиться к базе данных
	}

	myDB := db.NewDB(sqlDB)

	// Запуск приложения
	srv := server.NewApp(myDB).Run(cfg)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Завершение работы сервера ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Завершение работы сервера:", err)
	}

	<-ctx.Done()
	log.Println("Тайм-аут 5 секунд.")
	log.Println("Сервер завершает работу")
}

// Функция для подключения к базе данных
func connectToDB(cfg *config.Config) (*sql.DB, error) {
	var err error

	sqlDb, err := sql.Open("postgres", cfg.PG.URL) // для запуска локально использовать cfg.PG.URLLocal
	if err != nil {
		return nil, err // Возвращает ошибку, если не удается создать соединение
	}

	// Проверка соединения с базой данных
	if err := sqlDb.Ping(); err != nil {
		return nil, err // Возвращает ошибку, если не удается установить соединение
	}

	return sqlDb, nil // Возвращает подключение к базе данных
}
