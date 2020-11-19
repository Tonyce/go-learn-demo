package controller

import (
	"logical-example/internal/apiserver/service"
	"logical-example/internal/model"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func init() {
	gin.SetMode(gin.TestMode)
	// common.InitMongoDB()
}

func TestUserV1(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockServiceInstance := service.NewMockUserService(ctrl)
	// userMockRepo := model.NewMockUserRepository(ctrl)

	// userMockRepo.EXPECT().FindOne(1).Return(&model.User{Name: "张三李四"}, nil)
	userMockServiceInstance.EXPECT().GetUser("1", model.MongoUserRepository).Return(&model.User{Name: "张三李四"}, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "k", Value: "test"}}

	// c.Set("UserRepository", userMockRepo)
	c.Set("UserInstance", userMockServiceInstance)

	getUserEndpoint(c)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, `{"name":"test"}`, w.Body.String())
	assert.Equal(t, `{"name":"张三李四"}`, w.Body.String())
}

func TestUserHTTP(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockServiceInstance := service.NewMockUserService(ctrl)
	// userMockRepo := model.NewMockUserRepository(ctrl)

	// userMockRepo.EXPECT().FindOne(1).Return(&model.User{Name: "张三李四"}, nil)
	userMockServiceInstance.EXPECT().GetUserFromHTTP().Return("GetUserFromHTTP 张三李四")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "k", Value: "test"}}

	c.Set("UserInstance", userMockServiceInstance)

	getUserHTTPEndpoint(c)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, `{"name":"test"}`, w.Body.String())
	assert.Equal(t, `GetUserFromHTTP 张三李四`, w.Body.String())
}
