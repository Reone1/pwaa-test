package User

type User struct {
	ID string `json:"_id"`
	NickName string `json:"nickName"`
	Mail string `json:"mail"`
	Type string `json:"type"`
}

type Token struct {
	Key string `json:"key"`
	Type string `json:"type"`
}