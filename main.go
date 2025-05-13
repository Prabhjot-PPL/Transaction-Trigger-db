package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "postgresql://postgres:password@localhost:5433/e_commerce?sslmode=disable"

	// Database connection

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	er := db.Ping()
	if err != nil {
		log.Fatalln("Error connecting to db : ", er)
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Failed to begin transaction: %v", err)
	}

	userID := "828c990f-1acd-4a9e-95f5-c7284b854398"
	productID := 1

	// Inserting an order
	_, err = tx.Exec(`
		INSERT INTO orders (user_id, product_id, order_status)
		VALUES ($1, $2, 'Pending')
	`, userID, productID)
	if err != nil {
		tx.Rollback()
		log.Fatalf("Failed to insert order: %v", err)
	}

	// Decreasing product stock by 1.
	res, err := tx.Exec(`
		UPDATE products SET stock = stock - 1
		WHERE product_id = $1 AND stock > 0
	`, productID)
	if err != nil {
		tx.Rollback()
		log.Fatalf("Failed to update stock: %v", err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		tx.Rollback()
		log.Fatalf("Stock is 0 or product does not exist")
	}

	// Commit
	err = tx.Commit()
	if err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}

	fmt.Println("Order placed successfully and stock updated.")
}
