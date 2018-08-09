package resource

import (
	"github.com/aws/aws-lambda-go/events"
)

type Headers struct {
	AccessControlAllowMethods			string `json:"Access-Control-Allow-Methods"`
	AccessControlAllowOrigin			string `json:"Access-Control-Allow-Origin"`
	AccessControlAllowHeaders			string `json:"Access-Control-Allow-Headers"`
	ContentType										string `json:"Content-Type"`
	AccessControlAllowCredentials string `json:"Access-Control-Allow-Credentials"`
}

type GatewayResponse struct {
	IsBase64Encoded bool		`json:"isBase64Encoded"`
	StatusCode			int			`json:"statusCode"`
	Headers					Headers `json:"headers"`
	Body						string	`json:"body"`
}

type Handler func(events.APIGatewayProxyRequest) (GatewayResponse, error)

type Resource struct {
	post, unauthorized, options Handler
}

func New() Resource {
	return Resource{}
}

func (r *Resource) Options(handler Handler) *Resource {
	r.options = handler

	return r
}

func (r *Resource) Create(handler Handler) *Resource {
	r.post = handler

	return r
}

func (r *Resource) Unauthorized(handler Handler) *Resource {
	r.unauthorized = handler

	return r
}

func (r Resource) Run(event events.APIGatewayProxyRequest) (GatewayResponse, error) {
	if event.HTTPMethod == "POST" && r.post != nil {
		return r.post(event)
	} else if event.HTTPMethod == "OPTIONS" && r.options != nil {
		return r.options(event)
	}

	return r.unauthorized(event)
}
