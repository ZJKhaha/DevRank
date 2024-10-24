package models

type Developer struct {
	Username         string  `json:"username"`
	Repos            []Repo  `json:"repos"`
	Nation           string  `json:"nation"`
	OriginalLocation string  `json:"originalLocation"` // 新增字段
	TalentRank       float64 `json:"talentRank"`
}

type Repo struct {
	Name       string `json:"name"`
	Stars      int    `json:"stars"`
	Forks      int    `json:"forks"`
	OpenIssues int    `json:"openIssues"`
}
