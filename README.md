# Vending Machine

## To run

1. clone repository
1. go get github.com/julienschmidt/httprouter
2. go build
3. ./vendingmachine

Configuration option -c sets up a configuration file path.
As default it is ./config.json.

This is vending machine backend with JSON REST API.

## The user-side algorithm:

### 1. Ask for a list of available to buy products

URL: GET http://host/list

`curl -XGET http://localhost:12345/list`

Response:

```json
{
    "product1": {
        "title": "Product 1",
        "quantity": 2,
        "price": 50
    },
    "product2": {
        "title": "Product 2",
        "quantity": 5,
        "price": 12.5
    },
    ...
    "productN": {
        "title": "Product X",
        "quantity": 8,
        "price": 48
    }
}
```

### 2. To buy a product:

URL: http://host/get/:name

`curl -XGET http://localhost:12345/get/product1`

```json
{
    "title": "product X",
    "quantity": 2,
    "price": 100
}
```

Response OK:

```json
{
    "error_code": 0,
    "change": 22
}
```
change can be omitted

Response wrong quantity:

```json
{
    "error_code": number,
    "error_text": "Error description",
    "cash_back": 50
}
```

## The admin-side algorithm:

### 1. Add a product

URL: http://host/add/:name

`curl -XPUT http://localhost:12345/add/product1`

```json
{
    "title": "Product 123",
    "quantity": 10,
    "price": 20
}
```

Fields quantity and price may be omitted.

### 2. Update a product

URL: http://host/update/:name

`curl -XPOST http://localhost:12345/update/product1`

```json
{
    "title": "Product 123",
    "quantity": 10,
    "price": 20
}
```
Fields full_name, quantity and price may be skipped.

### 3. Delete a product

URL: http://host/delete/:name

`curl -XDELETE http://localhost:12345/delete/product1`

### Error reporting for an admin-side:

```json
{
    "error_code": number,
    "error_text": "Error description"
}
```