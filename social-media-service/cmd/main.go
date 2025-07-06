package main

import (
	"log"
	"net/http"
	"os"
	"social-media-service/config"

	"social-media-service/internal/handler"
	"social-media-service/internal/model"
	"social-media-service/internal/repository"
	"social-media-service/internal/router"
	"social-media-service/internal/service"
	"social-media-service/pkg/db"
)

// go func() {
//     lis, _ := net.Listen("tcp", ":50051")
//     s := grpc.NewServer()
//     user.RegisterUserServiceServer(s, &userServer{})
//     if err := s.Serve(lis); err != nil {
//         log.Fatalf("failed to serve gRPC: %v", err)
//     }
// }()

func main() {
	config.LoadEnv()
	db.Connect()

	// Auto-migrate models
	db.DB.AutoMigrate(&model.User{})

	// Initialize components
	repo := repository.NewUserRepository()
	svc := service.NewUserService(repo)
	h := &handler.UserHandler{Service: svc}
	r := router.NewRouter(h)

	port := os.Getenv("PORT")
	log.Println("Server starting on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
