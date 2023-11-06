package app

import (
	"server/internal/service"
)

type ServiceServer struct {
	exampleService service.ExampleService
	authService    service.AuthenticationService
	objectService  service.ObjectService
}

func NewService(
	exampleService service.ExampleService,
	authService service.AuthenticationService,
	objectService service.ObjectService,
) *ServiceServer {
	return &ServiceServer{
		exampleService: exampleService,
		authService:    authService,
		objectService:  objectService,
	}
}
