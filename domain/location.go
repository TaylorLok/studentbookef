package domain

type Location struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Title       string `json:"title"`
}
