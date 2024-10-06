# JAMPay Payment Platform

## Overview

JAMPay is a payment platform that allows users to process and refund payments. This project is built using Go,
PostgreSQL and Stripe to simulate an acquiring back.

## Prerequisites

- Go 1.20 or later
- PostgreSQL
- [Tern (for running the migrations)](https://github.com/jackc/tern)
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

docker run --name {container-name} -e POSTGRES_USER={postgres-user} -e POSTGRES_PASSWORD={postgres-password} -e POSTGRES_DB={database-name} -p 5432:5432 -v {local-path}:/var/lib/postgresql/data -d postgres:latest

- Update the database connection string in the `.env` file. You can use the `.env.example` file as a template.
- Also, need to install tern to run the migrations. To do so, run the following command:

   ```sh
   go install github.com/jackc/tern/v2@latest
   ```

- Set up the database information in the tern.conf file. You can use the tern.conf.example file as a template.
- Run the migrations:

   ```sh
   cd migrations
   tern migrate
   ```

3. **Install dependencies:**
   ```sh
   go mod download
   ```

4. **Set up the Stripe account:**

- Create a new account in Stripe.
- Get the API keys from the Stripe dashboard.
- Update the Stripe API keys in the `.env` file. using the `.env.example` file as a template, in the key `STRIPE_SECRET_KEY`.
- To do a test of payments, in this first phase you need to create a customer with a card to simulate the payment. You can use the Stripe dashboard to create a customer and get the customer ID and the payment method ID. Take in mind the payment method should belong to the customer.
- add the customer ID in the table `customers` and the card ID in the table `payment_methods` in the database. You can use this command to insert the data in query tool in the database:

   ```sh
    UPDATE customers SET external_id = {customer_stripe_id} WHERE status = 'ACTIVE';
    UPDATE payment_methods SET external_id = {payment_method_stripe_id} WHERE product_number = '4242424242424242';
     ```

## Running the Server

1. **Build the project:**
   ```sh
   go build ./cmd/api/jam_pay.go
   ```

2. **Run the server:**

- In development mode:

   ```sh
     go run ./cmd/api/jam_pay.go
   ```

- To run the binary file:

   ```sh
   ./jam_pay
   ```

## API Endpoints

To see the list of available endpoints, check the [API service definition](./service_definition.md).

