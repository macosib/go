package main

import (
	"Homework_30.5.1/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	store := handlers.GetStore()
	router := chi.NewRouter()
	router.Delete("/api/v1/user", store.DeleteUserView)
	router.Post("/api/v1/make_friends", store.MakeFriendsView)
	router.Post("/api/v1/create", store.CreateUserView)
	router.Get("/api/v1/friends/{UserId}", store.GetFriendsListView)
	router.Patch("/api/v1/{UserId}", store.UpdateUserAgeView)
	http.ListenAndServe("localhost:8080", router)
}
