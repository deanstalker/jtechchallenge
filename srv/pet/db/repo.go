package db

import (
	"database/sql"
	pet "github.com/deanstalker/jtechchallenge/srv/pet/proto"
)

type PetRepo interface {
	Add(pet *pet.Pet) error
	Delete(petId int64) error
	Update(pet *pet.Pet) error
	AddPhoto(pet *pet.Pet, photo pet.PhotoURL) error
	AddTag(pet *pet.Pet, tag *pet.Tag) error
	RemoveTag(pet *pet.Pet, tag *pet.Tag) error
	ByID(petID int64) (*pet.Pet, error)
	ByStatus(status string) ([]*pet.Pet, error)
	GetTagsByPetID(petID int64) ([]*pet.Tag, error)
}

type DefaultPetRepo struct {
	db *sql.DB
}

var _ PetRepo = (*DefaultPetRepo)(nil)

func NewPetRepo(db *sql.DB) *DefaultPetRepo {
	return &DefaultPetRepo{
		db: db,
	}
}

func (r *DefaultPetRepo) Add(pet *pet.Pet) error {
	stmt, err := r.db.Prepare(QueryAddPet)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pet.Name, pet.Category.Id, pet.Status)
	if err != nil {
		return err
	}

	return nil
}

func (r *DefaultPetRepo) Delete(petID int64) error {
	stmt, err := r.db.Prepare(QueryDeleteByID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(petID)
	if err != nil {
		return err
	}

	return nil
}

func (r *DefaultPetRepo) Update(pet *pet.Pet) error {
	stmt, err := r.db.Prepare(QueryUpdatePet)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pet.Name, pet.Category.Id, pet.Status, pet.Id)
	if err != nil {
		return err
	}

	for _, tag := range pet.GetTags() {
		if err = r.AddTag(pet, tag); err != nil {
			return err
		}
	}

	return nil
}

func (r *DefaultPetRepo) AddPhoto(pet *pet.Pet, photo pet.PhotoURL) error {
	stmt, err := r.db.Prepare(QueryAddPhoto)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pet.Id, photo.Url, photo.Filename)
	if err != nil {
		return err
	}

	return nil
}

func (r *DefaultPetRepo) AddTag(pet *pet.Pet, tag *pet.Tag) error {
	stmt, err := r.db.Prepare(QueryAddTagToPet)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pet.Id, tag.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *DefaultPetRepo) RemoveTag(pet *pet.Pet, tag *pet.Tag) error {
	stmt, err := r.db.Prepare(QueryRemoveTagFromPet)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pet.Id, tag.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *DefaultPetRepo) ByID(petID int64) (*pet.Pet, error) {
	stmt, err := r.db.Prepare(QueryGetByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(petID)
	p := &pet.Pet{}
	if err := row.Scan(&p); err != nil {
		return nil, err
	}

	p.Tags, err = r.GetTagsByPetID(p.Id)
	if err != nil {
		return nil, err
	}

	p.PhotoUrls, err = r.GetPhotosByPetID(p.Id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *DefaultPetRepo) GetTagsByPetID(petID int64) ([]*pet.Tag, error) {
	stmt, err := r.db.Prepare(QueryGetTagsByPetID)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(petID)
	if err != nil {
		return nil, err
	}

	tags := make([]*pet.Tag, 0)
	for rows.Next() {
		var tag *pet.Tag
		if err := rows.Scan(&tag); err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *DefaultPetRepo) GetPhotosByPetID(petID int64) ([]*pet.PhotoURL, error) {
	stmt, err := r.db.Prepare(QueryGetPhotosByPetID)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(petID)
	if err != nil {
		return nil, err
	}

	photos := make([]*pet.PhotoURL, 0)
	for rows.Next() {
		var photo *pet.PhotoURL
		if err := rows.Scan(&photo); err != nil {
			return nil, err
		}

		photos = append(photos, photo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return photos, nil
}

func (r *DefaultPetRepo) ByStatus(status string) ([]*pet.Pet, error) {
	stmt, err := r.db.Prepare(QueryGetByStatus)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, err
	}

	pets := make([]*pet.Pet, 0)
	for rows.Next() {
		var row *pet.Pet
		if err := rows.Scan(&row); err != nil {
			return nil, err
		}

		row.PhotoUrls, err = r.GetPhotosByPetID(row.Id)
		if err != nil {
			return nil, err
		}

		row.Tags, err = r.GetTagsByPetID(row.Id)
		if err != nil {
			return nil, err
		}

		pets = append(pets, row)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pets, nil
}
