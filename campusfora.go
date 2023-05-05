package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

type posts struct {
	Id   int64
	Text string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgresyash"
	dbname   = "campusfora_task"
)

func main() {
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	r := gin.Default()
	r.GET("/fetch", func(c *gin.Context) {
		getStmt := `select * from campusfora`
		res, e := db.Query(getStmt)
		CheckError(e)
		defer res.Close()
		var data posts
		var finalData []posts
		for res.Next() {
			err := res.Scan(&data.Id, &data.Text)
			if err != nil {
				log.Fatal(err)
			}
			finalData = append(finalData, data)
		}
		c.JSON(http.StatusOK, finalData)
	})

	r.POST("/post", func(c *gin.Context) {

		id := c.PostForm("id")
		i, e := strconv.ParseInt(id, 10, 64)
		CheckError(e)
		//fmt.Printf("id = %v", i)
		text := c.PostForm("text")
		//fmt.Printf("id: %d; post: %s", i, text)
		insertDynStmt := `insert into campusfora(id,textarea) values($1,$2)`
		_, e = db.Exec(insertDynStmt, i, text)
		CheckError(e)
	})

	r.DELETE("/delete", func(c *gin.Context) {
		id := c.PostForm("id")
		i, e := strconv.ParseInt(id, 10, 64)
		CheckError(e)
		deleteStmt := `delete from campusfora where id= $1`
		_, e = db.Exec(deleteStmt, i)
		CheckError(e)
	})

	r.PUT("/update", func(c *gin.Context) {
		id := c.PostForm("id")
		i, e := strconv.ParseInt(id, 10, 64)
		CheckError(e)
		text := c.PostForm("text")
		updateStmt := `update campusfora set textarea = $1 where id = $2`
		_, e = db.Exec(updateStmt, text, i)
		CheckError(e)
	})

	r.Run(":8080")

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
