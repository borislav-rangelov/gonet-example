package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/borislav-rangelov/gonet-example/handlers/home"
	"github.com/borislav-rangelov/gonet-example/handlers/users"
	"github.com/borislav-rangelov/gonet/handlers"
	"github.com/borislav-rangelov/gonet/util"
	muxhandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dbType := flag.String("dbtype", "", "Database driver type. Default: mysql")
	dbString := flag.String("dbstring", "root@/test", "DB connection string. Default: root@/test")
	flag.Parse()

	_, err := sql.Open(*dbType, *dbString)

	if err != nil {
		log.Fatal(err)
		os.Exit(2)
		return
	}

	// result, err := db.Exec("CREATE TABLE IF NOT EXISTS `test`.`t1` ( `col` VARCHAR(16) NOT NULL )")

	// if err != nil {
	// 	log.Fatal(err)
	// 	os.Exit(2)
	// 	return
	// }

	config := util.Config{}
	config.Read("config.json")

	log.Println("Setting up router...")
	router := mux.NewRouter()

	configureRouter(router)

	handler := configureMiddleware(router)

	http.Handle("/", handler)

	addr := ":" + strconv.FormatInt(config.GetInt("server.port", 8080), 10)
	log.Printf("Listening on %s", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func configureRouter(router *mux.Router) {
	users.ConfigureRouter(router.PathPrefix("/users").Subrouter())
	home.ConfigureRouter(router.PathPrefix("/").Subrouter())
	router.NotFoundHandler = &handlers.NotFoundHandler{}
}

func configureMiddleware(router *mux.Router) http.Handler {
	handler := muxhandlers.LoggingHandler(os.Stdout, router)

	handler = muxhandlers.CORS(
		muxhandlers.AllowedOriginValidator(func(origin string) bool {
			result, _ := regexp.MatchString("[a-z0-9]\\.website.com", origin)
			return result
		}))(handler)

	return handler
}
