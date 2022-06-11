package app

import (
	"github.com/viizgar/go-bookstore_users-api/controllers/ping"
	"github.com/viizgar/go-bookstore_users-api/controllers/user"
)

func mapsUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/search", user.GetUser)
	router.POST("/users", user.CreateUser)
	router.GET("/users/:user_id", user.GetUser)
	router.PUT("/users/:user_id", user.UpdateUser)

}
