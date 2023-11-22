package launch

import (
	"log"
	"net/http"
	"time"
	"xcluster/internal/database"
	"xcluster/internal/router"
	"xcluster/internal/server"
	"xcluster/internal/session"
	"xcluster/internal/user"
	"xcluster/pkg/redis"
)

func Server() {
	// init redis
	configRedis := redis.NewConfig("172.19.80.201:6379", "")
	session.InitStore(redis.NewRedis(configRedis))
	log.Println("redis: init success")
	// init db
	configDB := database.NewConfig("127.0.0.1:3306", "root", "root", "database_test")
	var err error
	if err = database.InitDatabase(configDB, true); err != nil {
		panic(err)
	}
	log.Println("database: init success")
	// init module
	if err = user.InitUserTable(); err != nil {
		panic(err)
	}
	if err = user.InitUserSessionTable(); err != nil {
		panic(err)
	}
	log.Println("user: init success")
	if err = server.InitServerTable(); err != nil {
		panic(err)
	}
	if err = server.InitGroupTable(); err != nil {
		panic(err)
	}
	log.Println("server: init success")
	// init server
	r := router.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("server listen at", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
