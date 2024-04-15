// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

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
        "/auth": {
            "get": {
                "description": "check email address before login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account_management"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "email address",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.OK"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account_management"
                ],
                "parameters": [
                    {
                        "description": "Login",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.LoginResponse"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "get": {
                "description": "logout user from the service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account_management"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.OK"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Register account for admin.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account_management"
                ],
                "parameters": [
                    {
                        "description": "Sign Up",
                        "name": "sign_up",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SignUpResponse"
                        }
                    }
                }
            }
        },
        "/categories": {
            "get": {
                "description": "Get categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "parameters": [
                    {
                        "type": "number",
                        "description": "limit number",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.CategoriesList"
                        }
                    }
                }
            },
            "post": {
                "description": "Insert new category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product_management"
                ],
                "parameters": [
                    {
                        "description": "Categories Request",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CategoriesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.LoginResponse"
                        }
                    }
                }
            }
        },
        "/common/healthcheck": {
            "get": {
                "description": "health check api server.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/newsletter": {
            "get": {
                "description": "Get Product Newsletter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit number of newsletter",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.NewsletterListResponse"
                        }
                    }
                }
            }
        },
        "/newsletters": {
            "post": {
                "description": "Create newsletter",
                "consumes": [
                    "multipart/form-data",
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product_management"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "image of newsletter",
                        "name": "img",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "subtitle",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "textcolor",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.OK"
                        }
                    }
                }
            }
        },
        "/oauth2/login": {
            "get": {
                "description": "Auth0 Login form.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "get product information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit number of products",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.ProductsListResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new product",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product_management"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "image of product",
                        "name": "img",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "available",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "categoryID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "price",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "supplierID",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.OK"
                        }
                    }
                }
            }
        },
        "/products/img/:id": {
            "post": {
                "description": "Insert new product image",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product_management"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "image of product",
                        "name": "img",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.OK"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    }
                }
            }
        },
        "/suppliers": {
            "get": {
                "description": "get suppliers information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product_management"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit number of suppliers",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SuppliersListResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "insert new suppliers information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product_management"
                ],
                "parameters": [
                    {
                        "description": "Suppliers Request",
                        "name": "SuppliersRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SuppliersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.OK"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "get information for users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account_management"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.UserInfo"
                        }
                    }
                }
            },
            "put": {
                "description": "update information for users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account_management"
                ],
                "parameters": [
                    {
                        "description": "Update User",
                        "name": "UserInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.OK"
                        }
                    }
                }
            }
        },
        "/users/image": {
            "put": {
                "description": "update information for users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account_management"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.OK"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Categories": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "domain.CategoriesList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Categories"
                    }
                }
            }
        },
        "domain.CategoriesRequest": {
            "type": "object",
            "required": [
                "description",
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "domain.Error": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        },
        "domain.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.LoginResponse": {
            "type": "object",
            "required": [
                "email",
                "success",
                "token"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "domain.NewsletterListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Newsletters"
                    }
                }
            }
        },
        "domain.Newsletters": {
            "type": "object",
            "required": [
                "description",
                "id",
                "image",
                "subtitle",
                "textcolor",
                "title",
                "type"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "subtitle": {
                    "type": "string"
                },
                "textcolor": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "domain.OK": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        },
        "domain.Products": {
            "type": "object",
            "properties": {
                "available": {
                    "type": "string"
                },
                "category_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "supplier_id": {
                    "type": "string"
                }
            }
        },
        "domain.ProductsListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Products"
                    }
                }
            }
        },
        "domain.SignUpRequest": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password",
                "phone_number"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "domain.SignUpResponse": {
            "type": "object",
            "required": [
                "msg",
                "success"
            ],
            "properties": {
                "msg": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "domain.Suppliers": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "domain.SuppliersListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Suppliers"
                    }
                }
            }
        },
        "domain.SuppliersRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "phone_number"
            ],
            "properties": {
                "city": {
                    "type": "string"
                },
                "district": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "ward": {
                    "type": "string"
                }
            }
        },
        "domain.UserInfo": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "id",
                "image",
                "last_name",
                "phone_number",
                "role",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.UserUpdate": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "id",
                "last_name",
                "phone_number",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swipe API Documentation",
	Description:      "This is a documentation for the Swipe API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
