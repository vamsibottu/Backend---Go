### Steps to run the worker

#### Requires:

* ![github.com/aws/aws-sdk-go/aws](https://github.com/aws/aws-sdk-go/aws)

* ![github.com/aws/aws-sdk-go/aws/credentials](https://github.com/aws/aws-sdk-go/aws/credentials)

* ![github.com/aws/aws-sdk-go/aws/session](https://github.com/aws/aws-sdk-go/aws/sessions)

* ![github.com/aws/aws-sdk-go/service/sqs](github.com/aws/aws-sdk-go/service/sqs)

* ![github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)

* ![github.com/lib/pq](https://"github.com/lib/pq)


#### Required Environmental Variables


```AWS
AWS_access_key = *******************
AWS_secret_key = *******************
AWS_region     = Us.East-1(virginia)
var token      = *******************
AWS_queue_url  = *******************

```
##### postgreSQL

```postgreSQL
PSQL_PORT = 5432
PSQL_URL  = AWS_ENDPOINT_DB
```
### How To Run
#### go run duntoday_worker.go
