package s2r

import (
	"lizardbt/server/model/response"
	"lizardbt/server/model/schema"
)

func Group(in *schema.Group) (out response.Group) {
	return response.Group{
		Id:   in.ID,
		Name: in.Name,
	}
}
