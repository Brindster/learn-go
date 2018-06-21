package service

import (
	"errors"
)

var (
	// ErrNotFound is returned when a factory for the service couldn't be found
	ErrNotFound = errors.New("Service could not be found")

	// ErrInvalidType is returned when a factory returns an unexpected type
	ErrInvalidType = errors.New("Service was incorrect type")
)

// Factory is a function that returns a service
type Factory func(Container) (interface{}, error)

// Container is an interface for a service that contains factories to create services
type Container interface {
	Build(string) (interface{}, error)
	Get(string) (interface{}, error)
	MustBuild(string) interface{}
	MustGet(string) interface{}
}

// Services is a service that contains factories to create services
type Services struct {
	factories map[string]Factory
	services  map[string]interface{}
}

// NewServices creates a new service container with the given factories
func NewServices(fcts map[string]Factory) *Services {
	var s map[string]interface{}
	s = make(map[string]interface{})
	return &Services{
		factories: fcts,
		services:  s,
	}
}

func (s *Services) build(a string) (interface{}, error) {
	factory, ok := s.factories[a]
	if !ok {
		return nil, ErrNotFound
	}

	service, err := factory(s)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (s *Services) get(a string) (interface{}, error) {
	service, ok := s.services[a]
	if ok {
		return service, nil
	}

	built, err := s.build(a)
	if err != nil {
		return nil, err
	}

	s.services[a] = built
	return built, nil
}

// Build will create a new instance of the service
func (s *Services) Build(a string) (interface{}, error) {
	return s.build(a)
}

// Get will retrieve an existing service if if exists, otherwise build a new one
func (s *Services) Get(a string) (interface{}, error) {
	return s.get(a)
}

// MustBuild will return the service, or panic if it doesnt exist
func (s *Services) MustBuild(a string) interface{} {
	srv, err := s.Build(a)
	if err != nil {
		panic(err)
	}
	return srv
}

// MustGet will return the service, or panic if it doesnt exist
func (s *Services) MustGet(a string) interface{} {
	srv, err := s.Get(a)
	if err != nil {
		panic(err)
	}
	return srv
}
