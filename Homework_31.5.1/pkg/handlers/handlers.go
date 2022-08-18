package handlers

import (
	"Homework_31.5.1/pkg/user"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"net/http"
	"strconv"
)

//Store - структура для хранения пользователей.
type Store struct {
	DB *sql.DB
}

//CreateUserView - Метод структуры Store, который обрабатывает POST запрос на создание пользователя.
//Принимает на вход объект, соответствующий интерфейсу http.ResponseWriter и ссылку объект структуры http.Request.
//В ответ отправляет ID созданного пользователя в формате JSON и код состояния HTTP, либо ошибку в текстовом формате
//и код состояния HTTP. Тело запроса должно быть в формате {"name":"some name","age":"24","friends":[]}
func (db *Store) CreateUserView(w http.ResponseWriter, r *http.Request) {
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
			UserId, err := userAddToDatabase(db, &newUser)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			response, _ := json.Marshal(map[string]int{"UserId": UserId})
			w.Write(response)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

//userAddToDatabase - Функция примает ссылку на объект структуры s Store, ссылку объект структуры newUser user.User
//Добавляет нового пользователя в базу данных. Возвращает id нового пользователя в формате int и статус ошибки.
func userAddToDatabase(s *Store, newUser *user.User) (int, error) {
	var userId int
	err := s.DB.QueryRow("insert into Users (user_name, age) values ($1, $2) returning id",
		newUser.Name, newUser.Age).Scan(&userId)
	if err != nil {
		return 0, err
	}
	if len(newUser.Friends) != 0 {
		for _, value := range newUser.Friends {
			s.DB.QueryRow("insert into FriendList (userId, friendId) values ($1, $2)", userId, value)
		}
	}
	return userId, nil
}

//GetFriendsListView - Метод структуры Store, который обрабатывает GET запрос на получение списка друзей пользователя.
//Принимает на вход объект, соответствующий интерфейсу http.ResponseWriter и ссылку объект структуры http.Request.
//В ответ отправляет список друзей пользователя в формате JSON и код состояния HTTP, либо ошибку в текстовом формате
//и код состояния HTTP.
func (s *Store) GetFriendsListView(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		userId, response, err := getFriendList(s, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(map[string][]int{userId: response})
		w.Write(res)

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

//getFriendList - Функция примает ссылку на объект структуры s Store и ссылку объект структуры r http.Request.
//Из URL извлекает параметр, который соответствует ID пользователя. Возвращает ID пользователя в формате string
// список друзей пользователя в формате []int и и состояние ошибки.
func getFriendList(s *Store, r *http.Request) (string, []int, error) {
	idParam := chi.URLParam(r, "UserId")
	idUser, err := strconv.Atoi(idParam)
	if err != nil {
		return "", nil, err
	}
	result := make([]int, 0)
	rows, err := s.DB.Query("select friendId from FriendList where userId = $1", idUser)
	if err != nil {
		return "", nil, err
	}
	for rows.Next() {
		var friendId int
		err := rows.Scan(&friendId)
		if err != nil {
			continue
		}
		result = append(result, friendId)
	}
	return idParam, result, nil
}

//UpdateUserAgeView - Метод структуры Store, который обрабатывает PATCH запрос на изменение возраста пользователя.
//Принимает на вход объект, соответствующий интерфейсу http.ResponseWriter и ссылку объект
//структуры http.Request. В ответ отправляет статус обновления возраста и код состояния HTTP,
//либо ошибку в текстовом формате и код состояния HTTP. Тело запроса должно быть в формате {"new age":"28"}
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

//updateUserAge - Функция примает ссылку на объект структуры s Store, ссылку объект структуры r http.Request
//и новый возраст пользователя age в формате int. Из URL извлекает параметр, который соответствует ID пользователя.
// Далее обновляет запись в БД и возвращает состояние ошибки.
func updateUserAge(s *Store, r *http.Request, age int) error {
	idParam := chi.URLParam(r, "UserId")
	idUser, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}
	var UserName string
	err = s.DB.QueryRow("select user_name from Users where id = $1", idUser).Scan(&UserName)
	if err != nil || UserName == "" {
		return errors.New("Такого пользователя нет!")
	}
	_, err = s.DB.Exec("update Users set age = $1 where id = $2", age, idUser)
	if err != nil {
		return err
	}
	return nil
}

//MakeFriendsView - Метод структуры Store, который обрабатывает POST запрос на добавление пользователей в друзья.
//Принимает на вход объект, соответствующий интерфейсу http.ResponseWriter и ссылку объект структуры http.Request.
//В ответ отправляет статус добавления в друзья и код состояния HTTP, либо ошибку в текстовом формате
//и код состояния HTTP. Тело запроса должно быть в формате {"source_id":"1","target_id":"2"}
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
				w.WriteHeader(http.StatusInternalServerError)
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

//makeFriends - Функция примает ссылку на объект структуры s Store, ссылку объект структуры newFriends user.Friend
//Извлекает из объекта структуры user.Friend ID пользователей добавляет их в друзья друг к другу.
//Возвращает имена пользователей в формате string и состояние ошибки.
func makeFriends(s *Store, newFriends *user.Friend) (string, string, error) {
	var sourseUserName string
	err := s.DB.QueryRow("select user_name from Users where id = $1", newFriends.SourceId).Scan(&sourseUserName)
	if err != nil || sourseUserName == "" {
		return "", "", err
	}
	var targetUserName string
	err = s.DB.QueryRow("select user_name from Users where id = $1", newFriends.TargetId).Scan(&targetUserName)
	if err != nil || targetUserName == "" {
		return "", "", err
	}
	s.DB.QueryRow("insert into FriendList (userId, friendId) values ($1, $2), ($2, $1)", newFriends.SourceId, newFriends.TargetId)
	return sourseUserName, targetUserName, nil
}

//DeleteUserView - Метод структуры Store, который обрабатывает POST запрос на удаление пользователя.
//Принимает на вход объект, соответствующий интерфейсу w http.ResponseWriter и ссылку объект структуры r http.Request.
//В ответ отправляет имя удаленного пользователя в текстовом формате и код состояния HTTP, либо ошибку в текстовом формате
//и код состояния HTTP. Тело запроса должно быть в формате {"target_id":"1"}
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

//deleteUser - Функция примает ссылку на объект структуры s Store, ссылку объект структуры userTarget user.UserId
//Извлекает из объекта структуры user.UserId ID пользователя и удалеет его каскадно из БД.
//Возвращает имя удаленного пользователя в формате string и состояние ошибки.
func deleteUser(s *Store, userTarget *user.UserId) (string, error) {
	var UserName string
	err := s.DB.QueryRow("select user_name from Users where id = $1", userTarget.TargetId).Scan(&UserName)
	if err != nil || UserName == "" {
		return "", errors.New("Такого пользователя нет!")
	}
	_, err = s.DB.Exec("delete from Users where id = $1", userTarget.TargetId)
	if err != nil {
		return "", err
	}
	return UserName, nil
}
