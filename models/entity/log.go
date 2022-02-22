package entity

type Log struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Worth int `json:"worth"`
	UserId string `json:"userId"`
	Bottle_Id string `json:"bottle_id"`
}