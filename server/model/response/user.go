package response

type User struct {
	Id       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Role     int    `json:"role"`
}
