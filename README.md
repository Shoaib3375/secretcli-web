# SecretCLI-Web

## Overview
SecretCLI-Web is a web application for securely managing sensitive information, such as passwords, personal notes, and account details. It allows users to register, login, create, list, and retrieve secret entries, with additional features for generating secure passwords.

### Register User
<details>

**Endpoint:** /auth/api/register

**Method:** POST

**Description:** Registers a new user in the system.

```
curl -X POST http://localhost:8080/auth/api/register \
     -H "Content-Type: application/json" \
     -d '{
          "name": "Test User",
          "email": "test5@example.com",
          "password": "test"
     }' 

```

</details>

---
### Login User
<details>

**Endpoint:** /auth/api/login

**Method:** POST

**Description:** Authenticates an existing user and provides a token for session management.

```
curl -X POST http://localhost:8080/auth/api/login \
     -H "Content-Type: application/json" \
     -d '{
          "email": "test1@example.com",
          "password": "test"
     }'
```

</details>

---
### Create Secret
<details>

**Endpoint:** /secret/api/create

**Method:** POST

**Description:** Creates a new secret entry. Requires an authentication token.

### Headers:
**Authorization:** Bearer \<token-here>

**Content-Type:** application/json

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

</details>

---
### List Secret
<details>

**Endpoint:** /secret/api/list

**Method:** GET

**Description:** Retrieves a list of secrets for the authenticated user.

### Headers:
**Authorization:** Bearer \<token-here>
```
curl -H "Authorization: Bearer <token-here>" -X GET http://localhost:8080/secret/api/list

```

```
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMwNTg1Njc1LCJ1c2VyX2lkIjozOX0.Nxc41WUhK5nBTaQjRAVT8W6J2NY0JPAW0f7lVWRbt7Q" -X GET http://localhost:8080/secret/api/list
```

</details>

---
### Generate Password
<details>

**Endpoint:** /secret/api/generatepassword

**Method:** POST

**Description:** Generates a secure password with specified length and character options.

### Headers:

**Authorization:** Bearer \<token-here>

**Content-Type:** application/json

### Request Parameters:
**length (int):** Desired password length.

**include_special_symbol (bool):** Whether to include special symbols.

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

</details>

---
### Get Secret Detail
<details>

**Endpoint:** /secret/api/secretdetail

**Method:** POST

**Description:** Retrieves detailed information of a specific secret by ID.

### Headers:
**Authorization:** Bearer \<token-here>

**Content-Type:** application/json

### Request Parameters:
**secret_id (int):** ID of the secret to retrieve.

```
curl -X 'POST' \
  'http://localhost:8080/secret/api/secretdetail' \
  -H 'accept: application/json' \
  -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMwNTg1NjAzLCJ1c2VyX2lkIjozOX0.4TQimAROxdk-_KdZLluz7hI32xzudGyi70GYe3-CYlQ' \
  -H 'Content-Type: application/json' \
  -d '{
  "secret_id": 16
}'
```

```
curl -X 'POST' \
  'http://localhost:8080/secret/api/secretdetail' \
  -H 'accept: application/json' \
  -H 'Authorization: <token>' \
  -H 'Content-Type: application/json' \
  -d '{
  "secret_id": 16
}'
```
</details>

---