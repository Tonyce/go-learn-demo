package service

import (
	"log"
	"logical-example/internal/model"
	"logical-example/internal/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockUserRepository(ctrl)
	// 期望FindOne(1)返回张三用户
	repo.EXPECT().FindOne(1).Return(&model.User{Name: "张三"}, nil)

	user, _ := UserServiceInstance.GetUser("c", repo)

	log.Println(user)

	assert.Equal(t, "张三", user.Name)
	// assert.Equal(t, `{"name":"test"}`, w.Body.String())
}
