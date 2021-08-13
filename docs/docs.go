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
            "url": "https://github.com/justfordevandtest/rabbit-finance-test",
            "email": "ekkasith.w@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/:id": {
            "get": {
                "description": "Return a decoded URL of a given shortened URL",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public"
                ],
                "summary": "Access a given shortened URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "410": {
                        "description": "Gone",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/admin": {
            "get": {
                "description": "Return a list of URLs according to the given paginator options",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Protected"
                ],
                "summary": "List a page of URLs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "A page number",
                        "name": "Page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "A total number of items per page",
                        "name": "PerPage",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "An ID filter; will search for a record with a given ID",
                        "name": "ID",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "A URL filter; will search for a record with URL that contains a given keyword",
                        "name": "Keyword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/admin.ListOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/admin/{id}": {
            "delete": {
                "description": "Accessing a deleted URL will get a 410 response",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Protected"
                ],
                "summary": "Delete a URL with a given ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "URL ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/admin.ListOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/shorten": {
            "post": {
                "description": "Return shorten version of a given URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public"
                ],
                "summary": "Shorten a given URL",
                "parameters": [
                    {
                        "description": "URL will not expire if 'expired' is set to null or excluded",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shortener.ShortenInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/shortener.ShortenOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "410": {
                        "description": "Gone",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ErrorItemResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "subCode": {
                    "type": "string"
                }
            }
        },
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ErrorItemResponse"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "SuccessResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "admin.ListOutput": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.URL"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "entity.URL": {
            "type": "object",
            "properties": {
                "expired": {
                    "type": "integer"
                },
                "hitCount": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "shortener.ShortenInput": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "expired": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "shortener.ShortenOutput": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "http://localhost:8080/api/v1/admin/login"
        }
    },
    "x-extension-openapi": {
        "example": "value on a json format"
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
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "URL Shortener API",
	Description: "Create an URL-shortener service to shorten URLs.\\n\\nAPI clients will be able to create short URLs from a full length URL.\\n\\nIt will also support redirecting the short urls to the correct url.",
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
