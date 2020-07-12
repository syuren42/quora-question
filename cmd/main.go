package main

import (
	crand "crypto/rand"
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math"
	"math/big"
	"math/rand"
	"net/http"
)

type Question struct {
	ID      int
	Content string
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", Index)
	router.Run()
}

func Index(c *gin.Context) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/app_development")
	log.Println("Connected to mysql.")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var count = 0
	err = db.QueryRow("SELECT count(*) FROM questions").Scan(&count)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("レコードが存在しません")
	case err != nil:
		panic(err.Error())
	default:
		fmt.Println(count)
	}

	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	var question_id = rand.Intn(count) + 1

	var question Question
	err = db.QueryRow("SELECT * FROM questions WHERE id = ?", question_id).Scan(&question.ID, &question.Content)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("レコードが存在しません")
	case err != nil:
		panic(err.Error())
	default:
		fmt.Println(question.ID, question.Content)
	}

	c.JSON(http.StatusOK, gin.H{"message": question.Content})
}
