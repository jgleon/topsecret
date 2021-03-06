// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "soberkoder@swagger.io"
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
        "/topsecret": {
            "post": {
                "description": "Determine the location of the spacecraft using the data sent",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "topsecret"
                ],
                "summary": "Determine the location of the spaceship",
                "parameters": [
                    {
                        "description": "information messages from satellites",
                        "name": "satellites",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Satellites"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Location"
                        }
                    },
                    "404": {}
                }
            }
        },
        "/topsecret_split": {
            "get": {
                "description": "Determine the location of the spacecraft using the data sent",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "topsecret_split"
                ],
                "summary": "Determine the location of the spaceship",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Location"
                        }
                    },
                    "404": {}
                }
            }
        },
        "/topsecret_split/{satellite_name}": {
            "post": {
                "description": "Records a satellite message in memory to calculate the location of the spacecraft later",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "topsecret_split"
                ],
                "summary": "Records a satellite message in memory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "satellite name",
                        "name": "satellite_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "infomation message satellite",
                        "name": "infoSatellite",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InfoSatellite"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {}
                }
            }
        }
    },
    "definitions": {
        "models.InfoSatellite": {
            "type": "object",
            "properties": {
                "Distance": {
                    "type": "number",
                    "example": 100
                },
                "Message": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.Location": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "position": {
                    "type": "object",
                    "properties": {
                        "x": {
                            "type": "number"
                        },
                        "y": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "models.Satellite": {
            "type": "object",
            "properties": {
                "Distance": {
                    "type": "number",
                    "example": 100
                },
                "Message": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "Name": {
                    "type": "string",
                    "example": "kenobi"
                }
            }
        },
        "models.Satellites": {
            "type": "object",
            "properties": {
                "Satellites": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Satellite"
                    }
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
	Host:        "localhost:3000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Localization API",
	Description: "Manage satellite messages and calculate the location of the spacecraft",
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
