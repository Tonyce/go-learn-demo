package service

import (
	"io/ioutil"
	"log"
	"logical-example/internal/model"
	"net/http"
	"time"
)

// UserService ...
type UserService interface {
	GetUser(v string, repo model.UserRepository) (*model.User, error)
	GetUserFromHTTP() string
}

// User ...
type User struct{}

// UserServiceInstance ...
var UserServiceInstance UserService

func init() {
	UserServiceInstance = &User{}
}

// GetUser ...
func (u *User) GetUser(v string, repo model.UserRepository) (*model.User, error) {

	user, err := repo.FindOne(1)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	return user, nil
}

// GetUserFromHTTP ...
func (u *User) GetUserFromHTTP() string {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Get("http://127.0.0.1:8989")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	// log.Printf("%s", )
	return string(body)
}
