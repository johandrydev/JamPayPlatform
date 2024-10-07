# JAMPay payment platform API

## Service definition

This document describes the API endpoints for the JAMPay payment platform.

## API Endpoints

### Login

**Endpoint:** `POST /api/login`

**Description:** Logs in a user.

**Request Body:**

- `LoginInput` (JSON): The input data for logging in a user.

```json
{
  "email": "peach@mail.com",
  "password": "password"
}
```

**Responses:**

- `200 OK`: User logged in successfully.

```json
{
  "message": "User logged in successfully",
  "data": {
    "token": "eyJ..."
  }
}
```

- `400 Bad Request`: Invalid request body.
- `401 Unauthorized`: Invalid email or password.
- `500 Internal Server Error`: error trying to login, please try again later.

### Get Merchant

**Endpoint:** `GET /api/merchant/{merchantID}/`

**Description:** Retrieves a merchant information by its ID.

**Request Parameters:**

- `merchantID` (path parameter): The ID of the merchant to retrieve.
- `Authorization` (header): The token of the user logged in with a Bearer format.

**Responses:**

- `200 OK`: Merchant information retrieved successfully.

```json
{
  "message": "Merchant information retrieved successfully",
  "data": {
    "id": "8d384f61-c92b-4867-abfc-f78ed3b6ea15",
    "name": "Peach Bros",
    "email": "peach@mail.com",
    "bank_account": "123456789",
    "status": "VERIFIED"
  }
}
```

- `404 Not Found`: Merchant not found.
- `500 Internal Server Error`: error finding merchant, please try again later.

### Get All Payments by Merchant

**Endpoint:** `GET /api/merchant/{merchantID}/payments`

**Description:** Retrieves all payments made to a merchant.

**Request Parameters:**

- `merchantID` (path parameter): The ID of the merchant to retrieve payments for.
- `Authorization` (header): The token of the user logged in with a Bearer format. The user must be the merchant owner of the payments.

**Responses:**

- `200 OK`: Payments retrieved successfully.

```json
{
  "message": "Payments retrieved successfully",
  "data": [
    {
      "id": "62789211-2c4c-4eb6-8c3e-9223de49e529",
      "external_id": "",
      "merchant_id": "8d384f61-c92b-4867-abfc-f78ed3b6ea15",
      "customer_id": "41cd94bb-50d4-4ff3-923e-9b3c919895ca",
      "payment_method_id": "ec228ab0-7abc-4e15-a1c6-5df6ba5534af",
      "amount": 1,
      "status": "PENDING",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z",
      "processed_at": null
    },
    {
      "id": "a46b4f3d-ba1e-4e3d-9e8a-2b82abdcfa79",
      "external_id": "",
      "merchant_id": "8d384f61-c92b-4867-abfc-f78ed3b6ea15",
      "customer_id": "41cd94bb-50d4-4ff3-923e-9b3c919895ca",
      "payment_method_id": "ec228ab0-7abc-4e15-a1c6-5df6ba5534af",
      "amount": 1,
      "status": "PENDING",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z",
      "processed_at": null
    }
  ]
}
```

- `400 Bad Request`: error finding payments, please try again later.

### Create Payment

**Endpoint:** `POST /api/payment`

**Description:** Creates a new payment.

**Request Body:**

- `PaymentInput` (JSON): The input data for creating a payment.
- `Authorization` (header): The token of the user logged in with a Bearer format. The user must be a customer.

```json
{
  "merchant_id": "8d384f61-c92b-4867-abfc-f78ed3b6ea15",
  "customer_id": "41cd94bb-50d4-4ff3-923e-9b3c919895ca",
  "payment_method_id": "ec228ab0-7abc-4e15-a1c6-5df6ba5534af",
  "amount": 1
}
```

**Responses:**

- `201 Created`: Payment created successfully

