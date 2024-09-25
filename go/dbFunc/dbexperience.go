package dbFunc

import "database/sql"

func (db *DBPortfolio) DeleteExperience(id int) error {
	stmt, err := db.core.Prepare("DELETE FROM Experiences WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return err
}

func (db *DBPortfolio) AddExperience(exp Experience) error {
	stmt, err := db.core.Prepare("INSERT INTO Experiences (nom, entreprise, poste, date_debut, date_fin) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(exp.Nom, exp.Entreprise, exp.Poste, exp.DateDebut, exp.DateFin)
	if err != nil {

		return err
	}

	return err
}

func (db *DBPortfolio) EditExperience(exp Experience) error {
	stmt, err := db.core.Prepare("UPDATE Experiences SET nom=?, entreprise=?, poste=?, date_debut=?, date_fin=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(exp.Nom, exp.Entreprise, exp.Poste, exp.DateDebut, exp.DateFin, exp.ID)
	if err != nil {
		return err
	}

	return err
}

func (db *DBPortfolio) ReadExperience(id int) (Experience, error) {
	var exp Experience

	stmt, err := db.core.Prepare("SELECT id, nom, entreprise, poste, date_debut, date_fin FROM Experiences WHERE id=?")
	if err != nil {
		return exp, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&exp.ID, &exp.Nom, &exp.Entreprise, &exp.Poste, &exp.DateDebut, &exp.DateFin)
	if err != nil {
		if err == sql.ErrNoRows {
			return exp, nil
		}
		return exp, err
	}

	return exp, nil
}

func (db DBPortfolio) GetExperience() ([]Experience, error) {
	rows, err := db.core.Query("SELECT id, nom, entreprise, poste, date_debut, date_fin FROM Experiences")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var experiences []Experience
	for rows.Next() {
		var exp Experience
		err := rows.Scan(&exp.ID, &exp.Nom, &exp.Entreprise, &exp.Poste, &exp.DateDebut, &exp.DateFin)
		if err != nil {
			return nil, err
		}
		experiences = append(experiences, exp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return experiences, nil
}

// DeleteExperience : supprime une expérience en fonction de l'id
/*func (db *DBPortfolio) DeleteExperience(id int) error {
	stmt, err := db.core.Prepare("DELETE FROM Experiences WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err // retourne directement l'erreur s'il y en a une
}

// AddExperience : ajoute une nouvelle expérience dans la base de données
func (db *DBPortfolio) AddExperience(exp Experience) error {
	stmt, err := db.core.Prepare("INSERT INTO Experiences (nom, entreprise, poste, date_debut, date_fin) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(exp.Nom, exp.Entreprise, exp.Poste, exp.DateDebut, exp.DateFin)
	return err
}

// EditExperience : met à jour une expérience existante
func (db *DBPortfolio) EditExperience(exp Experience) error {
	stmt, err := db.core.Prepare("UPDATE Experiences SET nom=?, entreprise=?, poste=?, date_debut=?, date_fin=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(exp.Nom, exp.Entreprise, exp.Poste, exp.DateDebut, exp.DateFin, exp.ID)
	return err
}

// ReadExperience : récupère une expérience par ID
func (db *DBPortfolio) ReadExperience(id int) (Experience, error) {
	var exp Experience

	stmt, err := db.core.Prepare("SELECT id, nom, entreprise, poste, date_debut, date_fin FROM Experiences WHERE id=?")
	if err != nil {
		return exp, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&exp.ID, &exp.Nom, &exp.Entreprise, &exp.Poste, &exp.DateDebut, &exp.DateFin)
	if err != nil {
		if err == sql.ErrNoRows {
			return exp, nil // Retourne nil pour une absence de résultat, car ce n'est pas une vraie "erreur"
		}
		return exp, err
	}

	return exp, nil
}

// GetExperience : récupère toutes les expériences
func (db DBPortfolio) GetExperience() ([]Experience, error) {
	rows, err := db.core.Query("SELECT id, nom, entreprise, poste, date_debut, date_fin FROM Experiences")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var experiences []Experience
	for rows.Next() {
		var exp Experience
		err := rows.Scan(&exp.ID, &exp.Nom, &exp.Entreprise, &exp.Poste, &exp.DateDebut, &exp.DateFin)
		if err != nil {
			return nil, err
		}
		experiences = append(experiences, exp)
	}

	// Vérifie les erreurs lors du parcours des lignes
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return experiences, nil
}*/
