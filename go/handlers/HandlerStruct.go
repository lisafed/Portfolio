package handlers

import "portfolio/go/dbFunc"

type Data struct {
	IsUserLogged       bool
	DataFormationList  []dbFunc.Formation
	DataExperienceList []dbFunc.Experience
	DataProjectList    []dbFunc.Projet

	UnitDataProject    dbFunc.Projet
	UnitDataExperience dbFunc.Experience
	DataFormation      dbFunc.Formation
}
