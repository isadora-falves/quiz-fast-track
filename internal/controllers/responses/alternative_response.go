package responses

type AlternativeResponse struct {
	Id     int    `json:"id" example:"1"`
	Option string `json:"option" example:"A"`
	Text   string `json:"text" example:"Malta"`
}
