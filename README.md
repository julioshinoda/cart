# Cart API

This API provide services to manager cart services.

## Documentation

The services are documented on ```doc/cart-doc.yaml``` using **OpenAPI Specification** version 3.0.0

## Overview

The cart api was written in Go and use Redis as database. 

The project architecture is following [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)



## How to run

Use command **make run** to run application, redis and swagger-ui.

- the api run on port 8084
- swagger-ui run on port 8091

## Running test

Run **make test** to run all test on application

## To improve

- Use UUID for cart and item ID
- Add metrics
- Add validation for inputs
- Implement test for usecase
- Improve coupon use.
