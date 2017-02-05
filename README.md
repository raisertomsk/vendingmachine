# Vending Machine

This is vending machine backend with JSON REST API.

## The user-side algorithm:

### 1. Ask for a list of available products

URL: http://host/list

Response:

```json
{
    "products": {
        "product1": {
            "name": "Product 1",
            "quantity": 2,
            "price": 50
        },
        "product2": {
            "name": "Product 2",
            "quantity": 5,
            "price": 12.5
        },
        ...
        "productN": {
            "name": "Product X",
            "quantity": 8,
            "price": 48
        }
    }
}
```

### 2. Make payment for a product:

URL: http://host/get

```json
{
    "product": "productX",
    "quantity": 2,
    "payment": 100
}
```

Response OK:

```json
{
    "error": false,
    "change": 22
}
```

Response wrong payment:

```json
{
    "error": true,
    "error_text": "Insufficient funds",
    "cash_back": 100
}
```

Response wrong quantity:

```json
{
    "error": true,
    "error_text": "Wrong quantity",
    "cash_back": 50
}
```

## The admin-side algorithm:

1. Add a product

URL: http://host/add

```json
{
    "name": "product123",
    "full_name": "Product 123",
    "quantity": 10,
    "price": 20
}
```

Fields quantity and price may be skipped.

2. Update a product

URL: http://host/update

```json
{
    "name": "product123",
    "full_name": "Product 123",
    "quantity": 10,
    "price": 20
}
```
Fields full_name, quantity and price may be skipped.

3. Delete a product

URL: http://host/delete/product1

Error reporting for an admin-side:

```json
{
    "error": true,
    "error_text": "Error description"
}
```