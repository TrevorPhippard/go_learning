package handler

import (
	"encoding/json"
	"net/http"
	"social-media-service/internal/middleware"
	"social-media-service/internal/model"
	"social-media-service/pkg/db"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	var post model.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.UserID = userID

	if err := db.DB.Create(&post).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(post)
}
