package handlers

import (
	"Homework_30.5.1/pkg/user"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Store - структура для хранения пользователей.
type Store struct {
	Store map[int]*user.User
}

// GetStore - Функция возвращает ссылку на экземпляр структуры Store в формате map[int]*user.User
func GetStore() *Store {
	return &Store{make(map[int]*user.User)}
}

// CreateUserView - Метод структуры Store, который обрабатывает POST запрос на создание пользователя.
// Принимает на вход объект, соответствующий интерфейсу http.ResponseWriter и ссылку объект структуры http.Request.
// В ответ отправляет ID созданного пользователя в формате JSON и код состояния HTTP, либо ошибку в текстовом формате
// и код состояния HTTP. Тело запроса должно быть в формате {"name":"some name","age":"24","friends":[]}
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
			w.Write(response)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// GetFriendsListView - Метод структуры Store, который обрабатывает GET запрос на получение списка друзей пользователя.
// Принимает на вход объект, соответствующий интерфейсу http.ResponseWriter и ссылку объект структуры http.Request.
// В ответ отправляет список друзей пользователя в формате JSON и код состояния HTTP, либо ошибку в текстовом формате
// и код состояния HTTP.
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

// UpdateUserAgeView - Метод структуры Store, который обрабатывает PATCH запрос на изменение возраста пользователя.
// Принимает на вход объект, соответствующий интерфейсу http.ResponseWriter и ссылку объект
// структуры http.Request. В ответ отправляет статус обновления возраста и код состояния HTTP,
// либо ошибку в текстовом формате и код состояния HTTP. Тело запроса должно быть в формате {"new age":"28"}
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

// MakeFriendsView - Метод структуры Store, который обрабатывает POST запрос на добавление пользователей в друзья.
// Принимает на вход объект, соответствующий интерфейсу http.ResponseWriter и ссылку объект структуры http.Request.
// В ответ отправляет статус добавления в друзья и код состояния HTTP, либо ошибку в текстовом формате
// и код состояния HTTP. Тело запроса должно быть в формате {"source_id":"1","target_id":"2"}
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

// DeleteUserView - Метод структуры Store, который обрабатывает POST запрос на удаление пользователя.
// Принимает на вход объект, соответствующий интерфейсу w http.ResponseWriter и ссылку объект структуры r http.Request.
// В ответ отправляет ID удаленного пользователя в текстовом формате и код состояния HTTP, либо ошибку в текстовом формате
// и код состояния HTTP. Тело запроса должно быть в формате {"target_id":"1"}
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

// getFriendList - Функция примает ссылку на объект структуры s Store и ссылку объект структуры r http.Request.
// Из URL извлекает параметр, который соответствует ID пользователя. Возвращает список друзей пользователя в формате
// map[string][]int и состояние ошибки.
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

// updateUserAge - Функция примает ссылку на объект структуры s Store, ссылку объект структуры r http.Request
// и новый возраст пользователя age в формате int. Из URL извлекает параметр, который соответствует ID пользователя.
// Возвращает список друзей пользователя в формате map[string][]int и состояние ошибки.
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

// userCreateData - Функция примает ссылку на объект структуры s Store, ссылку объект структуры newUser user.User
// Добавляет нового пользователя в хранилище s *Store. Возвращает информацию о новом пользователе в формате map[string][]int.
func userCreateData(s *Store, newUser *user.User) map[string]int {
	result := make(map[string]int)
	idUser := len(s.Store) + 1
	s.Store[idUser] = newUser
	result["UserId"] = idUser
	return result
}

// makeFriends - Функция примает ссылку на объект структуры s Store, ссылку объект структуры newFriends user.Friend
// Извлекает из объекта структуры user.Friend ID пользователей добавляет их в друзья друг к другу.
// Возвращает имена пользователей в формате string и состояние ошибки.
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

// deleteUser - Функция примает ссылку на объект структуры s Store, ссылку объект структуры userTarget user.UserId
// Извлекает из объекта структуры user.UserId ID пользователя и удалеет его из хранилища, а также из списка друзей всех
// пользователей. Возвращает имя удаленного пользователя в формате string и состояние ошибки.
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

// remove - Функция примает на вход массив s []int и индекс элемента, который необходимо удалить. Возвращает новый массив
// без указанного элемента.
func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
