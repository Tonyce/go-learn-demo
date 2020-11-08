// +build mock

// go test -tag mock 可使用 mock 文件

package service

import (
	"logical-example/internal/model"
)

// GetUser ...
func GetUser(v string) (*model.User, error) {
	return &model.User{
		Name: v,
	}, nil
}
