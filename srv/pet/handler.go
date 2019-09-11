package main

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

type DefaultPet struct {
	db *sql.DB
	log logrus.FieldLogger
}

var _ pet.PetService = (*DefaultPet)(nil)

func NewPetService(db *sql.DB, log logrus.FieldLogger) *DefaultPet {
	return &DefaultPet{
		db: db,
		log: log,
	}
}

func (p *DefaultPet) Add() {

}