package dbFunc

import "database/sql"

// GetProjets is a function used to return all experiences from the database
func (db DBPortfolio) GetProjets() ([]Projet, error) {
	rows, err := db.core.Query("SELECT id, nom_projet, lien_repo, langage FROM Projets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projets []Projet
	for rows.Next() {
		var proj Projet
		err := rows.Scan(&proj.ID, &proj.NomProjet, &proj.LienRepo, &proj.Langage)
		if err != nil {
			return nil, err
		}
		projets = append(projets, proj)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projets, nil
}

// DeleteProject is a function used to delete a project from the database based on the id entered as parameter
func (db *DBPortfolio) DeleteProject(id int) error {
	stmt, err := db.core.Prepare("DELETE FROM Projets WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

// AddProjet is a function used to add a project to the database based on the infos entered as parameter
func (db *DBPortfolio) AddProjet(proj Projet) error {
	stmt, err := db.core.Prepare("INSERT INTO Projets (nom_projet, lien_repo, langage) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(proj.NomProjet, proj.LienRepo, proj.Langage)
	return err
}

// EditProjet is a function used to edit project based on the new data entered as parameters
func (db *DBPortfolio) EditProjet(p Projet) error {
	stmt, err := db.core.Prepare("UPDATE Projets SET nom_projet=?, lien_repo=?, langage=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.NomProjet, p.LienRepo, p.Langage, p.ID)
	return err
}

// ReadProjet is a function used to return data of a project based on the id entered as parameter
func (db *DBPortfolio) ReadProjet(id int) (Projet, error) {
	var p Projet

	stmt, err := db.core.Prepare("SELECT id, nom_projet, lien_repo, langage FROM Projets WHERE id=?")
	if err != nil {
		return p, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&p.ID, &p.NomProjet, &p.LienRepo, &p.Langage)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, nil
		}
		return p, err
	}
	return p, nil
}
