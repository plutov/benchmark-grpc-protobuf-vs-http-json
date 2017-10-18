package httpjson

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/mail"
)

// Start entrypoint
func Start() {
	http.HandleFunc("/", CreateUser)
	log.Println(http.ListenAndServe(":60001", nil))
}

// User type
type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Response type
type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	ID      string `json:"id"`
}

// CreateUser handler
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
		ID:   "1000000",
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
