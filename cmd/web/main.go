package main

import (
	"database/sql"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// To run the web-server enter `go run ./cmd/web` from the root directory of the project file

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	database *databaseConn
}

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "admin:Admin123###@/todoapp?parseTime=true", "MYSQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		database: &databaseConn{DB: db},
	}

	// Initialize Gin router
	router := gin.Default()
	// Load all the html files
	router.LoadHTMLGlob("ui/html/*")

	// call the routes function from routes.go to define routes
	app.routes(router)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  router,
	}

	infoLog.Printf("Starting server on localhost%s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
