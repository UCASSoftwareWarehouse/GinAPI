package model

import (
	"GinAPI/models"
	"GinAPI/pb_gen"
)

type PingCodeSimRequest struct {
}

type PingCodeSimResponse struct {
	FullText string `json:"full_text"`
}

type SearchSourceCodeRequest struct {
	Content  string                               `form:"content"`
	CodeType pb_gen.CodeSimSearchRequest_CodeType `form:"code_type"`
	From     int                                  `form:"from"`
	Size     int                                  `form:"size"`
}

type SearchSourceCodeResponse struct {
	ProjectFiles []*models.ProjectFile `json:"project_files"`
}
