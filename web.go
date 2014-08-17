package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/bmizerany/pq"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	// "github.com/jmoiron/sqlx"
)

var global_img_cache = make(map[string][]map[string]interface{})
var db *sql.DB

// type SearchTerm struct {
// 	id   int
// 	term string
// }

func main() {
	r := gin.Default()
	// r.Use(InitDBHandler())
	db = dbConnect()

	defer db.Close()
	r.GET("/", root)
	r.GET("/ping", ping)
	r.GET("/links/:title", links)
	r.GET("/terms", terms)
	r.POST("/addsearchterm/*term", addsearchterm)

	var port = os.Getenv("PORT")
	r.Run(":" + port)
}

func root(c *gin.Context) {
	c.File("root.html")
}

func addsearchterm(c *gin.Context) {
	new_term := c.Params.ByName("term")
	if new_term != "" {
		fmt.Println("processing new term", new_term)
	} else {
		fmt.Println("skipping it since string is nil")
	}
	c.String(200, "OK")
}

func ping(c *gin.Context) {
	fmt.Println("processing ping request")
	c.JSON(200, gin.H{"ping": "pong"})
}

func terms(c *gin.Context) {
	terms := GetAllSearchTerms()
	terms_maps := make([]map[string]interface{}, 0)
	for _, term := range terms {
		terms_maps = append(terms_maps, term.ToMap())
	}
	c.JSON(200, terms_maps)
}

func dbConnect() *sql.DB {
	database_url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(database_url)
	conn, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}

	return conn
}

func links(c *gin.Context) {
	title := c.Params.ByName("title")
	var data []map[string]interface{}

	if _, key_exists_in_cache := global_img_cache[title]; key_exists_in_cache {
		fmt.Println("using cache")
		data = global_img_cache[title]
	} else {
		fmt.Println("cache miss. going to db")
		s := []string{"http://www.myapifilms.com/imdb?title=", title, "&exactFilter=0&limit=10"}
		fmt.Println(strings.Join(s, ""))
		resp, _ := http.Get(strings.Join(s, ""))

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		if err := json.Unmarshal(body, &data); err != nil {
			panic(err)
		}

		global_img_cache[title] = data
	}
	c.JSON(200, data)
}
