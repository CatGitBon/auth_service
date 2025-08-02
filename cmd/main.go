package main

import (
	"context"
	"fmt"
	"log"
	authService "microservices/auth_service/pkg" // Импортируем сгенерированные файлы
	"net"

	"google.golang.org/grpc"
)

type server struct {
	authService.UnimplementedAuthServiceServer
}

func (s *server) Authenticate(ctx context.Context, req *authService.AuthRequest) (*authService.AuthResponse, error) {
	// Логика аутентификации
	log.Printf(req.GetUserId(), req.GetPassword())
	if req.GetUserId() == "user124" && req.GetPassword() == "password123" {
		return &authService.AuthResponse{
			Success: true,
			Message: "Authentication successful",
		}, nil
	}

	// Если аутентификация не прошла
	return &authService.AuthResponse{
		Success: false,
		Message: "Authentication failed",
	}, nil
}

func main() {
	// Создаем слушателя для gRPC
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Создаем gRPC сервер
	grpcServer := grpc.NewServer()

	// Регистрируем наш AuthService
	authService.RegisterAuthServiceServer(grpcServer, &server{})

	// Запускаем сервер
	fmt.Println("Auth Service listening on port :50051")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
