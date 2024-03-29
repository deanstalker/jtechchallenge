package db

var (
	QueryAdd           = `INSERT INTO user.user(username, firstName, lastName, email, password, phone, userStatus) VALUES (?, ?, ?, ?, ?, ?, ?)`
	QueryBatchAdd      = `INSERT INTO user.user(username, firstName, lastName, email, password, phone, userStatus) VALUES`
	QueryGetByUsername = `SELECT id, username, firstName, lastName, email, password, phone, userStatus, token FROM user.user WHERE username = ?`
	QueryUpdate        = `UPDATE user.user SET username = ?, firstName = ?, lastName = ?, email = ?, phone = ?, userStatus = ?, token = ? WHERE id = ?`
	QueryDelete        = `DELETE FROM user.user WHERE id = ?`
)
