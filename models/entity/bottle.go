package entity

type Bottle struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Maturity_date string `json:"maturity_date"`
	Log_list []Log `json:"log_list"`

}