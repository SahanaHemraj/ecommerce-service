// internal/database/migrate.go
package database

import (
	"log"
	"github.com/lib/pq"
)

func Migrate() {
	// Product Table
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			description TEXT,
			price NUMERIC(10,2) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal("Error creating products table: ", err)
	}

	// User Table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(100) UNIQUE NOT NULL,
			password VARCHAR(100) NOT NULL,
			name VARCHAR(100),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal("Error creating users table: ", err)
	}

	// Order Table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
			total NUMERIC(10,2) NOT NULL,
			status VARCHAR(50) DEFAULT 'Pending',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal("Error creating orders table: ", err)
	}

	// Cart Table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS carts (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
			product_id INTEGER REFERENCES products(id),
			quantity INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal("Error creating carts table: ", err)
	}
}
