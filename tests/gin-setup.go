package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"api_test/storage/kv"
)

func SetupMinimumInstance(path string) error {

	_ = path
	cnf := viper.New()
	cnf.Set("mode", "test")

	kv.Init(kv.NewInMemory())
	return nil
}

func Serve(handler gin.HandlerFunc, uri string, c *gin.Context) error {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	switch uri {
	case "/user/create":
		r.POST(uri, handler)
	case "/user/get":
		r.GET(uri, handler)
	case "/user/update":
		r.PUT(uri, handler)
	case "/users/get":
		r.GET(uri, handler)
	case "/user/delete":
		r.DELETE(uri, handler)
	}

	r.ServeHTTP(c.Writer, c.Request)
	return nil
}
