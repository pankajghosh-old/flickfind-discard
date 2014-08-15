package main

import (
    "net/http"
    "io/ioutil"
    "github.com/gin-gonic/gin"
    "encoding/json"
    "strings"
    "fmt"
    "os"
)

var global_img_cache = make(map[string][]map[string]interface{});

func main() {
    r := gin.Default()
    r.GET("/", root)
    r.GET("/ping", ping)
    r.GET("/links/:title", links)

    var port = os.Getenv("PORT")
    r.Run(":"+port)
}

func root(c *gin.Context) {
    c.File("root.html")
}

func ping(c *gin.Context) {
    c.JSON(200, gin.H{"ping":"pong"})
}

func links(c *gin.Context) {
    title := c.Params.ByName("title")
    var data []map[string]interface{}

    if _, key_exists_in_cache := global_img_cache[title]; key_exists_in_cache{
        fmt.Println("using cache")
        data = global_img_cache[title]
    }else{
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
