package entities

type Question struct {
	Id           int
	Text         string
	Alternatives []Alternative
}
