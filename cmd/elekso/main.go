package main
//http://coderbox.herokuapp.com/

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"database/sql"

	
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)
var (
    repeat int
    db     *sql.DB = nil
)

func main() {

    var errd error
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	    db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if errd != nil {
        log.Fatalf("Error opening database: %q", errd)
    }
	
	//-----------------------------------------
	
	 router.GET("/test", func(c *gin.Context) {
        c.String(http.StatusOK, string("google mandar shinde github"))
    })

	
	//-----------------------------------------
	
   router.GET("/init", func(c *gin.Context) {

     if _, err := db.Exec("CREATE TABLE IF NOT EXISTS pothole(dtime text,lat text,lon text , userid text)"); err != nil {
      c.String(http.StatusInternalServerError,
          fmt.Sprintf("Error creating database table: %q", err))
      return
  }

        if _, err := db.Exec("CREATE TABLE IF NOT EXISTS potholetrail(dtime text,data text , userid text)"); err != nil {
     c.String(http.StatusInternalServerError, fmt.Sprintf("Error creating database table: %q", err))
     return
 }

  })
	
	

	
	//-----------------------------------------
	
   router.GET("/pothole", func(c *gin.Context) {
     usr := c.Query("userid")
     lat := c.Query("lat")
     lon := c.Query("lon")
     dtime := c.Query("dtime")

//	 //https://elekso.herokuapp.com/bump?dtime=10-27-2016 23:27:43&lat=18.6332967&lon=73.8114083&userid=758236590
//	 
//     if _, err := db.Exec("CREATE TABLE IF NOT EXISTS pothole(dtime text,lat text,lon text , userid text)"); err != nil {
//      c.String(http.StatusInternalServerError,
//          fmt.Sprintf("Error creating database table: %q", err))
//      return
//  }

      if _, err :=  db.Exec(  "INSERT INTO pothole (dtime, userid,lat,lon) VALUES ('"+dtime+"','"+usr+"','"+lat+"','"+lon+"')"  ); err != nil {
      c.String(http.StatusInternalServerError,
          fmt.Sprintf("Error ins: %q", err))
      return
  }

      c.String(http.StatusOK, "INS %s  %s %s+%s", usr,dtime,lat,lon)
  })
	
	
	//-----------------------------------------
	type Bump_path_struct struct {
    P_User     string `json:"userid" binding:"required"`
    P_Data string `json:"datapath" binding:"required"`
    P_Time string `json:"dtime" binding:"required"`
	}

	router.POST("/potholetrail", func (c *gin.Context) {
   var json Bump_path_struct
   c.Bind(&json)
  
//  //{"userid": "mandar", "datapath": "123", "dtime": "12-04-2016 05:00:00"}
//  
//      if _, err := db.Exec("CREATE TABLE IF NOT EXISTS potholetrail(dtime text,data text , userid text)"); err != nil {
//     //c.String(http.StatusInternalServerError, fmt.Sprintf("Error creating database table: %q", err))
//     //return
// }

     if _, err :=  db.Exec(  "INSERT INTO potholetrail (dtime, userid,data) VALUES ('"+json.P_Time+"','"+json.P_User+"','"+json.P_Data+"')"  ); err != nil {
     //c.String(http.StatusInternalServerError, fmt.Sprintf("Error ins: %q", err))
     //return
 }

    c.String(http.StatusOK, "name: %s; data: %s time: %s",  json.P_User , json.P_Data,json.P_Time)
    // c.JSON(http.StatusOK, gin.H{"status": "register done"})
	
 })
	
	
	
	
	
//	router.Run(":8080")
	router.Run(":" + port)
}