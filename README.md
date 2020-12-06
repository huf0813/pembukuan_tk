# pembukuan_tk API

PembukuanTK is an API for managing your store. Through this API, we hope that small business will grow in a positive
trajectory. This is so cool, isn't it?

## <b>Architecture</b>

![ARCH](https://raw.githubusercontent.com/bxcodec/go-clean-arch/master/clean-arch.png)
We are <b>different</b>. Using this Clean Arch by Uncle Bob Martin, we have a chance to create multiple data source and
multiple delivery. Right now, we provide blazing fast <b>DB SQLITE</b> and delivered by <b>JSON</b>. Also, by using <b>
GO + MUX Router Helper + JSON Web Token</b>, we can create secure endpoints without excuses. Anyway, Thx
to [bxcodec](https://github.com/bxcodec/go-clean-arch) has explained this examples.

## <b>How to Run</b>

1. Simply run this command

```
go run main.go
```

2. Want to see our deployment on server? go check out this [link](https://pembukuantk.herokuapp.com/)
3. Upss :scream:, want to see more? come and send us by issues

## <b>Meet the actors</b>

```
1. User
2. Admin 
3. Customer
```

## <b>API without JWT Token</b>

1. Login

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
    "token": "auth token goes here"
  }
}
```

2. Token Checking

```
POST -> /token/expired
header : null
payload :
{
	"token" : "input your auth token here"
}
response : 
{
  "status": "success",
  "message": "token is valid",
  "data": null
}
```

## <b>API USER with JWT Token</b>

1. Get All Products

```
GET -> /products
header : user token
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

2. Get All Customers

```
GET -> /customers
header : user token
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

3. Get All Invoices

```
GET -> /invoice
header : user token
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

4. Get Invoice Detail By ID

```
POST -> /invoice/detail
header : user token
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

5. Get Statistics By Year

```
POST -> /statistics
header : user token
payload : 
{
	"year": "2019"
}
response : 
{
  "status": "success",
  "message": "",
  "data": [
    {
      "year_and_month": "2019-01",
      "profit": 0
    },
    {
      "year_and_month": "2019-02",
      "profit": 0
    },
    {
      "year_and_month": "2019-03",
      "profit": 0
    },
    {
      "year_and_month": "2019-04",
      "profit": 0
    },
    {
      "year_and_month": "2019-05",
      "profit": 0
    },
    {
      "year_and_month": "2019-06",
      "profit": 0
    },
    {
      "year_and_month": "2019-07",
      "profit": 0
    },
    {
      "year_and_month": "2019-08",
      "profit": 0
    },
    {
      "year_and_month": "2019-09",
      "profit": 0
    },
    {
      "year_and_month": "2019-10",
      "profit": 0
    },
    {
      "year_and_month": "2019-11",
      "profit": 0
    },
    {
      "year_and_month": "2019-12",
      "profit": 0
    }
  ]
}
```

6. Input New Invoice With Products

```
POST -> /invoice
header : user token
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

7. Input New Stock for Product

```
POST -> /products/stock
header : user token
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

8. Input New Product

```
POST -> /products
header : user token
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

9. Input New Customer

```
POST -> /customers
header : user token
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

10. Edit Customer By ID

```
PUT -> /customers
header : user token
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

11. Edit Product By ID

```
PUT -> /products
header : user token
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

## <b>API ADMIN with JWT Token</b>

1. Get All Products

```
GET -> /admin/products
header : admin token
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

2. Get All Customers

```
GET -> /admin/customers
header : admin token
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

3. Get All Invoices

```
GET -> /admin/invoice
header : admin token
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

4. Get Invoice Detail By ID

```
POST -> /admin/invoice/detail
header : admin token
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

5. Get Statistics By Year

```
POST -> /admin/statistics
header : admin token
payload : 
{
	"year": "2019"
}
response : 
{
  "status": "success",
  "message": "",
  "data": [
    {
      "year_and_month": "2019-01",
      "profit": 0
    },
    {
      "year_and_month": "2019-02",
      "profit": 0
    },
    {
      "year_and_month": "2019-03",
      "profit": 0
    },
    {
      "year_and_month": "2019-04",
      "profit": 0
    },
    {
      "year_and_month": "2019-05",
      "profit": 0
    },
    {
      "year_and_month": "2019-06",
      "profit": 0
    },
    {
      "year_and_month": "2019-07",
      "profit": 0
    },
    {
      "year_and_month": "2019-08",
      "profit": 0
    },
    {
      "year_and_month": "2019-09",
      "profit": 0
    },
    {
      "year_and_month": "2019-10",
      "profit": 0
    },
    {
      "year_and_month": "2019-11",
      "profit": 0
    },
    {
      "year_and_month": "2019-12",
      "profit": 0
    }
  ]
}
```

6. Input New Invoice With Products

```
POST -> /admin/invoice
header : admin token
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

7. Input New Stock for Product

```
POST -> /admin/products/stock
header : admin token
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

8. Input New Product

```
POST -> /admin/products
header : admin token
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

9. Input New Customer

```
POST -> /admin/customers
header : admin token
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

10. Edit Customer By ID

```
PUT -> /admin/customers
header : admin token
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

11. Edit Product By ID

```
PUT -> /admin/products
header : admin token
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

12. Get All Users

```
GET -> /admin/users
header : admin token
payload : null
response : 
{
  "status": "success",
  "message": "",
  "data": [
    {
      "id": 1,
      "username": "jo",
      "password": "password encrypted"
      "user_type_id": 2
    },
    {
      "id": 3,
      "username": "huhu",
      "password": "password encrypted"
      "user_type_id": 2
    }
  ]
}
```

13. Input New User

```
POST -> /admin/users
header : admin token
payload : 
{
	"username" : "har",
	"password" : "har"
}
response : 
{
  "status": "success",
  "message": "",
  "data": {
    "id": 3,
    "username": "har",
    "password": "password encrpyted"
    "user_type_id": 2
  }
}
```

14. Edit User By ID

```
PUT -> /admin/users
header : admin token
payload : 
{
	"id": 3,
	"username": "huhu",
	"password": "huhu"
}
response : 
{
  "status": "success",
  "message": "",
  "data": {
    "id": 3,
    "username": "huhu",
    "password": "password encrypted"
    "user_type_id": 2
  }
}
```

