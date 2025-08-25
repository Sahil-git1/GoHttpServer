package main

import (
	"net/http"
	"github.com/Sahil-git1/GoHttpServer/internal/database"
	"github.com/Sahil-git1/GoHttpServer/internal/auth"
	"fmt"
)
type authedHandler func(w http.ResponseWriter, r *http.Request,user database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
			apiKey, err := auth.GetAPIKey(r.Header)
	if err !=nil {
		respondWithError(w, 403 , fmt.Sprintf("Auth error: %v",err))
		return
	}
	user , err := cfg.DB.GetUserByAPIKey(r.Context(),apiKey)
		if err !=nil {
		respondWithError(w, 400 , fmt.Sprintf("Couldn't get user: %v",err))
		return
	}
	handler(w,r,user)
	}
}