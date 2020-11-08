// +build !mock

package service

import (
	"logical-example/internal/model"
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
