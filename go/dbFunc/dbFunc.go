package dbFunc

import (
	"database/sql"
	"log"
)

var dbname = "./Portfolio./Portfolio.db"

//var DB = LoadDb()

type DBPortfolio struct {
	core *sql.DB
}

func LoadDb() DBPortfolio {
	dbIn, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture de la base de données:", err)
	}
	return DBPortfolio{core: dbIn}
}
