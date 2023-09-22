package response

type Bug struct {
	Id              uint    `json:"id"`
	CreateTime      int64   `json:"createTime"`
	UpdateTime      int64   `json:"updateTime"`
	Reporter        User    `json:"reporter"`
	Handler         User    `json:"handler"`
	Status          int     `json:"status"`
	Priority        int     `json:"priority"`
	Severity        int     `json:"severity"`
	Reproducibility int     `json:"reproducibility"`
	Project         Project `json:"project"`
	Version         Version `json:"version"`
	Group           Group   `json:"group"`
	Summary         string  `json:"summary"`
	Description     string  `json:"description"`
}

type BugBrief struct {
	Id              uint    `json:"id"`
	CreateTime      int64   `json:"createTime"`
	UpdateTime      int64   `json:"updateTime"`
	Reporter        User    `json:"reporter"`
	Handler         User    `json:"handler"`
	Status          int     `json:"status"`
	Priority        int     `json:"priority"`
	Severity        int     `json:"severity"`
	Reproducibility int     `json:"reproducibility"`
	Project         Project `json:"project"`
	Version         Version `json:"version"`
	Group           Group   `json:"group"`
}
