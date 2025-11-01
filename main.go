package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		start := time.Now()

		city := c.Query("city")
		limit := c.Query("limit")
		offset := c.Query("offset")

		query := `
			SELECT users.id, users.name, users.city,
			       COALESCE(COUNT(orders.id), 0) AS total_orders
			FROM users
			LEFT JOIN orders ON users.id = orders.user_id
		`

		var args []interface{}

		if city != "" {
			query += " WHERE users.city = ?"
			args = append(args, city)
		}

		query += `
			GROUP BY users.id
			ORDER BY total_orders DESC, users.id DESC
		`

		if limit != "" {
			query += " LIMIT ?"
			args = append(args, limit)
		}

		if offset != "" {
			query += " OFFSET ?"
			args = append(args, offset)
		}

		rows, err := db.Query(query, args...)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		type User struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			City        string `json:"city"`
			TotalOrders int    `json:"total_orders"`
		}

		users := []User{}
		for rows.Next() {
			var u User
			rows.Scan(&u.ID, &u.Name, &u.City, &u.TotalOrders)
			users = append(users, u)
		}

		elapsed := time.Since(start)
		c.Header("X-Query-Time", fmt.Sprintf("%v", elapsed))
		c.JSON(http.StatusOK, users)
	})

	fmt.Println("Server running on :8080")
	r.Run(":8080")
}
