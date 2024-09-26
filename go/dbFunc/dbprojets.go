package dbFunc

import "database/sql"

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

func (db *DBPortfolio) DeleteProject(id int) error {
	stmt, err := db.core.Prepare("DELETE FROM Projets WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func (db *DBPortfolio) AddProjet(proj Projet) error {
	stmt, err := db.core.Prepare("INSERT INTO Projets (nom_projet, lien_repo, langage) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(proj.NomProjet, proj.LienRepo, proj.Langage)
	return err
}

func (db *DBPortfolio) EditProjet(p Projet) error {
	stmt, err := db.core.Prepare("UPDATE Projets SET nom_projet=?, lien_repo=?, langage=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.NomProjet, p.LienRepo, p.Langage, p.ID)
	return err
}

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
