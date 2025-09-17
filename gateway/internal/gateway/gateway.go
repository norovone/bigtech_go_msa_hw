package gateway

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authPB "github.com/norovone/bigtech_go_msa_hw/auth/api/gen/proto"
	chatPB "github.com/norovone/bigtech_go_msa_hw/chat/api/gen/proto"
	socialPB "github.com/norovone/bigtech_go_msa_hw/social/api/gen/proto"
	usersPB "github.com/norovone/bigtech_go_msa_hw/users/api/gen/proto"
)

func Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Регистрируем все gRPC-сервисы
	err := authPB.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "auth_service:8082", opts)
	if err != nil {
		return err
	}

	err = usersPB.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "user_service:8083", opts)
	if err != nil {
		return err
	}

	err = socialPB.RegisterSocialServiceHandlerFromEndpoint(ctx, mux, "social_service:8084", opts)
	if err != nil {
		return err
	}

	err = chatPB.RegisterChatServiceHandlerFromEndpoint(ctx, mux, "chat_service:8085", opts)
	if err != nil {
		return err
	}

	// Swagger UI
	swaggerMux := http.NewServeMux()
	swaggerMux.Handle("/", http.FileServer(http.Dir("./swagger-ui")))
	swaggerMux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger-ui"))))

	// Объединяем маршруты
	mainMux := http.NewServeMux()
	mainMux.Handle("/v1/", mux)
	mainMux.Handle("/swagger/", swaggerMux)
	mainMux.Handle("/", swaggerMux) // главная страница — Swagger UI

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Gateway listening on :%s", port)
	return http.ListenAndServe(":"+port, mainMux)
}
