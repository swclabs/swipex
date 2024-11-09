// Package article Code generated by swaggo/swag. DO NOT EDIT
package article

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/comment": {
            "get": {
                "description": "get all comments of product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of products",
                        "name": "product_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.Comment"
                        }
                    }
                }
            },
            "post": {
                "description": "create comment into products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "parameters": [
                    {
                        "description": "comment data request",
                        "name": "banner",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.Comment"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.OK"
                        }
                    }
                }
            }
        },
        "/news": {
            "get": {
                "description": "get news",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "news"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "category of news",
                        "name": "category",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "limit of cards carousel",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.News"
                        }
                    }
                }
            },
            "post": {
                "description": "create news",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "news"
                ],
                "parameters": [
                    {
                        "description": "news Request",
                        "name": "collection",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.NewsDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.CollectionUpload"
                        }
                    }
                }
            }
        },
        "/news/image": {
            "put": {
                "description": "update news image",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "news"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "image of news",
                        "name": "img",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "news identifier",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.OK"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.CardArticle": {
            "type": "object",
            "required": [
                "category",
                "title"
            ],
            "properties": {
                "category": {
                    "type": "string"
                },
                "content": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.CardContent"
                    }
                },
                "src": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dtos.CardContent": {
            "type": "object",
            "required": [
                "content"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "src": {
                    "type": "string"
                }
            }
        },
        "dtos.CollectionUpload": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "dtos.Comment": {
            "type": "object",
            "required": [
                "content",
                "dislike",
                "id",
                "level",
                "like",
                "rating",
                "username"
            ],
            "properties": {
                "content": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "dislike": {
                    "type": "integer"
                },
                "id": {
                    "description": "Position string   ` + "`" + `json:\"position\" validate:\"required\"` + "`" + `",
                    "type": "integer"
                },
                "level": {
                    "description": "0: parent, 1: child",
                    "type": "integer"
                },
                "like": {
                    "type": "integer"
                },
                "parent_id": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dtos.News": {
            "type": "object",
            "required": [
                "cards",
                "header"
            ],
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.CardArticle"
                    }
                },
                "header": {
                    "type": "string"
                }
            }
        },
        "dtos.NewsDTO": {
            "type": "object",
            "required": [
                "cards",
                "category",
                "header"
            ],
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.CardArticle"
                    }
                },
                "category": {
                    "type": "string"
                },
                "header": {
                    "type": "string"
                }
            }
        },
        "dtos.OK": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swipe Public API v0.0.1",
	Description:      "This is a documentation for the Swipe API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
