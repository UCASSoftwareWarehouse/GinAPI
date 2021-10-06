package common_model

type Project struct {
	ProjectName string `json:"project_name"`
	Tag         string `json:"tag"`
}

type ProjectFile struct {
	*Project
	RelPath string
	Content string
}
