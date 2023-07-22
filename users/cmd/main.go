package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	rpc "github.com/intxff/mango/common/protos/users"
	"github.com/intxff/mango/users/internal/handlers"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, _ := cfg.Build()
	defer logger.Sync()

	link := "root:root321@(127.0.0.1:10000)/tudou_users?charset=utf8mb4&parseTime=True&loc=Local"
	gormdb, _ := gorm.Open(mysql.Open(link))

	addr := fmt.Sprintf(":%d", 50051)
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Cannot listen to address %s", addr)
	}
	s := grpc.NewServer()
	rpc.RegisterUsersServer(s, &handlers.UsersServer{WriteDB: gormdb, ReadDB: gormdb, Logger: *logger})
	reflection.Register(s)
	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
