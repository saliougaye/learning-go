package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func selectAllAlbums(db *pg.DB) []Album {

	var albums []Album

	err := db.Model(&albums).Select()

	if err != nil {
		panic(err)
	}

	return albums
}

func insertAlbum(db *pg.DB, album Album) error {
	_, err := db.Model(&album).Insert()

	if err != nil {
		return err
	}

	return nil

}

func selectSingleAlbum(db *pg.DB, id string) (*Album, error) {
	album := &Album{
		ID: id,
	}

	err := db.Model(album).WherePK().Select()

	if err != nil {
		return nil, err
	}

	return album, nil
}

func deleteAlbum(db *pg.DB, id string) error {

	album := &Album{
		ID: id,
	}

	_, err := db.Model(album).Where("id = ?", id).Delete()

	return err
}

func editAlbum(db *pg.DB, album Album) error {
	_, err := db.Model(&album).WherePK().Update()

	return err

}

func createSchema(db *pg.DB) error {

	models := []interface{}{
		(*Album)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			// Temp: true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
