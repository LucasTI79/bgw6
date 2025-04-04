package products

var (
	GetQuery = `
		SELECT id, name, price, quantity, type FROM products
	`
	GetOneQuery = `
		SELECT id, name, price, quantity, type FROM products
		WHERE id = ?
	`
	StoreQuery = `
		INSERT INTO products (name, price, quantity, type) VALUES (?, ?, ?, ?)
	`
	UpdateQuery = `
		UPDATE products SET name=?, price=?, quantity=?, type=? WHERE id=?
	`
	DeleteQuery = `
		DELETE FROM products WHERE id=?
	`
)
