# Project features

The project represents the backend of a customer relationship management (CRM) web application. As users interact with the app via some user interface, your server will support all of the functionalities:

All REST APIs (GET, POST, PATCH, DELETE)

### Getting a list of all customers

through a the /customers path

### Getting data for a single customer

through a /customers/{id} path

### Adding a customer

through a /customers path

### Updating a customer's information

through a /customers/{id} path

### Removing a customer

through a /customers/{id} path

# Api

- Getting a list of all customers

```url
url : GET localhost:3000/customers

```

- Response

```json
[
  {
    "id": "1",
    "name": "Tom",
    "role": "Software Engineer",
    "email": "tom@gmail.com",
    "phone": 123456892,
    "contacted": true
  },
  {
    "id": "2",
    "name": "Tim",
    "role": "Senior Software Engineer",
    "email": "tim@gmail.com",
    "phone": 213456892,
    "contacted": false
  },
  {
    "id": "3",
    "name": "Max",
    "role": "Product Manager",
    "email": "max@gmail.com",
    "phone": 231564892,
    "contacted": true
  },
  {
    "id": "4",
    "name": "Joe",
    "role": "Software Engineer",
    "email": "joe@gmail.com",
    "phone": 123456892,
    "contacted": true
  },
  {
    "id": "5",
    "name": "Ross",
    "role": "Civil Engineer",
    "email": "ross@gmail.com",
    "phone": 123456892,
    "contacted": true
  },
  {
    "id": "6",
    "name": "Adam",
    "role": "IT Support Engineer",
    "email": "adam@gmail.com",
    "phone": 123456892,
    "contacted": true
  }
]
```

- Getting a single customer

```url
url : GET localhost:3000/customers/{id}

```

- Response

```JSON
    {
        "id": "1",
        "name": "Tom",
        "role": "Software Engineer",
        "email": "tom@gmail.com",
        "phone": 123456892,
        "contacted": true
    }
```

- Adding a customer

```url
url : POST localhost:3000/customers

```

- Body

```JSON
  {
   "name": "John",
   "role": "Software Developer",
   "email": "johnEmail@anymail.com",
   "phone": 123456789,
   "contacted": true
   }
```

- Response

```JSON
[
    {
        "id": "1",
        "name": "Tom",
        "role": "Software Engineer",
        "email": "tom@gmail.com",
        "phone": 123456892,
        "contacted": true
    },
    {
        "id": "2",
        "name": "Tim",
        "role": "Senior Software Engineer",
        "email": "tim@gmail.com",
        "phone": 213456892,
        "contacted": false
    },
    {
        "id": "3",
        "name": "Max",
        "role": "Product Manager",
        "email": "max@gmail.com",
        "phone": 231564892,
        "contacted": true
    },
    {
        "id": "4",
        "name": "Joe",
        "role": "Software Engineer",
        "email": "joe@gmail.com",
        "phone": 123456892,
        "contacted": true
    },
    {
        "id": "5",
        "name": "Ross",
        "role": "Civil Engineer",
        "email": "ross@gmail.com",
        "phone": 123456892,
        "contacted": true
    },
    {
        "id": "6",
        "name": "Adam",
        "role": "IT Support Engineer",
        "email": "adam@gmail.com",
        "phone": 123456892,
        "contacted": true
    },
    {
        "id": "6cf26e2f-4d90-43b2-99e3-e4e89f2e891e",
        "name": "John",
        "role": "Software Developer",
        "email": "johnEmail@anymail.com",
        "phone": 123456789,
        "contacted": true
    }
]
```

- Updating a customer's information

```url

url : PATCH localhost:3000/customers/{id}
```

- Body

```json
        "name": "Jo",
        "role": "Software Developer",
        "contacted": true
```

- Response

```json
[
  {
    "id": "1",
    "name": "Tom",
    "role": "Software Engineer",
    "email": "tom@gmail.com",
    "phone": 123456892,
    "contacted": true
  },
  {
    "id": "2",
    "name": "Tim",
    "role": "Senior Software Engineer",
    "email": "tim@gmail.com",
    "phone": 213456892,
    "contacted": false
  },
  {
    "id": "3",
    "name": "Max",
    "role": "Product Manager",
    "email": "max@gmail.com",
    "phone": 231564892,
    "contacted": true
  },
  {
    "id": "4",
    "name": "Joe",
    "role": "Software Engineer",
    "email": "joe@gmail.com",
    "phone": 123456892,
    "contacted": true
  },
  {
    "id": "5",
    "name": "Ross",
    "role": "Civil Engineer",
    "email": "ross@gmail.com",
    "phone": 123456892,
    "contacted": true
  },
  {
    "id": "6",
    "name": "Adam",
    "role": "IT Support Engineer",
    "email": "adam@gmail.com",
    "phone": 123456892,
    "contacted": true
  },
  {
    "id": "6cf26e2f-4d90-43b2-99e3-e4e89f2e891e",
    "name": "Jo",
    "role": "Software Developer",
    "email": "johnEmail@anymail.com",
    "phone": 123456789,
    "contacted": true
  }
]
```

- Removing a customer

```url

url : DELETE localhost:3000/customers/{id}

```

- Response

```json
{
  "id": "6cf26e2f-4d90-43b2-99e3-e4e89f2e891e",
  "name": "Jo",
  "role": "Software Developer",
  "email": "johnEmail@anymail.com",
  "phone": 123456789,
  "contacted": true
}
```

## Data

Dummy data is in data.go file

# To Start the app

```
go run main.go data.go
```

## Packages used

```
encoding/json
```

The applications leverages the encoding/json package to parse JSON data.

```
io/ioutil
```

The application leverages the io/ioutil package to read I/O (e.g., request) data.

```
github.com/google/uuid
```

Google UUID GitHub project to create a unique id for the customers.

```
github.com/gorilla/mux
```

The application uses a router (e.g., gorilla/mux, http.ServeMux, etc.) that supports HTTP method-based routing and variables in URL paths.
