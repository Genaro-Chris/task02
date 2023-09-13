
# Project Documentation

This document provides detailed information on how to use the REST API for the "Person" resource. Please refer to this documentation for setup instructions, request/response formats, sample API usage, and any known limitations or assumptions made during development.

> LIVE API Endpoint is https://hgnxbackend-prmpsmart.b4a.run/api
>
> For the testing script go to [main_test.go](https://github.com/Genaro-Chris/task02/blob/main/main_test.go)


## Table of Contents

- [Setup Instructions](#setup-instructions)
- [API Endpoints](#api-endpoints)
- [Request/Response Formats](#requestresponse-formats)
- [Sample API Usage](#sample-api-usage)
- [Known Limitations and Assumptions](#known-limitations-and-assumptions)

---

## Setup Instructions

Follow these steps to set up and run the API locally:

1. **Clone the Repository:**
   ```bash
git clone https://github.com/Genaro-Chris/task02
cd ./task02
   ```

2. **Install Dependencies:**
   ```bash 
go mod tidy
   ```

3. **Run the API:**
   ```bash
go run .
   ```

4. The API will be available locally at `http://127.0.0.1:8000`.

---

## API Endpoints

The API provides the following endpoints for CRUD operations on the "Person" resource:

- **Create a Person**:
  - **POST /api/**
  - Add a new person to the database.

- **Read a Person**:
  - **GET /api/{user_id}**
  - Retrieve details of a person by name.

- **Update a Person**:
  - **PUT /api/{user_id}**
  - Modify details of an existing person by name.

- **Delete a Person**:
  - **DELETE /api/{user_id}**
  - Remove a person from the database by name.

---

## Request/Response Formats

### Create a Person (POST /api/)

**Request Format:**
```python
import requests

api_url = "http://127.0.0.1:8000/api/"

data = {
   "name": "Miracle Apata"
}

response = requests.post(api_url, json=data)
print(response.json())
```

**Response Format (Success - 200):**
```json
{
    "ID":  1,
    "name": "Miracle Apata"
}
```


### Read a Person (GET /api/{user_id})

**Request Format:**
```python
import requests

api_url = "http://127.0.0.1:8000/api/1"

response = requests.get(api_url)
print(response.json())
```

**Response Format (Success- 200):**
```json
{
    "ID": "1",
    "name": "Miracle Apata",
    
}
```

**Response Format (Not Found - 404):**
```json
{
    "error": "No row found with that ID"
}
```

### Update a Person (PUT /api/{user_id})

**Request Format:**
```python
import requests

api_url = "http://127.0.0.1:8000/api/1"

data = {
    "name": "Orji Adekunle"
}

response = requests.put(api_url, json=data)
print(response.json())
```

**Response Format (Success - 200):**
```json
{
    "ID": 1,
    "name": "Orji Adekunle",
}
```

**Response Format (Not Found - 404):**
```json
{
  "error": "No row found with that ID"
}
```

### Delete a Person (DELETE /api/{user_id})

**Request Format:**
```python
import requests

api_url = "http://127.0.0.1:8000/api/1"

response = requests.delete(api_url)
print(response.json())
```

**Response Format (Not Found - 404):**
```json
{
  "error": "No row found with that ID"
}
```

---

## Sample API Usage

Here are some sample API usage scenarios with Python code examples:

1. **Create a Person:**
   ```python
   import requests

   api_url = "http://127.0.0.1:8000/api/"

   data = {
       "name": "Alice Johnson",
   }

   response = requests.post(api_url, json=data)
   print(response.json())
   ```

2. **Read a Person:**
   ```python
   import requests

   api_url = "http://127.0.0.1:8000/api/1"

   response = requests.get(api_url)
   print(response.json())
   ```

3. **Update a Person:**
   ```python
   import requests

   api_url = "http://127.0.0.1:8000/api/1"

   data = {
      "ID": 1,
      "name": "Orji Adekunle"
   }

   response = requests.put(api_url, json=data)
   print(response.json())
   ```

4. **Delete a Person:**
   ```python
   import requests

   api_url = "http://127.0.0.1:8000/api/1"

   response = requests.delete(api_url)
   print(response.json())
   ```

---

## Known Limitations and Assumptions

- This API uses a built-in database (sqlite3) for demonstration purposes.
- Input validation is handled by GORM library in this task. Implement more robust validation and error handling in a production-ready application.
- Authentication and authorization mechanisms are not implemented here. Ensure secure access to your API in a real-world scenario.
- This documentation assumes that you have successfully set up the API locally.
