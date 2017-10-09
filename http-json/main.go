package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	_ "net/http/pprof"
	"net/mail"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6061", nil))
	}()

	http.HandleFunc("/", CreateUser)
	log.Println(http.ListenAndServe(":60001", nil))
}

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Id      string `json:"id"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	decoder.Decode(&user)
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	validationErr := validate(user)
	if validationErr != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    500,
			Message: validationErr.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(Response{
		Code: 200,
		Id:   "1000000",
	})
}

func validate(in User) error {
	_, err := mail.ParseAddress(in.Email)
	if err != nil {
		return err
	}

	if len(in.Name) < 4 {
		return errors.New("Name is too short")
	}

	if len(in.Password) < 4 {
		return errors.New("Password is too weak")
	}

	return nil
}
