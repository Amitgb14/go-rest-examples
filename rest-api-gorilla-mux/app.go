package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

type User struct {
	ID   int
	Name string
	Age  int
}

var users map[int]User

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRouters()
}

func (a *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, a.Router))
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

func init() {
	users = map[int]User{
		1: User{ID: 1, Name: "Amit", Age: 28},
		2: User{ID: 2, Name: "Rohit", Age: 30},
	}

}
func getUsers() (map[int]User, error) {
	return users, errors.New("")
}

func getUser(id int) (User, error) {
	if value, ok := users[id]; ok {
		return value, nil
	}
	return User{}, errors.New("")
}

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {

	users, _ := getUsers()
	respondWithJson(w, http.StatusOK, users)

}

func (a *App) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	user, err := getUser(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "User ID not Found")
		return
	}
	respondWithJson(w, http.StatusOK, user)

}

func (a *App) addUser(w http.ResponseWriter, r *http.Request) {
	var user User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	users[user.ID] = User{ID: user.ID, Name: user.Name, Age: user.Age}
	// fmt.Println(user.ID)
	respondWithJson(w, http.StatusCreated, map[string]string{"msg": "Successfully added"})

}

func (a *App) addFile(w http.ResponseWriter, r *http.Request) {

	// type Result struct {
	// 	ID   int    `xml:"ID"`
	// 	Name string `xml:"Name"`
	// 	Age  int    `xml:"Name"`
	// }
	// v := Result{}

	var user User

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := xml.Unmarshal(body, &user); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	users[user.ID] = User{ID: user.ID, Name: user.Name, Age: user.Age}
	// fmt.Println(user.ID)
	respondWithJson(w, http.StatusCreated, map[string]string{"msg": "Successfully added"})

}

func (a *App) initializeRouters() {
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
	a.Router.HandleFunc("/user/{id:[0-9]+}", a.getUser).Methods("GET")
	a.Router.HandleFunc("/user", a.addUser).Methods("POST")
	a.Router.HandleFunc("/file", a.addFile).Methods("POST")
}
