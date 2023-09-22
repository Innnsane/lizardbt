package s2r

import (
	"lizardbt/server/model/response"
	"lizardbt/server/model/schema"
)

func Bug(in *schema.Bug) (out response.Bug) {
	return response.Bug{
		Id:              in.ID,
		CreateTime:      in.CreatedAt.UnixMilli(),
		UpdateTime:      in.UpdatedAt.UnixMilli(),
		Reporter:        User(in.Reporter),
		Handler:         User(in.Handler),
		Status:          in.Status,
		Priority:        in.Priority,
		Severity:        in.Severity,
		Reproducibility: in.Reproducibility,
		Project:         Project(in.Project),
		Version:         Version(in.Version),
		Group:           Group(in.Group),
		Summary:         in.Summary,
		Description:     in.Description,
	}
}

func BugBrief(in *schema.Bug) (out response.BugBrief) {
	return response.BugBrief{
		Id:              in.ID,
		CreateTime:      in.CreatedAt.UnixMilli(),
		UpdateTime:      in.UpdatedAt.UnixMilli(),
		Reporter:        User(in.Reporter),
		Handler:         User(in.Handler),
		Status:          in.Status,
		Priority:        in.Priority,
		Severity:        in.Severity,
		Reproducibility: in.Reproducibility,
		Project:         Project(in.Project),
		Version:         Version(in.Version),
		Group:           Group(in.Group),
	}
}
