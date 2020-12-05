# pembukuan_tk

## <b>Meet the actor</b>

```
1. User
2. Admin 
3. Customer
```

## <b>API without JWT Token</b>

```
POST -> /login
header : null
payload :
{
	"username": "jo",
	"password": "jo"
}
response : 
{
  "status": "success",
  "message": "",
  "data": {
    "username": "jo",
    "user_type_id": 2,
    "user_id": 1,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvIiwidXNlcl90eXBlX2lkIjoyLCJ1c2VyX2lkIjoxLCJleHAiOjE2MDcyMDYzMDB9.KWwfjI0HRFi-yeTiviemffAuS3ZX6sjktCX4lJQRT2c"
  }
}
```
```
POST -> /token/expired
header : null
payload :
{
	"token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvIiwidXNlcl90eXBlX2lkIjoyLCJ1c2VyX2lkIjoxLCJleHAiOjE2MDcyMDYzMDB9.KWwfjI0HRFi-yeTiviemffAuS3ZX6sjktCX4lJQRT2c"
}
response : 
{
  "status": "success",
  "message": "token is valid",
  "data": null
}
```

## <b>API USER with JWT Token</b>

```
GET -> /products
header : token
payload : null
response : 
{
  "status": "success",
  "message": "",
  "data": [
    {
      "id": 1,
      "name": "kopi janji jiwa",
      "price": "1000",
      "stock": 99
    }
  ]
}
```
```
GET -> /customers
header : token
payload : null
response : 
{
  "status": "success",
  "message": "",
  "data": [
    {
      "id": 1,
      "name": "randi julon",
      "phone": "082208220822",
      "email": "randi@gmail.com",
      "address": "malang"
    }
  ]
}
```
```
GET -> /invoice
header : token
payload : null
response : 
{
  "status": "success",
  "message": "",
  "data": [
    {
      "id": 1,
      "customer_name": "randi julon",
      "customer_phone": "082208220822",
      "customer_email": "randi@gmail.com",
      "customer_address": "malang",
      "total_invoice_price": "1000",
      "created_at": "1607178612",
      "updated_at": "1607178612",
      "products": [
        {
          "product_name": "kopi janji jiwa",
          "product_qty": "1",
          "product_price": "1000",
          "product_total_price": "1000"
        }
      ]
    }
  ]
}
```
```
POST -> /invoice/detail
header : token
payload : 
{
	"invoice_id": 1
}
response : 
{
  "status": "success",
  "message": "",
  "data": {
    "id": 1,
    "customer_name": "randi julon",
    "customer_phone": "082208220822",
    "customer_email": "randi@gmail.com",
    "customer_address": "malang",
    "total_invoice_price": "1000",
    "created_at": "1607178612",
    "updated_at": "1607178612",
    "products": [
      {
        "product_name": "kopi janji jiwa",
        "product_qty": "1",
        "product_price": "1000",
        "product_total_price": "1000"
      }
    ]
  }
}
```
```
POST -> /invoice
header : token
payload : 
{
	"customer_id": 1,
	"user_id": 1,
	"list_product": [
		{
			"product_id": 1,
			"qty": 1
		}
	]
}
response : 
{
  "status": "success",
  "message": "",
  "data": {
    "id": 1,
    "customer_id": 1,
    "user_id": 1
  }
}
```
```
POST -> /products/stock
header : token
payload : 
{
	"product_id" : 1,
	"quantity" : 100,
	"user_id" : 1
}
response : 
{
  "status": "success",
  "message": "",
  "data": {
    "id": 1,
    "product_id": 1,
    "quantity": 100,
    "user_id": 1
  }
}
```
```
POST -> /products
header : token
payload : 
{
	"name": "kopi janji jiwa",
	"price": "1000"
}
response : 
{
  "status": "success",
  "message": "",
  "data": {
    "id": 1,
    "name": "kopi janji jiwa",
    "price": "1000"
  }
}
```
```
POST -> /customers
header : token
payload : 
{
	"name": "randi julon",
	"phone": "082208220822",
	"email": "randi@gmail.com",
	"address": "malang"
}
response : 
{
  "status": "success",
  "message": "inserted successfully",
  "data": {
    "id": 1,
    "name": "randi julon",
    "phone": "082208220822",
    "email": "randi@gmail.com",
    "address": "malang"
  }
}
```
```
PUT -> /customers
header : token
payload : 
{
	"name": "randi edit haha",
	"phone": "082208220822",
	"email": "randi@gmail.com",
	"address": "malang",
	"id": 1
}
response : 
{
  "status": "success",
  "message": "inserted successfully",
  "data": {
    "id": 1,
    "name": "randi edit haha",
    "phone": "082208220822",
    "email": "randi@gmail.com",
    "address": "malang"
  }
}
```
```
PUT -> /products
header : token
payload : 
{
	"id": 1,
	"name": "kopi excelso",
	"price": "500000"
}
response : 
{
  "status": "success",
  "message": "",
  "data": {
    "id": 1,
    "name": "kopi excelso",
    "price": "500000"
  }
}
```