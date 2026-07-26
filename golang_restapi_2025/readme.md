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

Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InpnIiwiZXhwIjoxNzg1MTY2NDA5LCJpc3MiOiJnb3RlbXAifQ.roROQ9qSpc2aVxNAcS6EKKbltSRSyarGEFU2OpL9w70

Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InpnIiwiZXhwIjoxNzg1MTY2NDc0LCJpc3MiOiJnb3RlbXAifQ.b60mVX8It0jGTZPARDDNfOqWUxDqmfEO0ONVjyFs5AU