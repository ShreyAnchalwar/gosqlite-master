package main

import (
	"flag"
	"gosqlite/internal/entity"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func dbConnect() (*gorm.DB, error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address") //returns a pointer

	db, err := dbConnect()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entity.User{})

	app := &App{

		UserModel: &entity.UserModel{Db: db},
	}

	srv := &http.Server{
		Addr: *addr,

		Handler: app.registerroutes(),
	}

	log.Printf("Starting server on %s", *addr) // dereferencing the pointer
	// Call the ListenAndServe() method on the new http.Server struct.
	err = srv.ListenAndServe()
	log.Fatal(err)

}
