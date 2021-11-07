package graph

import (
	"app/graph/model"
	"math/rand"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var users []*model.User
var professions []*model.Profession

type DB struct {

	usersDb []*model.User
}

func (db *DB) GetAllServants () []*model.User {
	return users
}

func (db *DB) GetAllProfessions () []*model.Profession {
	return professions
}

func (db *DB) CreateUser (newUserName *model.NewUser) *model.User {

	rand.Seed(time.Now().UnixNano())

	newUser := model.User {
		ID: rand.Intn(100000),
		Username: newUserName.Username,
	}

	users = append(users, &newUser)

	return &newUser
} 

func (db *DB) CreateProfession (name string) *model.Profession {

	newProfession := model.Profession {
		ID: len(professions) + 1,
		Name: name,
	}

	professions = append(professions, &newProfession)

	return &newProfession
} 


type Resolver struct{

	*DB
}
