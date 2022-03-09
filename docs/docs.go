// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/bottle": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get string by Bottle ID",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "bottle"
                ],
                "summary": "Show an Bottle",
                "parameters": [
                    {
                        "type": "string",
                        "example": "bottle ID",
                        "name": "bottleId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetBottleResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "새로운 유리병을 생성합니다. \u003cbr /\u003e 유리병의 이름을 입력할 수 있습니다. 아무값 없이 요청하면 \"default\" 이름을 갖게 됩니다.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "bottle"
                ],
                "summary": "Create BOTTLE by userID",
                "parameters": [
                    {
                        "description": "Create Bottle request body",
                        "name": "title",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateBottleRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Bottle"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/bottle/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GET bottle list",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "bottle"
                ],
                "summary": "유리병 목록 조회",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Bottle"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/hplog": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "단일 기록 세부사항 조회",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "hplog"
                ],
                "summary": "기록 세부사항 조회",
                "parameters": [
                    {
                        "type": "string",
                        "example": "hplog ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.HpLog"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "유리병에 단일 로그 생성",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "hplog"
                ],
                "summary": "로그 생성",
                "parameters": [
                    {
                        "description": "create hplog",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateHplogRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateHplogResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/hplog/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "다중기록 조회 (유리병 단위)",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "hplog"
                ],
                "summary": "기록 목록 조회",
                "parameters": [
                    {
                        "type": "string",
                        "example": "bottle ID",
                        "name": "bottleId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.HpLog"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/test/login": {
            "post": {
                "description": "테스트 유저 로그인",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "test"
                ],
                "summary": "테스트 유저 로그인",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.loginResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.CreateBottleRequestBody": {
            "type": "object",
            "properties": {
                "maturityDate": {
                    "type": "string",
                    "example": "date JSON string"
                },
                "title": {
                    "type": "string",
                    "example": "bottle title (optional)"
                }
            }
        },
        "controllers.CreateHplogRequestBody": {
            "type": "object",
            "properties": {
                "bottleId": {
                    "type": "string",
                    "example": "bottle ID"
                },
                "text": {
                    "type": "string",
                    "example": "each log text"
                },
                "worth": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "controllers.CreateHplogResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "controllers.GetBottleResponse": {
            "type": "object",
            "properties": {
                "hplogList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.HpLog"
                    }
                },
                "maturityDate": {
                    "type": "string",
                    "example": "2022-03-04T03:16:49.767Z"
                },
                "title": {
                    "type": "string",
                    "example": "Bottle title"
                },
                "totalWorth": {
                    "type": "integer",
                    "example": 2300000
                }
            }
        },
        "controllers.loginResponseBody": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "token string (JWT)"
                }
            }
        },
        "entity.Bottle": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": ""
                },
                "maturityDate": {
                    "type": "string",
                    "example": ""
                },
                "title": {
                    "type": "string",
                    "example": "default"
                }
            }
        },
        "entity.HpLog": {
            "type": "object",
            "required": [
                "text",
                "worth"
            ],
            "properties": {
                "text": {
                    "type": "string"
                },
                "worth": {
                    "type": "integer"
                }
            }
        },
        "httputil.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080/",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
