package storage

const (
	insertQuery       = "INSERT INTO products (name, description, price, created) VALUES (?, ?, ?, ?)"
	selectAllQuery    = "SELECT * FROM products"
	selectByNameQuery = "SELECT * FROM products WHERE name = ?"
	updateQuery       = "UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?"
	deleteQuery       = "DELETE FROM products WHERE id = ?"
)
