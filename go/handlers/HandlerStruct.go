package handlers

import "portfolio/go/dbFunc"

type Data struct {
	IsUserLogged   bool
	DataFormation  []dbFunc.Formation
	DataExperience []dbFunc.Experience
	DataProject    []dbFunc.Projet
}
