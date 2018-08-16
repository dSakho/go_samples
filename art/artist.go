package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

/*
 * Artist Data Struct
 */
type Artist struct {
	ArtistID     uint64 `db:"artist_id"`
	ArtistName   string `db:"artist_name"`
	ArtistType   string `db:"artist_type"`
	CreationDate string `db:"creation_date"`
}

func main() {
	connStr := "postgres://dsakho:Guenny12@localhost/testdb"

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	artists := []Artist{}
	queryErr := db.Select(&artists, "SELECT a.artist_id, a.artist_name, a.artist_type, a.creation_date FROM art.artist a")
	if queryErr != nil {
		log.Fatal(queryErr)
	}

	fmt.Printf("%#v\n", artists[0])

	pablo := Artist{}
	err = db.Get(&pablo, "SELECT a.artist_id, a.artist_name, a.artist_type, a.creation_date FROM art.artist a WHERE a.artist_name=$1", "Pablo Picasso")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%#v\n", pablo)
	}

	artist := Artist{}
	rows, err := db.Queryx("SELECT a.artist_id, a.artist_name, a.artist_type, a.creation_date FROM art.artist a")
	for rows.Next() {
		err := rows.StructScan(&artist)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", artist)
	}

	defer db.Close()
}
