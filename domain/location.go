package domain

type Location struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	Description string `json:"description"`
}
