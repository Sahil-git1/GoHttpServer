package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sahil-git1/GoHttpServer/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User){
	
	type parameters struct{
		 FeedId uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
    
	err := decoder.Decode(&params) 
    if err !=nil{
		respondWithError(w , 400 , fmt.Sprintf("Error parsing json: %v",err) )
	    return
	}
    feedFollow, err := apiCfg.DB.CreateFeedFollows(
		r.Context(), database.CreateFeedFollowsParams{
        ID : uuid.New(),
		CreatedAt : time.Now().UTC(),
		UpdatedAt : time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedId,
	})
    if err !=nil{
		respondWithError(w , 400 , fmt.Sprintf("Couldn't create the feed: %v",err) )
	    return
	}	
	respondsWithJson(w , 201 , databaseFeedFollowToFeedFollow(feedFollow))
}


func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User){
	
	
    feedFollow, err := apiCfg.DB.GetFeedFollows(
		r.Context(), user.ID)
    if err !=nil{
		respondWithError(w , 400 , fmt.Sprintf("Couldn't get the feed follows: %v",err) )
	    return
	}	
	respondsWithJson(w , 201 , databaseFeedsFollowsToFeedsFollows(feedFollow))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User){
	
	feedFollowIDStr := chi.URLParam(r,"feedFollowID")
	feedFollowID ,err := uuid.Parse(feedFollowIDStr)
    if err !=nil{
		respondWithError(w , 400 , fmt.Sprintf("Couldn't parse feed follow ID: %v",err) )
	    return
	}	


    err = apiCfg.DB.DeleteFeedFollows(
		r.Context(), database.DeleteFeedFollowsParams{
			ID: feedFollowID,
			UserID: user.ID,
		})
    if err !=nil{
		respondWithError(w , 400 , fmt.Sprintf("Couldn't delete the feed follows: %v",err) )
	    return
	}	
	respondsWithJson(w , 200 , struct{}{})
}