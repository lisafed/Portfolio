package dbFunc

type Experience struct {
	ID         int
	Nom        string
	Entreprise string
	Poste      string
	DateDebut  string
	DateFin    string
}

type Formation struct {
	ID            int
	Nom           string
	Formation     string
	Etablissement string
	DateDebut     string
	DateFin       string
}

type Projet struct {
	ID        int
	NomProjet string
	LienRepo  string
	Langage   string
}
type Utilisateur struct {
	ID             int
	NomUtilisateur string
	MotDePasse     string
}
