// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the go-openapi swagger example",
    "title": "go-openapi",
    "version": "1.0.0"
  },
  "basePath": "/v2",
  "paths": {
    "/healthz": {
      "get": {
        "tags": [
          "health"
        ],
        "summary": "Health check",
        "responses": {
          "default": {
            "description": "OK"
          }
        }
      }
    },
    "/user": {
      "post": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Create user",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Error creating user"
          },
          "409": {
            "description": "User exists"
          },
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/user/login": {
      "get": {
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs user into the system",
        "operationId": "loginUser",
        "parameters": [
          {
            "type": "string",
            "description": "The user name for login",
            "name": "username",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "The password for login in clear text",
            "name": "password",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "string"
            },
            "headers": {
              "X-Expires-After": {
                "type": "string",
                "format": "date-time",
                "description": "date in UTC when token expires"
              },
              "X-Rate-Limit": {
                "type": "integer",
                "format": "int32",
                "description": "calls per hour allowed by the user"
              }
            }
          },
          "400": {
            "description": "Invalid username/password supplied"
          },
          "500": {
            "description": "Error creating user"
          }
        }
      }
    },
    "/user/logout": {
      "get": {
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs out current logged in user session",
        "operationId": "logoutUser",
        "responses": {
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/user/{username}": {
      "get": {
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Get user by user name",
        "operationId": "getUserByName",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be fetched. Use user1 for testing. ",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid username supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "put": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Updated user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "name that need to be updated",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid user supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "delete": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Delete user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be deleted",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "User deleted"
          },
          "400": {
            "description": "Invalid username supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      }
    }
  },
  "definitions": {
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "lastName": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "phone": {
          "type": "string"
        },
        "userStatus": {
          "description": "User Status",
          "type": "integer",
          "format": "int32"
        },
        "username": {
          "type": "string"
        }
      },
      "xml": {
        "name": "User"
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    },
    "user_auth": {
      "type": "oauth2",
      "flow": "implicit",
      "authorizationUrl": "http://petstore.swagger.io/oauth/dialog",
      "scopes": {
        "read:pets": "read your pets",
        "write:pets": "modify pets in your account"
      }
    }
  },
  "tags": [
    {
      "description": "Operations about user",
      "name": "user",
      "externalDocs": {
        "description": "Find out more about our store",
        "url": "http://swagger.io"
      }
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the go-openapi swagger example",
    "title": "go-openapi",
    "version": "1.0.0"
  },
  "basePath": "/v2",
  "paths": {
    "/healthz": {
      "get": {
        "tags": [
          "health"
        ],
        "summary": "Health check",
        "responses": {
          "default": {
            "description": "OK"
          }
        }
      }
    },
    "/user": {
      "post": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "user"
        ],
        "summary": "Create user",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Error creating user"
          },
          "409": {
            "description": "User exists"
          },
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/user/login": {
      "get": {
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs user into the system",
        "operationId": "loginUser",
        "parameters": [
          {
            "type": "string",
            "description": "The user name for login",
            "name": "username",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "The password for login in clear text",
            "name": "password",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "string"
            },
            "headers": {
              "X-Expires-After": {
                "type": "string",
                "format": "date-time",
                "description": "date in UTC when token expires"
              },
              "X-Rate-Limit": {
                "type": "integer",
                "format": "int32",
                "description": "calls per hour allowed by the user"
              }
            }
          },
          "400": {
            "description": "Invalid username/password supplied"
          },
          "500": {
            "description": "Error creating user"
          }
        }
      }
    },
    "/user/logout": {
      "get": {
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs out current logged in user session",
        "operationId": "logoutUser",
        "responses": {
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/user/{username}": {
      "get": {
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "user"
        ],
        "summary": "Get user by user name",
        "operationId": "getUserByName",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be fetched. Use user1 for testing. ",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid username supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "put": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "user"
        ],
        "summary": "Updated user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "name that need to be updated",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid user supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "delete": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "user"
        ],
        "summary": "Delete user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be deleted",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "User deleted"
          },
          "400": {
            "description": "Invalid username supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      }
    }
  },
  "definitions": {
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "lastName": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "phone": {
          "type": "string"
        },
        "userStatus": {
          "description": "User Status",
          "type": "integer",
          "format": "int32"
        },
        "username": {
          "type": "string"
        }
      },
      "xml": {
        "name": "User"
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    },
    "user_auth": {
      "type": "oauth2",
      "flow": "implicit",
      "authorizationUrl": "http://petstore.swagger.io/oauth/dialog",
      "scopes": {
        "read:pets": "read your pets",
        "write:pets": "modify pets in your account"
      }
    }
  },
  "tags": [
    {
      "description": "Operations about user",
      "name": "user",
      "externalDocs": {
        "description": "Find out more about our store",
        "url": "http://swagger.io"
      }
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
}
