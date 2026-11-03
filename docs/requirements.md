Order Statuses

CREATED
PAYMENT_PENDING
PAID
PAYMENT_FAILED
CANCELLED

Payment Statuses

PENDING
PROCESSING
SUCCESS
FAILED

Functional Requirements

1. A user should be able to create an order.

2. A user should be able to retrieve an order using its ID.

3. Every order should have a status.

4. A payment should be created for an order.

5. Payment processing can succeed or fail.

6. Order status should change based on payment result.

7. Duplicate payment requests should not create duplicate payments.

8. A notification should be generated after payment succeeds or fails.

9. Each service should run independently.

10. Each service should own its own data.


Non-Functional Requirements

1. The system should handle failures gracefully.

2. Services should be independently deployable.

3. Payment processing should be idempotent.

4. The system should support logging and monitoring.

5. APIs should return meaningful error responses.

6. Services should support graceful shutdown.

7. The system should eventually support distributed tracing.