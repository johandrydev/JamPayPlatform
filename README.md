# JAMPay Payment Platform

## Overview

JAMPay is a payment platform that allows users to process and refund payments. This project is built using Go,
PostgreSQL and Stripe to simulate an acquiring back.

## Prerequisites

- Go 1.20 or later
- PostgreSQL
- Git
- Stripe account

## Installation

1. **Clone the repository:**
   ```sh
   https://github.com/johandrydev/JamPayPlatform.git
   cd JamPayPlatform
   ```

2. **Set up the database:**

- Create a new database in PostgreSQL.

You could use docker to run the PostgreSQL database. To do so, run the following command:

```sh
docker run --name {container-name} -e POSTGRES_USER={postgres-user} -e POSTGRES_PASSWORD={postgres-password} -e POSTGRES_DB={database-name} -p 5432:5432 -d postgres
```

- Update the database connection string in the `.env` file. You can use the `.env.example` file as a template.

3. **Install dependencies:**
   ```sh
   go mod download
   ```

4. **Set up the Stripe account:**

- Create a new account in Stripe.
- Get the API keys from the Stripe dashboard.
- Update the Stripe API keys in the `.env` file. using the `.env.example` file as a template. in the key `STRIPE_SECRET_KEY`.

## Running the Server

1. **Build the project:**
   ```sh
   go build ./cmd/api/jam_pay.go
   ```

2. **Run the server:**

In development mode:

```sh
  go run ./cmd/api/jam_pay.go
```

To run the binary file:

   ```sh
   ./jam_pay
   ```

## API Endpoints

To see the list of available endpoints, check the [API service definition](./service_definition.md).

