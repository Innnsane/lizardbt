package s2r

import (
	"lizardbt/model/response"
	"lizardbt/model/schema"
)

func Project(in *schema.Project) (out response.Project) {
	return response.Project{
		Id:   in.ID,
		Name: in.Name,
	}
}
