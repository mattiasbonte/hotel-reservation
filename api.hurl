GET http://127.0.0.1:3333/api/v1/users

HTTP 200
[Asserts]
jsonpath "$.[0].firstName" == "ronny"
jsonpath "$..firstName" == "ronny"

POST http://127.0.0.1:3333/api/v1/users
{
    "firstName": "ronny",
    "lastName": "watercool",
    "email": "ronny@bonny.com",
    "password": "testing"
}

[Captures]
theBody: body


