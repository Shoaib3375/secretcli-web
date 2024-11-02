# secretcli-web


### Register User
---

```
curl -X POST http://localhost:8080/auth/api/register \
     -H "Content-Type: application/json" \
     -d '{
          "name": "Test User",
          "email": "test5@example.com",
          "password": "test"
     }' 

```

### Login User
---

```
curl -X POST http://localhost:8080/auth/api/login \
     -H "Content-Type: application/json" \
     -d '{
          "email": "test1@example.com",
          "password": "test"
     }'
```


### Create Secret
---

```
curl -X POST http://localhost:8080/secret/api/create \
     -H "Authorization: Bearer <token-here>" \
     -H "Content-Type: application/json" \
     -d '{
          "title": "Test Secret 1",
          "username": "test1",
          "password": "test1",
          "note": "This is a test secret.",
          "email": "test1@example.com",
          "website": "https://test1.com"
     }'
```

```
curl -X POST http://localhost:8080/secret/api/create \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMwNTg1NjAzLCJ1c2VyX2lkIjozOX0.4TQimAROxdk-_KdZLluz7hI32xzudGyi70GYe3-CYlQ" \
     -H "Content-Type: application/json" \
     -d '{
          "title": "Test Secret 1",
          "username": "test1",
          "password": "test1",
          "note": "This is a test secret.",
          "email": "test1@example.com",
          "website": "https://test1.com"
     }'
```


### List Secret
---

```
curl -H "Authorization: Bearer <token-here>" -X GET http://localhost:8080/secret/api/list

```

```
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMwNTg1Njc1LCJ1c2VyX2lkIjozOX0.Nxc41WUhK5nBTaQjRAVT8W6J2NY0JPAW0f7lVWRbt7Q" -X GET http://localhost:8080/secret/api/list
```

### Generate Password

```
curl -X POST http://localhost:8080/secret/api/generatepassword \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMwNTg1NjAzLCJ1c2VyX2lkIjozOX0.4TQimAROxdk-_KdZLluz7hI32xzudGyi70GYe3-CYlQ" \
     -H "Content-Type: application/json" \
     -d '{
          "length": 10,
          "include_special_symbol": true
     }'
```

```
curl -X POST http://localhost:8080/secret/api/generatepassword \
     -H "Authorization: Bearer <token-here> \
     -H "Content-Type: application/json" \
     -d '{
          "length": 100,
          "include_special_symbol": true
     }'
```