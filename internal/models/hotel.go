package models

type Comment struct {
	User    string
	Message string
	Rating  int
}

type Hotel struct {
	ID          int
	City        string
	Name        string
	Description string
	ImageURL    string
	MapURL      string
	Address     string
	Facilities  []string
	Comments    []Comment
}
