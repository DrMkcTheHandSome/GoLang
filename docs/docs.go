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
        "contact": {
            "name": "Marc Kenneth Lomio",
            "email": "marckenneth.lomio@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "show html that navigates to google auth login",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/googlecallback": {
            "get": {
                "description": "Get details of the user",
                "produces": [
                    "application/json"
                ],
                "summary": "Get details of the user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Google"
                        }
                    }
                }
            }
        },
        "/migration": {},
        "/product": {},
        "/product/{id}": {
            "get": {
                "description": "retrieve the product corresponding to the input productId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "retrieve product identified by the given productId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the product to be retrieve",
                        "name": "productId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "Delete the product corresponding to the input productId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Delete product identified by the given productId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the product to be deleted",
                        "name": "productId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Get details of all products",
                "produces": [
                    "application/json"
                ],
                "summary": "Get details of all products",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user": {},
        "/user/login": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Login user with password",
                "parameters": [
                    {
                        "description": "Create user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/loginViaGoogle": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Login user using google account",
                "parameters": [
                    {
                        "description": "Create user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get details of all users",
                "produces": [
                    "application/json"
                ],
                "summary": "Get details of all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.User"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "verified_email": {
                    "type": "boolean"
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
	Host:        "localhost:9000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Users Product Go Rest API",
	Description: "Go Rest API with SQL SERVER DB",
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