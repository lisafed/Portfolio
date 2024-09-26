package dbFunc

import (
	"database/sql"
	"log"
)

var dbname = "./assets/db/forum.db"

type DBPortfolio struct {
	core *sql.DB
}

//var DB = LoadDb()

func LoadDb() DBPortfolio {
	dbIn, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture de la base de données:", err)
	}
	return DBPortfolio{core: dbIn}
}
