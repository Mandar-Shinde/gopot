package main

import (
    "bytes"
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "database/sql"

    "github.com/gin-gonic/gin"

    _ "github.com/lib/pq"
)

var (
    repeat int
    db     *sql.DB = nil
)

func repeatFunc(c *gin.Context) {
    var buffer bytes.Buffer
    for i := 0; i < repeat; i++ {
        buffer.WriteString("Hello from Go!")
    }
    c.String(http.StatusOK, buffer.String())
}

//----------------
// Path location
//---------------
type Bump_path_struct struct {
    P_User     string `json:"userid" binding:"required"`
    P_Data string `json:"datapath" binding:"required"`
    P_Time string `json:"dtime" binding:"required"`
}

// func Bump_path_Func(c *gin.Context) {
//   var json Bump_path_struct
//   c.Bind(&json)
//
//   // if json.User == "man" && json.Password == "123" {
//   //     c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
//   // } else {
//   //     c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
//   // }
//
//
//   fmt.Sprintf("name: %s; data: %s time: %s",  json.P_User , json.P_Data,json.P_Time)
//     c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
// }

//-----------------------

func main() {
    var err error
    var errd error
    port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
    }

    tStr := os.Getenv("REPEAT")
    repeat, err = strconv.Atoi(tStr)
    if err != nil {
        log.Print("Error converting $REPEAT to an int: %q - Using default", err)
        repeat = 5
    }

    db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if errd != nil {
        log.Fatalf("Error opening database: %q", errd)
    }

    router := gin.New()
    router.Use(gin.Logger())
    router.LoadHTMLGlob("templates/*.tmpl.html")
    router.Static("/static", "static")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl.html", nil)
    })

    router.GET("/test", func(c *gin.Context) {
        c.String(http.StatusOK, string("google mandar ... shinde "))
    })


  //router.POST("/bumppath:", Bump_path_Func)

      router.GET("/bump", func(c *gin.Context) {
       usr := c.Query("userid")
       lat := c.Query("lat")
       lon := c.Query("lon")
       dtime := c.Query("dtime")

       if _, err := db.Exec("CREATE TABLE IF NOT EXISTS pothole(dtime text,lat text,lon text , userid text)"); err != nil {
        c.String(http.StatusInternalServerError,
            fmt.Sprintf("Error creating database table: %q", err))
        return
    }

        if _, err :=  db.Exec(  "INSERT INTO pothole (dtime, userid,lat,lon) VALUES ('"+dtime+"','"+usr+"','"+lat+"','"+lon+"')"  ); err != nil {
        c.String(http.StatusInternalServerError,
            fmt.Sprintf("Error ins: %q", err))
        return
    }

        c.String(http.StatusOK, "INS %s  %s %s+%s", usr,dtime,lat,lon)
    })
    router.Run(":8080")

  router.Run(":" + port)
}
