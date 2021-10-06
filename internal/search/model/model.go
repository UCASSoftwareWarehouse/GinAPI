package model

import (
	"GinAPI/internal/common_model"
	"GinAPI/pb_gen"
)

type PingCodeSimRequest struct {
}

type PingCodeSimResponse struct {
	FullText string `json:"full_text"`
}

type QueryCodeRequest struct {
	Content  string                               `json:"content"`
	CodeType pb_gen.CodeSimSearchRequest_CodeType `json:"code_type"`
	From     int                                  `json:"from"`
	Size     int                                  `json:"size"`
}

type QueryCodeResponse struct {
	ProjectFiles []*common_model.ProjectFile `json:"project_files"`
}
