package dbFunc

import (
	"database/sql"
)

// GetUtilisateur is a function used to get the data of a user to check if it exists
func GetUtilisateur(db DBPortfolio, nomUtilisateur string) (*Utilisateur, error) {
	var utilisateur Utilisateur
	query := "SELECT id, nom_utilisateur, mot_de_passe FROM utilisateur WHERE nom_utilisateur = ?"
	err := db.core.QueryRow(query, nomUtilisateur).Scan(&utilisateur.ID, &utilisateur.NomUtilisateur, &utilisateur.MotDePasse)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &utilisateur, nil
}

// AddUtilisateur is a function used to add a user to the database
func (db *DBPortfolio) AddUtilisateur(username string, password string) error {
	stmt, err := db.core.Prepare("INSERT INTO Utilisateur(nom_utilisateur, mot_de_passe) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, password)
	return err
}
