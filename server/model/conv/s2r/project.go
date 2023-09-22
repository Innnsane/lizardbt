package s2r

import (
	"lizardbt/server/model/response"
	"lizardbt/server/model/schema"
)

func Project(in *schema.Project) (out response.Project) {
	return response.Project{
		Id:   in.ID,
		Name: in.Name,
	}
}
