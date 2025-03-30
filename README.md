# Holiday Management API

This is a simple RESTful API built with **Golang, Echo framework**, and **CouchDB** for managing holidays. It allows you to **add, retrieve, update, and delete** holiday records.

## Features
- Add a new holiday
- Get all holidays
- Get a specific holiday by ISO date
- Update an existing holiday
- Delete a specific holiday
- Delete all holidays

## Prerequisites

Ensure you have the following installed:
- **Go** (1.18+ recommended)
- **CouchDB** (running locally or remotely)
- **Postman** or `cURL` for testing API

## Setup & Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/N-Jangra/cc.git
   cd cc
   ```

2. **Configure the Database:**
   - Update `.env` with your CouchDB credentials:
     ```env
     DB_HOST=localhost
     DB_PORT=5984
     DB_USER=USERNAME
     DB_PASSWORD=PASSWORD
     DB_NAME=DATABASE_NAME
     apikey=<your-apikey>
     ```

3. **Install dependencies:**
   ```sh
   go mod tidy
   ```

4. **Run the API server:**
   ```sh
   go run main.go
   ```
   The server should start at `http://localhost:8080`

## API Endpoints

### **1. Add a New Holiday**
- **Endpoint:** `POST /n`
- **Query Parameters:**
  | Parameter    | Type   | Description |
  |-------------|--------|-------------|
  | Name        | string | Holiday name |
  | iso_date    | string | Holiday ISO date (YYYY-MM-DD) |
  | international | bool | Whether it's an international holiday |
- **cURL Example:**
  ```sh
  curl -X POST "http://localhost:8080/n?Name=New%20Year&iso_date=2025-01-01&international=true"
  ```

### **2. Get All Holidays**
- **Endpoint:** `GET /ga`
- **cURL Example:**
  ```sh
  curl -X GET "http://localhost:8080/ga"
  ```

### **3. Get a Holiday by ISO Date**
- **Endpoint:** `GET /g/:iso_date`
- **Example:** `GET /g/2025-01-01`
- **cURL Example:**
  ```sh
  curl -X GET "http://localhost:8080/g/2025-01-01"
  ```

### **4. Update a Holiday**
- **Endpoint:** `PUT /u/:id`
- **cURL Example:**
  ```sh
  curl -X PUT "http://localhost:8080/u/01" -H "Content-Type: application/json" -d '{
      "name": "Updated Holiday",
      "iso_date": "2025-01-01",
      "international": false
  }'
  ```

### **5. Delete a Specific Holiday**
- **Endpoint:** `DELETE /d/:iso_date`
- **cURL Example:**
  ```sh
  curl -X DELETE "http://localhost:8080/d/2025-01-01"
  ```

### **6. Delete All Holidays**
- **Endpoint:** `DELETE /da`
- **cURL Example:**
  ```sh
  curl -X DELETE "http://localhost:8080/da"
  ```

## Running Tests
You can test the API using **Postman** or the provided `cURL` commands.

## License
This project is open-source and free to use under the **MIT License**.

## Contact 

For any questions or suggestions, feel free to reach out:

    Author: Nitin
    Email: itznitinjangra@gmail.com
    GitHub: n-jangra


