package handlers

import (
	"Homework_30.5.1/pkg/user"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Store struct {
	Store map[int]*user.User
}

func GetStore() *Store {
	return &Store{make(map[int]*user.User)}
}

func (s *Store) CreateUserView(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "POST":
		{
			content, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			defer r.Body.Close()
			var newUser user.User
			if err := json.Unmarshal(content, &newUser); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			response, _ := json.Marshal(userCreateData(s, &newUser))
			fmt.Println("response", response)
			w.Write(response)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (s *Store) GetFriendsListView(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		response, err := getFriendList(s, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(response)
		w.Write(res)

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (s *Store) UpdateUserAgeView(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "PATCH":
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		var newAge user.UserAge
		if err := json.Unmarshal(content, &newAge); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = updateUserAge(s, r, newAge.Age)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		response := "Возраст пользователя успешно обновлен!"
		w.Write([]byte(response))
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (s *Store) MakeFriendsView(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "POST":
		{
			content, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			defer r.Body.Close()
			var newFriends user.Friend
			if err := json.Unmarshal(content, &newFriends); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			sourseUser, targetUser, err := makeFriends(s, &newFriends)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			response := sourseUser + " и " + targetUser + " теперь друзья"
			w.Write([]byte(response))
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (s *Store) DeleteUserView(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "DELETE":
		{
			content, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			defer r.Body.Close()
			var user user.UserId
			if err := json.Unmarshal(content, &user); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			userName, err := deleteUser(s, &user)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			response := userName + " удален!"
			w.Write([]byte(response))
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getFriendList(s *Store, r *http.Request) (map[string][]int, error) {
	idParam := chi.URLParam(r, "UserId")
	idUser, err := strconv.Atoi(idParam)
	if err != nil {
		return nil, err
	}
	result := make(map[string][]int)
	user, ok := s.Store[idUser]
	if !ok {
		return nil, errors.New("Такого пользователя нет!")
	}
	result["friends"] = user.Friends
	return result, nil
}

func updateUserAge(s *Store, r *http.Request, age int) error {
	idParam := chi.URLParam(r, "UserId")
	idUser, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}
	user, ok := s.Store[idUser]
	if !ok {
		return errors.New("Такого пользователя нет!")
	}
	user.Age = age
	return nil
}

func userCreateData(s *Store, newUser *user.User) map[string]int {
	result := make(map[string]int)
	idUser := len(s.Store) + 1
	s.Store[idUser] = newUser
	result["UserId"] = idUser
	fmt.Println("res", result)
	return result
}

func makeFriends(s *Store, newFriends *user.Friend) (string, string, error) {

	sourseUser, ok := s.Store[newFriends.SourceId]
	if !ok {
		return "", "", errors.New("Такого пользователя нет!")
	}
	targetUser, ok := s.Store[newFriends.TargetId]
	if !ok {
		return "", "", errors.New("Такого пользователя нет!")
	}
	// Можно добавить дополднительные проверки, что они уже есть в друзьях и т.д.
	sourseUser.Friends = append(sourseUser.Friends, newFriends.TargetId)
	targetUser.Friends = append(targetUser.Friends, newFriends.SourceId)
	return sourseUser.Name, targetUser.Name, nil
}

func deleteUser(s *Store, userTarget *user.UserId) (string, error) {

	userName, ok := s.Store[userTarget.TargetId]
	if !ok {
		return "", errors.New("Такого пользователя нет!")
	}
	for _, value := range s.Store {
		userIndex := 0
		flag := false
		for index, userId := range value.Friends {
			if userTarget.TargetId == userId {
				flag = true
				userIndex = index
			}
		}
		if flag {
			value.Friends = remove(value.Friends, userIndex)
		}
	}
	return userName.Name, nil
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
