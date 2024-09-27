package dbFunc

import (
	"database/sql"
)

// Fonction pour récupérer un utilisateur avec son nom d'utilisateur
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

func (db *DBPortfolio) AddUtilisateur(username string, password string) error {
	stmt, err := db.core.Prepare("INSERT INTO Utilisateur(nom_utilisateur, mot_de_passe) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, password)
	return err
}
