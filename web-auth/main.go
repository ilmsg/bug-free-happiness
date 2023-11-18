package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/shaj13/libcache"
	_ "github.com/shaj13/libcache/fifo"

	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/basic"
	"github.com/shaj13/go-guardian/v2/auth/strategies/token"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
)

var strategy union.Union
var tokenStrategy auth.Strategy
var cacheObj libcache.Cache

func setupGoGuardian() {
	cacheObj = libcache.FIFO.New(0)
	cacheObj.SetTTL(time.Minute * 5)

	basicStrategy := basic.NewCached(validateUser, cacheObj)
	tokenStrategy = token.New(token.NoOpAuthenticate, cacheObj)
	strategy = union.New(tokenStrategy, basicStrategy)
}

func main() {
	log.Println("Web Auth !!")
	setupGoGuardian()

	router := mux.NewRouter()
	router.HandleFunc("/v1/auth/token", middleware(http.HandlerFunc(createToken))).Methods("GET")
	router.HandleFunc("/v1/book/{id}", middleware(http.HandlerFunc(getBookAuthor))).Methods("GET")

	log.Println("server startd and listening on :7001")
	http.ListenAndServe(":7001", router)
}

func createToken(w http.ResponseWriter, r *http.Request) {
	token := uuid.New().String()
	user := auth.User(r)
	auth.Append(tokenStrategy, token, user)

	body := fmt.Sprintf("token: %s\n", token)
	w.Write([]byte(body))
}

func getBookAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	books := map[string]string{
		"1449311601": "Ryan Boyd",
		"148425094X": "Yvonne Wilson",
		"1484220498": "Prabath Siriwarden",
	}
	body := fmt.Sprintf("Author: %s\n", books[id])
	w.Write([]byte(body))
}

func validateUser(ctx context.Context, r *http.Request, username, password string) (auth.Info, error) {
	if username == "scott" && password == "tiger" {
		return auth.NewDefaultUser("scott", "1", nil, nil), nil
	}
	return nil, fmt.Errorf("invalid credentials")
}

func middleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing Auth Middleware")

		_, user, err := strategy.AuthenticateRequest(r)
		if err != nil {
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}

		log.Printf("User %s Authenticated\n", user.GetUserName())
		r = auth.RequestWithUser(user, r)
		next.ServeHTTP(w, r)
	})
}
