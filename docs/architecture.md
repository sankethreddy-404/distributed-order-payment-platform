# Architecture

The system contains three main services:

1. Order Service
2. Payment Service
3. Notification Service

Initial flow:

Client
  |
  v
Order Service
  |
  v
Payment Service
  |
  v
Notification Service


Order Service:
- Creates orders
- Retrieves orders
- Maintains order status

Payment Service:
- Processes payments
- Stores payment status
- Prevents duplicate payments

Notification Service:
- Receives payment result
- Generates success or failure notifications

Future Architecture

Client
  |
  v
Order Service
  |
  +------ PostgreSQL
  |
  +------ Redis
  |
  v
Kafka
  |
  +------ Payment Service
  |          |
  |          +------ PostgreSQL
  |
  +------ Notification Service

All services
   |
   v
OpenTelemetry
   |
   v
SigNoz

