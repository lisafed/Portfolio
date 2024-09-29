package dbFunc

import "database/sql"

// DeleteExperience is a func used to delete an experience based on an id entered as a parameter
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

// AddExperience is a function used to add an experience based on an Experience struct entered as a parameter
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

// EditExperience is a function used to edit an experience based on the new value used as parameters
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

// ReadExperience is a function used to return an experience based on the id used as a parameter
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

// GetExperience is a function used to return all experiences from the database
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
