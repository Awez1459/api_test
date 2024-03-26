// ====================== Fasthttp ===================
// package tests

// import (
// 	"encoding/json"
// 	"github.com/k0kubun/pp"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// 	"github.com/valyala/fasthttp"
// 	"api_test/handler"
// 	"api_test/storage"
// 	"testing"
// )

// func TestApi(t *testing.T) {
// 	require.NoError(t, SetupMinimumInstance(""))
// 	buf, err := OpenFile("user.json")
// 	require.NoError(t, err)

// 	req := NewRequest(fasthttp.MethodPost, "/user/create", buf)
// 	res := NewResponse()

// 	assert.NoError(t, Serve(handler.AddUser, "/user/create", req, res))
// 	assert.Equal(t, fasthttp.StatusOK, res.StatusCode())
// 	var user *storage.User
// 	require.NoError(t, json.Unmarshal(res.Body(), &user))
// 	require.Equal(t, user.Email, "test@gmail.com")
// 	require.Equal(t, user.Username, "testusername")
// 	require.NotNil(t, user.ID)

// 	getReq := NewRequest(fasthttp.MethodPost, "/user/get", buf)
// 	getRes := NewResponse()
// 	args := getReq.URI().QueryArgs()
// 	args.Set("id", user.ID)
// 	assert.NoError(t, Serve(handler.GetUser, "/user/get", getReq, getRes))
// 	pp.Print(string(getRes.Body()))
// 	assert.Equal(t, fasthttp.StatusOK, res.StatusCode())
// 	var getUser *storage.User
// 	require.NoError(t, json.Unmarshal(res.Body(), &getUser))
// 	assert.Equal(t, user.ID, getUser.ID)
// }
// ==============================================



// ====================== Gin ===================
package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"api_test/handler"
	"api_test/storage"
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
)

func OpenFile(fileName string) ([]byte, error) {
	return os.ReadFile(fileName)
}

func TestApi(t *testing.T) {
    // Setup
    require.NoError(t, SetupMinimumInstance(""))
    buf, err := OpenFile("user.json")
    require.NoError(t, err)

    router := gin.Default()
    router.POST("/user/create", handler.AddUser)
    router.GET("/user/get", handler.GetUser)
    router.PUT("/user/update", handler.UpdateUser)
    router.GET("/users/get", handler.GetAllUsers)
    router.DELETE("/user/delete", handler.DeleteUser)

    // Test AddUser
    addUserReq, err := http.NewRequest("POST", "/user/create", bytes.NewBuffer(buf))
    require.NoError(t, err)
    addUserReq.Header.Set("Content-Type", "application/json")
    addUserRes := httptest.NewRecorder()
    router.ServeHTTP(addUserRes, addUserReq)
    assert.Equal(t, http.StatusOK, addUserRes.Code)

    // Test GetUser
    var user storage.User
    require.NoError(t, json.Unmarshal(addUserRes.Body.Bytes(), &user))
    getUserReq, err := http.NewRequest("GET", "/user/get?id="+user.ID, nil)
    require.NoError(t, err)
    getUserRes := httptest.NewRecorder()
    router.ServeHTTP(getUserRes, getUserReq)
    assert.Equal(t, http.StatusOK, getUserRes.Code)

    // Test UpdateUser
    var newUser storage.User
    require.NoError(t, json.Unmarshal(addUserRes.Body.Bytes(), &newUser))
    updateUserReq, err := http.NewRequest("PUT", "/user/update?id="+user.ID, bytes.NewBuffer(buf))
    require.NoError(t, err)
    updateUserRes := httptest.NewRecorder()
    router.ServeHTTP(updateUserRes, updateUserReq)
    assert.Equal(t, http.StatusOK, updateUserRes.Code)

    // Test GetAllUsers
    getAllUsersReq, err := http.NewRequest("GET", "/users/get", nil)
    require.NoError(t, err)
    getAllUsersRes := httptest.NewRecorder()
    router.ServeHTTP(getAllUsersRes, getAllUsersReq)
    assert.Equal(t, http.StatusOK, getAllUsersRes.Code)
	
    // Test DeleteUser
    deleteUserReq, err := http.NewRequest("DELETE", "/user/delete?id="+user.ID, nil)
    require.NoError(t, err)
    deleteUserRes := httptest.NewRecorder()
    router.ServeHTTP(deleteUserRes, deleteUserReq)
    assert.Equal(t, http.StatusOK, deleteUserRes.Code)
}
