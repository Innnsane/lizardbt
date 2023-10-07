package s2r

import (
	"lizardbt/model/response"
	"lizardbt/model/schema"
)

func Version(in *schema.Version) (out response.Version) {
	return response.Version{
		Id:   in.ID,
		Name: in.Name,
	}
}
