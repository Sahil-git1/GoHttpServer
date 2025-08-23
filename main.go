package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"net/http"
	"github.com/go-chi/cors"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the environmet")
	}

    router := chi.NewRouter()
    srv := &http.Server{
		Handler : router,
		Addr : ":" + portString,
	}
    
	v1Router := chi.NewRouter()

    router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET","POST","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge:             300,
	}))

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	router.Mount("/v1",v1Router)

    
	log.Printf("Server starting at port %v",portString)
	err := srv.ListenAndServe()
    if err != nil {
		log.Fatal(err)
	}
	
}