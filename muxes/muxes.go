package muxes

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strings"
	"strconv"
	"log"
)

type (
	User struct {
		Id      int
		Name    string
		Surname string
		Age     int
	}
)

type Users = []User

var users = Users{
	User{
		0,
		"Fatih",
		"Totrakanlı",
		27,
	},
}

func SERVE() *http.ServeMux {
	log.Print("Server started at http://127.0.0.1:3000 port.")
	mux := http.NewServeMux()

	mux.HandleFunc("/newUser", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			http.NotFound(w, req)
			return
		}
		var newUser = convertRequestToUser(req)

		okStatus(w)
		users = append(users, newUser)
		log.Printf("New User %s %s added successfully.", newUser.Name, newUser.Surname)
		json.NewEncoder(w).Encode(newUser)

		return
	})

	mux.HandleFunc("/getAll", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			http.NotFound(w, req)
			return
		}

		okStatus(w)
		log.Printf("All users listed successfully.")
		json.NewEncoder(w).Encode(users)

		return
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, req *http.Request) {
		var method = req.Method
		if method != "GET" && method != "DELETE" && method != "PUT" {
			http.NotFound(w, req)
			return
		}

		id, err := strconv.Atoi(strings.TrimPrefix(req.URL.Path, "/users/"))
		if err != nil {
			panic(err)
		}

		okStatus(w)

		for index, user := range users {
			if user.Id == id {
				if method == "GET" {
					log.Printf("The user who is name %s %s listed successfully.", user.Name, user.Surname)
					json.NewEncoder(w).Encode(user)
					return
				}

				if method == "DELETE" {
					users = remove(users, index)
					log.Printf("The user who is name %s %s deleted successfully.", user.Name, user.Surname)
					json.NewEncoder(w).Encode(users)
					return
				}

				if method == "PUT" {
					var newUser = convertRequestToUser(req)
					users[index] = newUser
					log.Printf("The user who is name %s %s updated successfully.", user.Name, user.Surname)
					json.NewEncoder(w).Encode(users)
					return
				}

			}
		}

		json.NewEncoder(w).Encode(nil)
		return
	})

	return mux
}

func okStatus(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	return
}

func remove(slice Users, s int) Users {
	return append(slice[:s], slice[s+1:]...)
}

func convertRequestToUser(req *http.Request) User {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var newUser User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		panic(err)
	}

	return newUser
}
