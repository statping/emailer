package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "emailer.db")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})

	InitConfig()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/request", RequestHandler).Methods("POST")
	r.HandleFunc("/confirm/{id}", ConfirmHandler).Methods("GET")
	r.HandleFunc("/check", CheckHandler).Methods("GET")
	r.HandleFunc("/resend", ResendHandler).Methods("GET")
	r.HandleFunc("/unsubscribe", UnsubscribeHandler).Methods("GET")
	r.HandleFunc("/send", SendHandler).Methods("POST")
	http.ListenAndServe(":8000", r)
}

var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString generates a random string of n length
func RandomString(n int) string {
	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}
