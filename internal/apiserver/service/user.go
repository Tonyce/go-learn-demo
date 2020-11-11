// +build !mock

package service

import (
	"io/ioutil"
	"log"
	"logical-example/internal/model"
	"net/http"
	"time"
)

// GetUser ...
func GetUser(v string) (*model.User, error) {

	user, err := model.MongoUserRepository.FindOne(1)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	return user, nil
}

// GetUserFromHTTP ...
func GetUserFromHTTP() string {
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
