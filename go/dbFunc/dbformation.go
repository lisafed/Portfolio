package dbFunc

import "database/sql"

// DeleteFormation is a function used to delete formation based on the id inserted as a parameter
func (db *DBPortfolio) DeleteFormation(id int) error {
	stmt, err := db.core.Prepare("DELETE FROM Formations WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

// AddFormation is a function used to add an experience based on a Formation struct entered as a parameter
func (db *DBPortfolio) AddFormation(f Formation) error {
	stmt, err := db.core.Prepare("INSERT INTO Formations (nom, formation, etablissement, date_debut, date_fin) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(f.Nom, f.Formation, f.Etablissement, f.DateDebut, f.DateFin)
	return err
}

// EditFormation is a function used to edit a formation based on the new value entered as parameters
func (db *DBPortfolio) EditFormation(f Formation) error {
	stmt, err := db.core.Prepare("UPDATE Formations SET nom=?, formation=?, etablissement=?, date_debut=?, date_fin=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(f.Nom, f.Formation, f.Etablissement, f.DateDebut, f.DateFin, f.ID)
	return err
}

// ReadFormation is a function used to return a formation based on the id used as a parameter
func (db *DBPortfolio) ReadFormation(id int) (Formation, error) {
	var f Formation

	stmt, err := db.core.Prepare("SELECT id, nom, formation, etablissement, date_debut, date_fin FROM Formations WHERE id=?")
	if err != nil {
		return f, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&f.ID, &f.Nom, &f.Formation, &f.Etablissement, &f.DateDebut, &f.DateFin)
	if err != nil {
		if err == sql.ErrNoRows {
			return f, nil
		}
		return f, err
	}

	return f, nil
}

// GetFormation is a function used to return all formations from the database
func (db DBPortfolio) GetFormation() ([]Formation, error) {
	rows, err := db.core.Query("SELECT id, nom,	formation, etablissement, date_debut, date_fin FROM Formations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var formations []Formation
	for rows.Next() {
		var exp Formation
		err := rows.Scan(&exp.ID, &exp.Nom, &exp.Formation, &exp.Etablissement, &exp.DateDebut, &exp.DateFin)
		if err != nil {
			return nil, err
		}
		formations = append(formations, exp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return formations, nil
}
