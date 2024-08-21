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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "1478488313@qq.com"
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
        "/auth/login": {
            "post": {
                "description": "输入` + "`" + `用户名-密码` + "`" + `以登录",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.AuthLoginResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "输入` + "`" + `用户名-密码-确认密码` + "`" + `以注册",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "确认密码",
                        "name": "confirm_password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.AuthRegisterResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/avatar": {
            "post": {
                "security": [
                    {
                        "AccessToken": []
                    },
                    {
                        "RefreshToken": []
                    }
                ],
                "description": "上传头像，文件类型支持jpg(jpeg)、png",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "上传用户头像",
                "parameters": [
                    {
                        "type": "file",
                        "description": "头像",
                        "name": "avatar",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.UploadAvatarResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    },
                    {
                        "RefreshToken": []
                    }
                ],
                "description": "可查看信息包括：用户名、昵称、用户状态、邮箱、头像",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "查看用户信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.ShowUserInfoResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "AccessToken": []
                    },
                    {
                        "RefreshToken": []
                    }
                ],
                "description": "可修改昵称和邮箱",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "修改用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "昵称",
                        "name": "nickname",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.UserInfoUpdateResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/password": {
            "put": {
                "security": [
                    {
                        "AccessToken": []
                    },
                    {
                        "RefreshToken": []
                    }
                ],
                "description": "输入` + "`" + `原密码-新密码-确认密码` + "`" + `以更改密码",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "更改用户密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "原密码",
                        "name": "origin_password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "新密码",
                        "name": "new_password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "确认密码",
                        "name": "confirm_password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.UserPasswordChangeResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/e.Code"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "e.Code": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4,
                5,
                6,
                7,
                8,
                9,
                10,
                11,
                12,
                13,
                14,
                15,
                16
            ],
            "x-enum-comments": {
                "ErrorAccountInvalid": "用户名或密码错误",
                "ErrorContextValue": "上下文值传递错误",
                "ErrorCreateUser": "创建用户错误",
                "ErrorEncryptMoney": "金额加密错误",
                "ErrorEncryptPassword": "密码加密错误",
                "ErrorFileError": "文件错误",
                "ErrorGenerateToken": "token生成错误",
                "ErrorGetUserByID": "根据id获取用户失败",
                "ErrorIncorrectPassword": "密码错误",
                "ErrorOSSUploadError": "OSS文件上传错误",
                "ErrorUpdateUser": "更新用户失败",
                "ErrorUploadAvatar": "头像上传错误",
                "ErrorUploadFile": "文件上传错误",
                "ErrorUserExists": "用户已存在",
                "InvalidParams": "参数错误",
                "Success": "响应成功",
                "UnknownError": "未知错误"
            },
            "x-enum-varnames": [
                "Success",
                "InvalidParams",
                "UnknownError",
                "ErrorUserExists",
                "ErrorEncryptPassword",
                "ErrorEncryptMoney",
                "ErrorCreateUser",
                "ErrorAccountInvalid",
                "ErrorGetUserByID",
                "ErrorUpdateUser",
                "ErrorIncorrectPassword",
                "ErrorUploadAvatar",
                "ErrorGenerateToken",
                "ErrorContextValue",
                "ErrorUploadFile",
                "ErrorFileError",
                "ErrorOSSUploadError"
            ]
        },
        "response.AuthLoginResp": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "avatar": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "response.AuthRegisterResp": {
            "type": "object"
        },
        "response.ShowUserInfoResp": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.UploadAvatarResp": {
            "type": "object"
        },
        "response.UserInfoUpdateResp": {
            "type": "object"
        },
        "response.UserPasswordChangeResp": {
            "type": "object"
        }
    },
    "securityDefinitions": {
        "AccessToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "RefreshToken": {
            "type": "apiKey",
            "name": "X-Refresh-Token",
            "in": "header"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "gin-mall",
	Description:      "gin-mall API Documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
