package s2r

import (
	"lizardbt/model/response"
	"lizardbt/model/schema"
)

func User(in *schema.User) (out response.User) {
	return response.User{
		Id:       in.ID,
		Name:     in.Name,
		Nickname: in.Nickname,
		Role:     in.Role,
	}
}
