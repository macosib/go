package database

import (
	"fmt"
	"os"
)

type PostgresDB struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
	PostgresHost     string
	PostgresPort     string
}

type Config struct {
	Postgres PostgresDB
}

//GetConfigDb - Функция возвращает ссылку на экземпляр структуры Config (конфигурация подключения к БД)
func GetConfigDb() *Config {
	return &Config{
		Postgres: PostgresDB{
			PostgresUser:     getEnv("POSTGRES_USER", ""),
			PostgresPassword: getEnv("POSTGRES_PASSWORD", ""),
			PostgresDb:       getEnv("POSTGRES_DB", ""),
			PostgresHost:     getEnv("POSTGRES_HOST", ""),
			PostgresPort:     getEnv("POSTGRES_PORT", ""),
		},
	}
}

//getEnv - Функция принимает аргументы key string и defaultVal string, возвращает string.
//Проверяет наличие переменной окружения, и возвращает её, либо значение по умолчанию
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

//GetConnectionToDatabase - Функция возвращает строку DSN для подключения к БД.
func GetConnectionToDatabase() string {
	configDB := GetConfigDb()
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configDB.Postgres.PostgresHost,
		configDB.Postgres.PostgresPort,
		configDB.Postgres.PostgresUser,
		configDB.Postgres.PostgresPassword,
		configDB.Postgres.PostgresDb,
	)
	return conn

}
