// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/binary/local": {
            "patch": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "binary"
                ],
                "summary": "本地上传二进制包",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "projectName",
                        "name": "projectName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tag",
                        "name": "tag",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.LocalPatchBinaryResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/source_code/local": {
            "patch": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "source_code"
                ],
                "summary": "local上传源代码",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "projectName",
                        "name": "projectName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tag",
                        "name": "tag",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "fileType",
                        "name": "fileType",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.LocalPatchSourceCodeResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/source_code/search": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "source_code"
                ],
                "summary": "模糊查询代码",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "name": "codeTypes",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "content",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "name": "withSource",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SearchSourceCodeResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/test/error_handler": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "test"
                ],
                "summary": "test error handler",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/test/ping_code_sim": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "test"
                ],
                "summary": "ping code sim微服务测试",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PingCodeSimResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.LocalPatchBinaryResponse": {
            "type": "object"
        },
        "model.LocalPatchSourceCodeResponse": {
            "type": "object"
        },
        "model.PingCodeSimResponse": {
            "type": "object",
            "properties": {
                "full_text": {
                    "type": "string"
                }
            }
        },
        "model.SearchSourceCodeResponse": {
            "type": "object",
            "properties": {
                "project_files": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ProjectFile"
                    }
                }
            }
        },
        "models.ProjectFile": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "project_name": {
                    "type": "string"
                },
                "relPath": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "SoftwareWarehouse Web API",
	Description: "This is a SoftwareWarehouse API server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
