# AWS Lambda Resource

Helps form restful resource endpoints for AWS API Gateway + Lambda.


```golang
package main

import (
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/espinola-designs/aws-lambda-resource"
)

func main() {
  users := resource.New() 

  users.Create(func(_ events.APIGatewayProxyRequest) (resource.GatewayRespoonse, error) {})

  users.Get(func(_ events.APIGatewayProxyRequest) (resource.GatewayRespoonse, error) {})

  lambda.Start(users)
}
```
