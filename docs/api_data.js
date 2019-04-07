define({ "api": [
  {
    "type": "post",
    "url": "/delete",
    "title": "delete values",
    "name": "delete_values",
    "group": "all",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "key",
            "description": "<p>key of the value</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"key\":\"dsad\"\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"key\": \"dsad\",\n   \"value\": \"ddd\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controller/delete.go",
    "groupTitle": "all"
  },
  {
    "type": "post",
    "url": "/get",
    "title": "get values",
    "name": "get_values",
    "group": "all",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "key",
            "description": "<p>key of the value</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n   \"key\": \"dsad\"\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"key\": \"dsad\",\n   \"value\": \"ddd\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controller/get.go",
    "groupTitle": "all"
  },
  {
    "type": "post",
    "url": "/put",
    "title": "put values",
    "name": "put_values",
    "group": "all",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "key",
            "description": "<p>key of the value</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "value",
            "description": "<p>value of the key</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n   \"key\": \"dsad\",\n   \"value\": \"ddd\"\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"key\": \"dsad\",\n   \"value\": \"ddd\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controller/put.go",
    "groupTitle": "all"
  },
  {
    "type": "post",
    "url": "/update",
    "title": "update values",
    "name": "update_values",
    "group": "all",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "key",
            "description": "<p>key of the value</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "value",
            "description": "<p>new value of the key</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n   \"key\": \"dsad\",\n   \"value\": \"ddd\"\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"key\": \"dsad\",\n   \"value\": \"ddd\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controller/update.go",
    "groupTitle": "all"
  }
] });
