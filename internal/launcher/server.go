package launcher

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"xcluster/internal/database"
	"xcluster/internal/router"
	"xcluster/internal/server"
	"xcluster/internal/session"
	"xcluster/internal/user"
	"xcluster/pkg/redis"
	"xcluster/pkg/utils"
)

type Server struct {
	// configurations
	srv            *http.Server
	shutdown       bool
	sigShutdown    chan interface{}
	finishShutdown chan interface{}
}

func NewServer() *Server {
	return &Server{
		sigShutdown:    make(chan interface{}),
		finishShutdown: make(chan interface{}),
	}
}

func (s *Server) Launch() {
	// init redis
	configRedis := redis.NewConfig("127.0.0.1:6379", "")
	log.Println("redis: init")
	redisWrapper, err := redis.NewRedisWrapper(configRedis)
	if err != nil {
		panic(err)
	}
	defer utils.CloseAndPrintError(redisWrapper)
	session.InitStore(redisWrapper)
	//var err error
	// init db
	configDB := database.NewConfig("127.0.0.1:3306", "root", "root", "database_test")
	log.Println("database: init")
	if err = database.InitDatabase(configDB, true); err != nil {
		panic(err)
	}
	defer utils.CloseAndPrintError(database.DB)
	// init module
	log.Println("user: init")
	if err = user.InitUserTable(); err != nil {
		panic(err)
	}
	if err = user.InitUserSessionTable(); err != nil {
		panic(err)
	}
	log.Println("server: init")
	if err = server.InitServerTable(); err != nil {
		panic(err)
	}
	if err = server.InitGroupTable(); err != nil {
		panic(err)
	}
	if err = server.InitTokenTable(); err != nil {
		panic(err)
	}
	// init server
	r := router.NewServerRouter()
	s.srv = &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("server listen at %s", s.srv.Addr)
	if err = s.srv.ListenAndServe(); err != nil && !s.shutdown {
		log.Printf("server: listen and serve failed, err=%s", err)
		s.finishShutdown <- nil
		return
	}
	// wait for graceful shutdown
	<-s.sigShutdown
	log.Println("server: shutdown completed gracefully")
	s.finishShutdown <- nil
}

func (s *Server) ListenForShutdown() {
	sys := make(chan os.Signal, 1)
	signal.Notify(sys, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-sys:
		log.Println("server: shutdown signal received")
		s.shutdown = true
		if err := s.srv.Shutdown(context.Background()); err != nil {
			log.Printf("server: shutdown error, err=%s", err)
		}
		s.sigShutdown <- nil
		// wait for resource release
		<-s.finishShutdown
	case <-s.finishShutdown:
	}
}
