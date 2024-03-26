// ===================== Fasthttp =====================
// package main

// import (
// 	"github.com/fasthttp/router"
// 	"api_test/handler"
// 	"api_test/storage/kv"
// 	"log"
// 	"net/http"
// )

// func main() {

// 	kv.Init(kv.NewInMemory())
// 	rtr := router.New()

// 	rtr.GET("/users", handler.GetUser)
// 	rtr.POST("/users/add", handler.AddUser)
// 	rtr.POST("/users/delete", handler.DeleteUser)

// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }
// ====================================================





// ===================== Gin =====================
package main

import (
	"github.com/gin-gonic/gin"
	"api_test/handler"
	"api_test/storage/kv"
)

func main() {
	r := gin.Default()

	kv.Init(kv.NewInMemory())
	r.POST("/user/create", handler.AddUser)
	r.GET("/user/get", handler.GetUser)
	r.PUT("/user/update", handler.UpdateUser)
	r.GET("/users/get", handler.GetAllUsers)
	r.DELETE("/user/delete", handler.DeleteUser)

	r.Run(":8000")
}
