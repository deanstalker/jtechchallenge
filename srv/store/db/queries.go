package db

var (
	QueryGetOrderByID       = `SELECT id, petId, quantity, shipDate, status, complete FROM store.order WHERE id = ?`
	QueryGetOrdersByPetID   = `SELECT id, petId, quantity, shipDate, status, complete FROM store.order WHERE petId = ?`
	QueryGetOrdersByStatus  = `SELECT id, petId, quantity, shipDate, status, complete FROM store.order WHERE status = ?`
	QueryGetCompletedOrders = `SELECT id, petId, quantity, shipDate, status, complete FROM store.order WHERE complete = 1`
	QueryGetInventory       = `SELECT status, COUNT(status) AS total FROM store.order GROUP BY status ORDER BY status ASC `
	QueryAddOrder           = `INSERT INTO store.order (petId, quantity, status) VALUES (?, ?, ?)`
	QueryUpdateOrder        = `UPDATE store.order SET quantity = ?, shipDate = ?, status = ?, complete = ? WHERE id = ?`
)
