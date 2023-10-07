package s2r

import (
	"lizardbt/model/response"
	"lizardbt/model/schema"
)

func Group(in *schema.Group) (out response.Group) {
	return response.Group{
		Id:   in.ID,
		Name: in.Name,
	}
}
