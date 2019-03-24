# microservices-redis-stream

This project makes use of redis-streams [https://redis.io/topics/streams-intro] to communicaate business events among micro-services.

The project has the following microservices written in GO:

1. Accounts (HTTP enpointds accessible on port 8090)
This services provides the following endpoints:
- Get all accounts, with the following detail for each account:
  - Account ID
  - Owned By
  - Total revenue for the account
  - List of customers belonging to the account
  
Accounts service has two pre-loaded accounts:
_Account Name: one, Owned by: user1
Account Name: two, Onwed by: user2_

Account service also host authentication endpoint:
/auth/login
Ideally auth is best hosted as different service.

2. Customers service (HTTP enpointds accessible on port 8080)
- Create a new customer
- Create a new invoice

### Design

The core idea behind the design is that _Services publish business events that occur, other services listen to these events from the event-stream and update their data store or react in a way which is fit for them_


```
----------+                                                                                       +---------
 Customer |-----------------> Publish to redis-stream ~~~~~~~ <------ Read from redis stream ---- |  Accounts service
 service  |  (CustomerCreated, InvoiceCreated,                                                    | 
          |   InvoiceDeleted, CustomerDeleted...)                                                 | 
          |                                                                                       |
          |-----------------> Read fromredis-stream ~~~~~~~ <------ Publish to redis stream ------| 
----------+                                                     (AccountCreated, AccountDeleted)  +---------

```
Thus when an a customer is created, the accounts service which is responsible for providing the details about the aggregate revenue for an account does the following:

> It creates a new record for the customer in its own data store (saving only the relevant details)

Later when an Invoice is created for a customer, the accounts service gets to know about it from the event stream, and then it:
> It adds the new invoice amount as delta to the existing sum of all invoices for that customer

Thus each services reacts to business events in a way its sees fit.

### Keeping track of events consumed
The services start consuming from the event stream at start-up. It is imperative for the service to _start consuming the events from the right event_. The right event is from where the server had last consumed successfullly. The service has to maintain an event log in persistance store. This event log has to save the id of the event last consumed.

The service can then pass the last event id to `subscriber.Listen` method, which then returns the list of events from the event onwards.

### Creating a new customer:
```
POST: http://localhost:8080/customers
Body:
{
	"name": "John Adams", 
	"account_id": "1"
}
```
(Currently the account is not being validated for existence)

When a customer is created, its details are published as `CustomerCreated` event to `event-stream` redis-stream:
```
127.0.0.1:6379> XREAD STREAMS event-stream 0
1) 1) "event-stream"
   2) 1) 1) "1553318690326-0"
         2) 1) "Payload"
            2) "{\"id\":\"7\",\"account_id\":\"1\",\"name\":\"Martin North 4\",\"created_date\":\"0001-01-01T00:00:00Z\"}"
            3) "EventTypeName"
            4) "CustomerCreated"
```

#### Creating a new invoice
```
POST: http://localhost:8080/invoices
{
	"customer_id": "7", 
	"purchase_date": "2019-01-06",
	"purchase_price_cents": 10000
}
```

When an invoice is created, its details are published as `InvoiceCreated` event to redis stream (`event-stream`):
```
9) 1) "1553343729026-0"
         2) 1) "EventTypeName"
            2) "InvoiceCreated"
            3) "Payload"
            4) "{\"id\":\"10\",\"customer_id\":\"9\",\"purchase_price_cents\":8000,\"created_date\":\"0001-01-01T00:00:00Z\",\"purchase_date\":\"2019-01-09\"}"
```


