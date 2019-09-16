package db

const (
	QueryGetByStatus      = `SELECT p.id, status, p.name, c.name AS categoryName, c.id AS categoryId FROM pets.pets p LEFT JOIN pets.categories c ON c.id = p.categoryId WHERE status = ?`
	QueryGetByID          = `SELECT p.id, status, p.name, c.name AS categoryName, c.id AS categoryId FROM pets.pets LEFT JOIN pets.categories c ON c.id = p.categoryId WHERE id = ?`
	QueryDeleteByID       = `DELETE FROM pets.pets WHERE petId = ?`
	QueryUpdatePet        = `UPDATE pets.pets SET name = ?, categoryId = ?, status = ? WHERE id = ?`
	QueryAddTagToPet      = `INSERT INTO pets.pet_tags (petId, tagId) VALUES`
	QueryRemoveTagFromPet = `DELETE FROM pets.pet_tags WHERE petId = ? AND tagId = ?`
	QueryGetTagsByPetID   = `SELECT t.id, t.name FROM pets.tags t JOIN pets.pet_tags pt ON pt.petId = ? ORDER BY t.name ASC`
	QueryGetPhotosByPetID = `SELECT p.* FROM pets.photos WHERE petId = ?`
	QueryAddPet           = `INSERT INTO pets.pets (categoryId, status, name) VALUES (?, ?, ?)`
	QueryAddPhoto         = `INSERT INTO pets.photos (petId, url, filename) VALUES (?, ?, ?)`
)
