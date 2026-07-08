package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/controller"
	"github.com/RafaelFleitas/API-Golang/src/controller/model/response"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUserController_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fakeService := &FakeUserDomainService{
		CreateUserServiceFunc: func(u model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
			u.SetID(1)
			return u, nil
		},
	}

	userController := controller.NewUserControllerInterface(fakeService)

	body := `{"email":"novo@gmail.com","password":"senha123!","name":"Novo Usuario","age":30}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/createUser", bytes.NewBufferString(body))
	c.Request.Header.Set("Content-Type", "application/json")

	userController.CreateUser(c)

	require.Equal(t, http.StatusOK, w.Code)

	var responseBody response.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	require.NoError(t, err)
	assert.Equal(t, int64(1), responseBody.ID)
	assert.Equal(t, "novo@gmail.com", responseBody.Email)
}

func TestCreateUserController_ValidationError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	serviceWasCalled := false
	fakeService := &FakeUserDomainService{
		CreateUserServiceFunc: func(u model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
			serviceWasCalled = true
			return u, nil
		},
	}

	userController := controller.NewUserControllerInterface(fakeService)

	body := `{"email":"nao-e-email","password":"123","name":"No","age":30}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/createUser", bytes.NewBufferString(body))
	c.Request.Header.Set("Content-Type", "application/json")

	userController.CreateUser(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.False(t, serviceWasCalled, "a service não deveria ser chamada quando a validação falha")
}
