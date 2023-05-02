package types

type Country struct {
	Id          int
	CountryName string
}

type University struct {
	Id             int
	CountryName    string
	UniversityName string
}

type RankingCriteria struct {
	Id           int
	SystemID     int
	CriteriaName string
}

type ChangeStudentStaffRatio struct {
	UniversityName string
	Year           int
	NewStaffRatio  int
}
