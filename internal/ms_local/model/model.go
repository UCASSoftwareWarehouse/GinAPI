package model

type PingMSLocalRequest struct {
}

type PingMSLocalResponse struct {
	FullText string `json:"full_text"`
}

type MSLocalUploadRequest struct {
}
type MSLocalUploadBinaryResponse struct {
}

type MSLocalUploadCodesResponse struct {
}

type CreateProjectRequest struct {
	ProjectName         string `form:"projectName" json:"projectName" binding:"required"`
	UserId              uint64 `form:"userId" json:"userId" binding:"required"`
	Tags                string `form:"tags" json:"tags" binding:"required"`
	License             string `form:"license" json:"license" binding:"required"`
	ProjectDescription  string `from:"projectDescription" json:"projectDescription" binding:"required"`
	OperatingSystem     string `form:"operatingSystem" json:"operatingSystem" binding:"required"`
	ProgrammingLanguage string `form:"programmingLanguage" json:"programmingLanguage" inding:"required"`
	NaturalLanguage     string `form:"naturalLanguage" json:"naturalLanguage" binding:"required"`
	Topic               string `form:"topic" json:"topic" binding:"required"`
}

type CreateProjectResponse struct {
}

func GetClassifier(osValue uint8, plValue uint8, nlValue uint8, toValue uint8) uint32 {
	var res uint32
	res += uint32(osValue)
	res += (uint32(plValue) << 8)
	res += (uint32(nlValue) << 16)
	res += (uint32(toValue) << 24)
	return res
}
