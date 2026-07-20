success and error response
password util (hashing and comparing passwords)
jwt
signup user

curl -X POST http://localhost:8081/user/register 
-H "Content-Type: application/json" 
-d '{
    "username":"mimi",
    "email":"mimi@c.c",
    "password":"admin"
}'