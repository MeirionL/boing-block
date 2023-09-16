package main

import (
	"net/http"

	"github.com/MeirionL/personal-finance-app/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		apiKey, err := auth.GetAPIKey(r.Header)
// 		if err != nil {
// 			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
// 			return
// 		}

// 		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
// 		if err != nil {
// 			respondWithError(w, 404, fmt.Sprintf("Couldn't get user: %v", err))
// 			return
// 		}

// 		handler(w, r, user)
// 	}
// }