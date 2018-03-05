define({ "api": [
  {
    "type": "get",
    "url": "/v1/health",
    "title": "GET Health",
    "version": "1.0.0",
    "description": "<p>Is the service healthy</p>",
    "name": "GetHealth",
    "group": "Health",
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK",
          "type": "json"
        }
      ]
    },
    "filename": "./route/route.go",
    "groupTitle": "Health"
  },
  {
    "type": "get",
    "url": "/v1/location",
    "title": "GET Location",
    "version": "1.0.0",
    "description": "<p>get the location of an ip address</p>",
    "name": "GetLocation",
    "group": "Location",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "ip",
            "description": "<p>ipv4 address</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "curl -i http://localhost:[port]/v1/location?ip=\"174.29.5.118\"",
        "type": "curl"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "    HTTP/1.1 200 OK\n    {\n\t   \t\t\"locationId\": 14288,\n\t\t\t\"country\": \"US\",\n\t\t\t\"region\": \"CO\",\n\t\t\t\"city\": \"Arvada\",\n\t\t\t\"postalCode\": \"\",\n\t\t\t\"latitude\": 39.8131,\n\t\t\t\"longitude\": -105.1257,\n\t\t\t\"metroCode\": \"751\",\n\t\t\t\"areaCode\": \"\"\n\t  }",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 400 Bad Request\n{\n  \"message\": \"error text\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./route/route.go",
    "groupTitle": "Location"
  },
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./docs/main.js",
    "group": "_app_docs_main_js",
    "groupTitle": "_app_docs_main_js",
    "name": ""
  }
] });