```json
{
  "message": "Payment created successfully",
  "data": {
    "id": "62789211-2c4c-4eb6-8c3e-9223de49e529",
    "external_id": "",
    "merchant_id": "8d384f61-c92b-4867-abfc-f78ed3b6ea15",
    "customer_id": "41cd94bb-50d4-4ff3-923e-9b3c919895ca",
    "payment_method_id": "ec228ab0-7abc-4e15-a1c6-5df6ba5534af",
    "amount": 1,
    "status": "PENDING",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
    "processed_at": null
  }
}
```

- `400 Bad Request`: Invalid request body.
- `500 Internal Server Error`: error creating payment, please try again later.

### Get Payment

**Endpoint:** `GET /api/payments/{paymentID}`

**Description:** Retrieves a payment by its ID.

**Request Parameters:**

- `paymentID` (path parameter): The ID of the payment to retrieve.
- `Authorization` (header): The token of the user logged in with a Bearer format.

**Responses:**

- `200 OK`: Payment retrieved successfully.

```json
{
  "message": "Payment created successfully",
  "data": {
    "id": "62789211-2c4c-4eb6-8c3e-9223de49e529",
    "external_id": "",
    "merchant_id": "8d384f61-c92b-4867-abfc-f78ed3b6ea15",
    "customer_id": "41cd94bb-50d4-4ff3-923e-9b3c919895ca",
    "payment_method_id": "ec228ab0-7abc-4e15-a1c6-5df6ba5534af",
    "amount": 1,
    "status": "PENDING",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
    "processed_at": null
  }
}
```

- `404 Not Found`: payment not found.
- `500 Internal Server Error`: error finding payment, please try again later

### Process Payment

**Endpoint:** `POST /payments/{paymentID}/process`

**Description:** Processes an existing payment.

**Request Parameters:**

- `paymentID` (path parameter): The ID of the payment to process.
- `Authorization` (header): The token of the user logged in with a Bearer format. The user must be the merchant owner of the payment.

**Responses:**

- `200 OK`: Payment processed successfully.

```json
{
  "message": "Payment processed successfully",
  "data": {
    "id": "62789211-2c4c-4eb6-8c3e-9223de49e529",
    "external_id": "7b2b4b3d-ba1e-4e3d-9e8a-2b82abdcfa79",
    "merchant_id": "8d384f61-c92b-4867-abfc-f78ed3b6ea15",
    "customer_id": "41cd94bb-50d4-4ff3-923e-9b3c919895ca",
    "payment_method_id": "ec228ab0-7abc-4e15-a1c6-5df6ba5534af",
    "amount": 1,
    "status": "PROCESSED",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
    "processed_at": "2021-09-29T00:00:00Z"
  }
}
```

- `400 Bad Request`: payment has already been processed
- `404 Not Found`: Payment not found.
- `500 Internal Server Error`: error finding payment, please try again later
- `500 Internal Server Error`: error processing payment, please try again later

### Refund Payment

**Endpoint:** `POST /payments/{paymentID}/refund`

**Description:** Refunds an existing payment.

**Request Parameters:**

- `paymentID` (path parameter): The ID of the payment to refund.
- `Authorization` (header): The token of the user logged in with a Bearer format. The user must be the merchant owner of the payment.

**Responses:**

- `200 OK`: Payment refunded successfully.

```json
{
  "message": "Payment refunded successfully",
  "data": {
    "id": "62789211-2c4c-4eb6-8c3e-9223de49e529",
    "external_id": "7b2b4b3d-ba1e-4e3d-9e8a-2b82abdcfa79",
    "merchant_id": "8d384f61-c92b-4867-abfc-f78ed3b6ea15",
    "customer_id": "41cd94bb-50d4-4ff3-923e-9b3c919895ca",
    "payment_method_id": "ec228ab0-7abc-4e15-a1c6-5df6ba5534af",
    "amount": 1,
    "status": "REFUNDED",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
    "processed_at": "2021-09-29T00:00:00Z"
  }
}
```

- `400 Bad Request`: payment cannot be refunded
- `404 Not Found`: Payment not found.
- `500 Internal Server Error`: error finding payment, please try again later
- `500 Internal Server Error`: error refunding payment, please try again later
