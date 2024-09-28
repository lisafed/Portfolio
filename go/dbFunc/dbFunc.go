package dbFunc

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var dbname = "Portfolio.db"

type DBPortfolio struct {
	core *sql.DB
}

func LoadDb() DBPortfolio {
	dbIn, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture de la base de donnÃ©es:", err)
	}
	return DBPortfolio{core: dbIn}
}
