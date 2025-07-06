package router

import (
	"net/http"
	"social-media-service/internal/handler"
	"social-media-service/internal/middleware"

	"github.com/rs/cors"
)

func NewRouter(h *handler.UserHandler) http.Handler {
	mux := http.NewServeMux()

	// Define handlers without CORS
	mux.Handle("/register", http.HandlerFunc(h.Register))
	mux.Handle("/login", http.HandlerFunc(h.Login))
	mux.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(h.Profile)))
	mux.Handle("/posts", middleware.AuthMiddleware(http.HandlerFunc(handler.CreatePost)))

	// Wrap the entire mux with CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	handlerWithCORS := c.Handler(mux)

	return handlerWithCORS
}
