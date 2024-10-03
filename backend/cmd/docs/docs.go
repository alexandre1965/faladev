// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Suporte da API Faladev",
            "url": "http://www.faladev.com/support",
            "email": "support@faladev.com"
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
        "/": {
            "get": {
                "description": "Renders the form.html page to display user information form.",
                "produces": [
                    "text/html"
                ],
                "summary": "Render form page",
                "responses": {
                    "200": {
                        "description": "HTML content of the form page",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/callback": {
            "get": {
                "description": "Exchange code for token and save it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Handle OAuth2 callback",
                "parameters": [
                    {
                        "type": "string",
                        "description": "State token",
                        "name": "state",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization code",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "303": {
                        "description": "Redirects to /"
                    },
                    "400": {
                        "description": "State token doesn't match",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Unable to retrieve or save token",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/event": {
            "post": {
                "description": "Handles event creation, guest addition, email sending, and redirects to Google Meet link.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Handle event creation and interaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the student",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Email of the student",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Phone number of the student",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "303": {
                        "description": "Redirects to Google Meet link",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "No Google Meet link available or other errors",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI Specification",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Faladev API",
	Description:      "Esta é uma API exemplo do Faladev, que integra com o Google Calendar e o Gmail.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}