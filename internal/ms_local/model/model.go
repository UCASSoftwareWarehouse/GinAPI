package model

import (
	"GinAPI/pb_gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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

type ProjectDetail struct {
	Id                  uint64                 `form:"id" json:"id,omitempty"`
	ProjectName         string                 `form:"projectName" json:"projectName,omitempty"`
	UserId              uint64                 `form:"userId" json:"userId,omitempty"`
	Tags                string                 `form:"tags" json:"tags,omitempty"`
	License             string                 `form:"license" json:"license,omitempty"`
	Updatetime          *timestamppb.Timestamp `form:"updatetime" json:"updatetime,omitempty"`
	ProjectDescription  string                 `form:"projectDescription" json:"projectDescription,omitempty"`
	CodeAddr            string                 `form:"codeAddr" json:"codeAddr,omitempty"`
	BinaryAddr          string                 `form:"binaryAddr" json:"binaryAddr,omitempty"`
	OperatingSystem     string                 `form:"operatingSystem" json:"operatingSystem,omitempty"`
	ProgrammingLanguage string                 `form:"programmingLanguage" json:"programmingLanguage,omitempty"`
	NaturalLanguage     string                 `form:"naturalLanguage" json:"naturalLanguage,omitempty"`
	Topic               string                 `form:"topic" json:"topic,omitempty"`
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
const (
	OSmap uint32 = 0x000000ff //operating system
	PLmap uint32 = 0x0000ff00 //programing language
	NLmap uint32 = 0x00ff0000 // natural language
	Tomap uint32 = 0xff000000 //topic
)


func GetOSName(classifer uint32) string {
	return pb_gen.OperatingSystem_name[int32(classifer & OSmap)]
}

func GetPLName(classifer uint32) string {
	return pb_gen.ProgrammingLanguage_name[int32((classifer & PLmap) >> 8)]
}

func GetNLName(classifer uint32) string {
	return pb_gen.NaturalLanguage_name[int32((classifer & NLmap) >> 16)]
}

//topic value
func GetToName(classifer uint32) string {
	return pb_gen.Topic_name[int32(((classifer & Tomap) >> 24))]
}