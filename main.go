package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

type user struct {
	ID          string `json:"ID"`
	UserName       string `json:"UserName"`
	Password string `json:"Password"`
}

type allUsers []user

var users = allUsers{
	{
		ID:	"1",
		UserName: "test",
		Password: "test",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser user
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the UserName and Password only in order to update")
	}
	
	json.Unmarshal(reqBody, &newUser)
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

func getOneUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	for _, singleUser := range users {
		if singleUser.ID == userID {
			json.NewEncoder(w).Encode(singleUser)
		}
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	var updatedUser user

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with UserName and Password only in order to update")
	}
	json.Unmarshal(reqBody, &updatedUser)

	for i, singleUser := range users {
		if singleUser.ID == userID {
			singleUser.UserName = updatedUser.UserName
			singleUser.Password = updatedUser.Password
			users = append(users[:i], singleUser)
			json.NewEncoder(w).Encode(singleUser)
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	for i, singleUser := range users {
		if singleUser.ID == userID {
			users = append(users[:i], users[i+1:]...)
			fmt.Fprintf(w, "The with ID %v has been deleted successfully", userID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/users", getAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getOneUser).Methods("GET")
	router.HandleFunc("/users/{id}", updateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}