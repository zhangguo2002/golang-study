success and error response
password util (hashing and comparing passwords)
jwt
signup user

curl -X POST http://localhost:8081/users/register 
-H "Content-Type: application/json" 
-d '{
    "username":"mimi",
    "email":"mimi@c.c",
    "password":"admin"
}'


Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InpnIiwiZXhwIjoxNzg1MDY3ODMzLCJpc3MiOiJnb3RlbXAifQ.34OWfrC2XAQQTbeyRkJar7Z2nMclLK0NnVuk_jPM69c