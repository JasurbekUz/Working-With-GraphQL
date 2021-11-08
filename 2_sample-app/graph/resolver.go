package graph

import (

	"app/graph/model"
	"database/sql"
	"log"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{

	DB DB
}

type DB struct {

	Conn *sql.DB
}


func (r *DB) GetClassifieds () []*model.Classified {


	rows, err := r.Conn.Query(
		"select classified_id, title, price, created_at from classifieds",
	)

	defer rows.Close()

	if err != nil {
		log.Fatalf("db query error:\n %v", err)
	}

	var classifieds []*model.Classified

	for rows.Next() {

		var classified model.Classified

		rows.Scan(
			&classified.ID,
			&classified.Title,
			&classified.Price,
			&classified.CreatedAt,
		)

		classifieds = append(classifieds, &classified)
	}

	return classifieds

}

func (r *DB) GetClassified (id int) *model.Classified {

	var classified model.Classified

	err := r.Conn.QueryRow(
		"select classified_id, title, price, created_at from classifieds where classified_id = $1",
		id,
	).Scan(
			&classified.ID,
			&classified.Title,
			&classified.Price,
			&classified.CreatedAt,
		)

	if err != nil {
		return &classified
		//log.Fatalf("db query error:\n %v", err)
	}

	return &classified

}

func (r *DB) GetCategories () []*model.Category {

	var categories []*model.Category

	rows, err := r.Conn.Query(
		"select category_id, name from categories",
	)

	defer rows.Close()

	if err != nil {
		return categories
		//log.Fatalf("db query error:\n %v", err)
	}

	for rows.Next() {

		var category model.Category

		rows.Scan(
			&category.ID,
			&category.Name,
			
		)

		categories = append(categories, &category)
	}

	return categories

}

func (r *DB) GetCategory (id int) *model.Category {

	var category model.Category

	err := r.Conn.QueryRow(
		"select category_id, name from categories where category_id = $1",
		id,
	).Scan(
			&category.ID,
			&category.Name,	
		)

	if err != nil {
		return &category
		//log.Fatalf("db query error:\n %v", err)
	}

	return &category

}

func (r *DB) GetClassifiedsByCategotyId(catId int) []*model.Classified {


	rows, err := r.Conn.Query(
		"select classified_id, title, price, created_at from classifieds where category_id = $1",
		catId,
	)

	defer rows.Close()

	if err != nil {
		log.Fatalf("db query error:\n %v", err)
	}

	var classifieds []*model.Classified

	for rows.Next() {

		var classified model.Classified

		rows.Scan(
			&classified.ID,
			&classified.Title,
			&classified.Price,
			&classified.CreatedAt,
		)

		classifieds = append(classifieds, &classified)
	}

	return classifieds

}

func (r *DB) GetCurrentClassifiedSCategoryId (clId int) (catId int) {

	err := r.Conn.QueryRow(
		"select category_id from classifieds where classified_id = $1",
		clId,
	).Scan(
		&catId,
	)

	if err != nil {
		return 
		//log.Fatalf("db query error:\n %v", err)
	}

	return catId
}

func (r *DB) GetClassifiedSCategory (id int)  *model.Category {

	var category model.Category

	err := r.Conn.QueryRow(
		"select category_id, name from categories where category_id = $1",
		id,
	).Scan(
			&category.ID,
			&category.Name,	
		)

	if err != nil {
		return &category
		//log.Fatalf("db query error:\n %v", err)
	}

	return &category

}