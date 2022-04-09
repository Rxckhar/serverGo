package main

import(
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func pgDataBase() (con *pg.DB) {
	address := fmt.Sprintf("%s:%s", "localhost", "5432")
	options := &pg.Options{
		User: "postgres",
		Password: "",
		Addr: address,
		Database: "notice",
		PoolSize: 50,

	}

	con = pg.Connect(options)
	if con == nil {
		log.Fatalf("Не удалось подключиться к postgress")
	}
	return
}

func Api(c *gin.Context) {
	db := pgDataBase()

	db.Close()

	c.JSON(200, gin.H{
		"api": "notice",

	})
}

func main() {
	r := gin.Default()
	r.GET("/api", Api)
	r.Run("0.0.0.0:9090")
}