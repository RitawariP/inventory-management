package storage

const (
	insertQuery       = "INSERT INTO products (name, description, price, created) VALUES (?, ?, ?, ?)"
	selectAllQuery    = "SELECT * FROM products"
	selectByNameQuery = "SELECT * FROM products WHERE name = ?"
	updateQuery       = "UPDATE products SET description = ?, price = ? WHERE name = ?"
)
