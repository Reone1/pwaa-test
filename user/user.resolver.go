package User

func GetUserById(s string) User {
	return User{
		ID: "user id string",
		Name: "user name string",
		Type: "user type string",
	}
}

func GetUsers(ids ...string) []User {
	return []User{{
		ID: "user id string1",
		Name: "user name string1",
		Type: "user type string1",
	},
	{
		ID: "user id string2",
		Name: "user name string2",
		Type: "user type string2",
	}}
}