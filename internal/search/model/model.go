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

// SearchSourceCodeRequest
// Content 待搜索文本
// CodeTypes 将搜索文本作为何种代码进行搜索，可以作为多种代码，0=>纯文本，1=>python, 2=>golang, 3=>java
type SearchSourceCodeRequest struct {
	Content    string                               `form:"content"`
	CodeTypes   []pb_gen.CodeSimSearchRequest_CodeType `form:"codeType"`
	From       int                                  `form:"from"`
	Size       int                                  `form:"size"`
	WithSource bool                                 `form:"withSource"`
}

type SearchSourceCodeResponse struct {
	ProjectFiles []*models.ProjectFile `json:"project_files"`
}

// DeleteSourceCodeRequest 删除源代码请求
type DeleteSourceCodeRequest struct {
	ProjectName string `json:"project_name"`
	Tag string `json:"tag"`
}

type DeleteSourceCodeResponse struct {

}