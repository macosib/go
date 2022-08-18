package main

import (
	"Homework_31.5.1/pkg/database"
	"Homework_31.5.1/pkg/handlers"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

//init - Загружает переменные окружения из .env в операционную систему.
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

//main - Основная функция программы. Регистрирует маршруты и обработчики запросов API.
func main() {
	db, err := sql.Open("postgres", database.GetConnectionToDatabase())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	database.InitDatabase(db)
	s := &handlers.Store{db}

	router := chi.NewRouter()
	router.Post("/api/v1/create", s.CreateUserView)
	router.Post("/api/v1/make_friends", s.MakeFriendsView)
	router.Delete("/api/v1/user", s.DeleteUserView)
	router.Get("/api/v1/friends/{UserId}", s.GetFriendsListView)
	router.Patch("/api/v1/{UserId}", s.UpdateUserAgeView)
	http.ListenAndServe(":80", router)
}
