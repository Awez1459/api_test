package handler

// import (
// 	"encoding/json"
// 	"github.com/google/uuid"
// 	"github.com/valyala/fasthttp"
// 	"api_test/storage"
// 	"api_test/storage/kv"
// 	"net/http"
// )

// func GetUser(ctx *fasthttp.RequestCtx) {
// 	userID := ctx.QueryArgs().Peek("id")
// 	if user, err := kv.Get(string(userID)); err != nil {
// 		ctx.Response.SetStatusCode(http.StatusNotFound)
// 		ctx.Response.SetBody([]byte(err.Error()))
// 		return
// 	} else {
// 		ctx.SetStatusCode(http.StatusOK)
// 		rbyte, _ := json.Marshal(user)
// 		ctx.Response.SetBody(rbyte)
// 	}
// }

// func AddUser(ctx *fasthttp.RequestCtx) {
// 	// Parse the request body into a User struct
// 	var newUser *storage.User
// 	err := json.Unmarshal(ctx.PostBody(), &newUser)
// 	if err != nil {
// 		ctx.SetStatusCode(http.StatusBadRequest)
// 		return
// 	}

// 	newUser.ID = uuid.NewString()

// 	if err := kv.Set(newUser.ID, newUser); err != nil {
// 		ctx.Response.SetStatusCode(500)
// 		return
// 	}

// 	// Return the newly created user
// 	respBody, _ := json.Marshal(&newUser)
// 	ctx.Response.SetBody(respBody)
// 	ctx.Response.SetStatusCode(200)
// }

// func DeleteUser(ctx *fasthttp.RequestCtx) {
// 	userID := ctx.QueryArgs().Peek("id")
// 	if err := kv.Delete(string(userID)); err != nil {
// 		ctx.Response.SetStatusCode(500)
// 		return
// 	} else {
// 		ctx.Response.SetStatusCode(http.StatusOK)
// 	}
// }
