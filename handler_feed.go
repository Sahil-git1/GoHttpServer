package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	
	"github.com/Sahil-git1/GoHttpServer/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User){
	
	type parameters struct{
		Name string `json:"name"`
		URL string `json:"url"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
    
	err := decoder.Decode(&params) 
    if err !=nil{
		respondWithError(w , 400 , fmt.Sprintf("Error parsing json: %v",err) )
	    return
	}
    feed, err := apiCfg.DB.CreateFeeds(
		r.Context(), database.CreateFeedsParams{
        ID : uuid.New(),
		CreatedAt : time.Now().UTC(),
		UpdatedAt : time.Now().UTC(),
		Name : params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
    if err !=nil{
		respondWithError(w , 400 , fmt.Sprintf("Couldn't create the feed: %v",err) )
	    return
	}	
	respondsWithJson(w , 201 , databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request){
	
    feeds, err := apiCfg.DB.GetFeeds(
		r.Context())
    if err !=nil{
		respondWithError(w , 400 , fmt.Sprintf("Couldn't get the feed: %v",err) )
	    return
	}	
	respondsWithJson(w , 201 , databaseFeedsToFeeds(feeds))
}
