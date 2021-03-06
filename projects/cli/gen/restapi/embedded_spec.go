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
  "swagger": "2.0",
  "info": {
    "title": "Fish History",
    "version": "1.0.0"
  },
  "paths": {
    "/history/chunk": {
      "post": {
        "produces": [
          "application/json"
        ],
        "operationId": "chunkHistory",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ChunkHistoryParams"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "History blocks filtered.",
            "schema": {
              "$ref": "#/definitions/HistoryBlocks"
            }
          },
          "400": {
            "description": "Bad request."
          }
        }
      }
    },
    "/timezone": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "getTimezone",
        "responses": {
          "200": {
            "description": "Timezone used in server.",
            "schema": {
              "$ref": "#/definitions/TimezoneResponse"
            }
          },
          "400": {
            "description": "Bad request."
          }
        }
      }
    }
  },
  "definitions": {
    "ChunkHistoryParams": {
      "type": "object",
      "properties": {
        "filter": {
          "$ref": "#/definitions/TimeRangeFilter"
        },
        "limit": {
          "description": "Limit number of results in a single chunk. Undefined means +inf.",
          "type": "integer",
          "x-nullable": true
        },
        "unit": {
          "$ref": "#/definitions/TimeUnit"
        }
      }
    },
    "History": {
      "type": "object",
      "required": [
        "when",
        "cmd"
      ],
      "properties": {
        "cmd": {
          "type": "string"
        },
        "paths": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "string"
          }
        },
        "when": {
          "$ref": "#/definitions/Timestamp"
        }
      }
    },
    "HistoryBlock": {
      "type": "object",
      "required": [
        "count",
        "from",
        "to"
      ],
      "properties": {
        "count": {
          "type": "integer"
        },
        "from": {
          "$ref": "#/definitions/Timestamp"
        },
        "histories": {
          "type": "array",
          "minItems": 1,
          "items": {
            "$ref": "#/definitions/History"
          }
        },
        "to": {
          "$ref": "#/definitions/Timestamp"
        }
      }
    },
    "HistoryBlocks": {
      "type": "object",
      "required": [
        "unit"
      ],
      "properties": {
        "blocks": {
          "type": "array",
          "minItems": 1,
          "items": {
            "$ref": "#/definitions/HistoryBlock"
          }
        },
        "unit": {
          "$ref": "#/definitions/TimeUnit"
        }
      }
    },
    "TimeRangeFilter": {
      "type": "object",
      "properties": {
        "from": {
          "$ref": "#/definitions/TimestampNullable"
        },
        "to": {
          "$ref": "#/definitions/TimestampNullable"
        }
      }
    },
    "TimeUnit": {
      "type": "string",
      "default": "all",
      "enum": [
        "all",
        "year",
        "month",
        "week",
        "day",
        "hour",
        "minute",
        "second"
      ]
    },
    "Timestamp": {
      "description": "Timestamp in unixtime seconds.",
      "type": "integer",
      "format": "int64"
    },
    "TimestampNullable": {
      "description": "Timestamp in unixtime seconds.",
      "type": "integer",
      "format": "int64",
      "x-nullable": true
    },
    "TimezoneResponse": {
      "type": "object",
      "required": [
        "name",
        "offset"
      ],
      "properties": {
        "name": {
          "type": "string"
        },
        "offset": {
          "description": "Timezone offset in seconds.",
          "type": "integer"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Fish History",
    "version": "1.0.0"
  },
  "paths": {
    "/history/chunk": {
      "post": {
        "produces": [
          "application/json"
        ],
        "operationId": "chunkHistory",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ChunkHistoryParams"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "History blocks filtered.",
            "schema": {
              "$ref": "#/definitions/HistoryBlocks"
            }
          },
          "400": {
            "description": "Bad request."
          }
        }
      }
    },
    "/timezone": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "getTimezone",
        "responses": {
          "200": {
            "description": "Timezone used in server.",
            "schema": {
              "$ref": "#/definitions/TimezoneResponse"
            }
          },
          "400": {
            "description": "Bad request."
          }
        }
      }
    }
  },
  "definitions": {
    "ChunkHistoryParams": {
      "type": "object",
      "properties": {
        "filter": {
          "$ref": "#/definitions/TimeRangeFilter"
        },
        "limit": {
          "description": "Limit number of results in a single chunk. Undefined means +inf.",
          "type": "integer",
          "x-nullable": true
        },
        "unit": {
          "$ref": "#/definitions/TimeUnit"
        }
      }
    },
    "History": {
      "type": "object",
      "required": [
        "when",
        "cmd"
      ],
      "properties": {
        "cmd": {
          "type": "string"
        },
        "paths": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "string"
          }
        },
        "when": {
          "$ref": "#/definitions/Timestamp"
        }
      }
    },
    "HistoryBlock": {
      "type": "object",
      "required": [
        "count",
        "from",
        "to"
      ],
      "properties": {
        "count": {
          "type": "integer"
        },
        "from": {
          "$ref": "#/definitions/Timestamp"
        },
        "histories": {
          "type": "array",
          "minItems": 1,
          "items": {
            "$ref": "#/definitions/History"
          }
        },
        "to": {
          "$ref": "#/definitions/Timestamp"
        }
      }
    },
    "HistoryBlocks": {
      "type": "object",
      "required": [
        "unit"
      ],
      "properties": {
        "blocks": {
          "type": "array",
          "minItems": 1,
          "items": {
            "$ref": "#/definitions/HistoryBlock"
          }
        },
        "unit": {
          "$ref": "#/definitions/TimeUnit"
        }
      }
    },
    "TimeRangeFilter": {
      "type": "object",
      "properties": {
        "from": {
          "$ref": "#/definitions/TimestampNullable"
        },
        "to": {
          "$ref": "#/definitions/TimestampNullable"
        }
      }
    },
    "TimeUnit": {
      "type": "string",
      "default": "all",
      "enum": [
        "all",
        "year",
        "month",
        "week",
        "day",
        "hour",
        "minute",
        "second"
      ]
    },
    "Timestamp": {
      "description": "Timestamp in unixtime seconds.",
      "type": "integer",
      "format": "int64"
    },
    "TimestampNullable": {
      "description": "Timestamp in unixtime seconds.",
      "type": "integer",
      "format": "int64",
      "x-nullable": true
    },
    "TimezoneResponse": {
      "type": "object",
      "required": [
        "name",
        "offset"
      ],
      "properties": {
        "name": {
          "type": "string"
        },
        "offset": {
          "description": "Timezone offset in seconds.",
          "type": "integer"
        }
      }
    }
  }
}`))
}
