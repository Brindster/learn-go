package service

import (
	"testing"
)

func testfac(c Container) (interface{}, error) {
	return interface{}("Result"), nil
}

// TestNewServices tests the creation of a container
func TestNewServices(t *testing.T) {
	var fct map[string]Factory
	services := NewServices(fct)
	if _, ok := interface{}(services).(Container); !ok {
		t.Error("Expected instance of Container")
	}
}

// TestBuildReturnsCorrectError tests the building of a service will return no error
func TestBuildReturnsCorrectError(t *testing.T) {
	var fct map[string]Factory

	services := NewServices(fct)
	if _, err := services.Build("unknown"); err.Error() != "Service could not be found" {
		t.Error("Expected error of type ErrNotFound")
	}
}

// TestBuildReturnsNoError tests the building of a service will return no error
func TestBuildReturnsNoError(t *testing.T) {
	var fct map[string]Factory
	fct = make(map[string]Factory)
	fct["test"] = testfac

	services := NewServices(fct)
	if _, err := services.Build("test"); err != nil {
		t.Error("Expected no error")
	}
}

// TestBuildReturnsService tests the building of a service
func TestBuildReturnsService(t *testing.T) {
	var fct map[string]Factory
	fct = make(map[string]Factory)
	fct["test"] = testfac

	services := NewServices(fct)
	service, _ := services.Build("test")

	if service != "Result" {
		t.Error("Expecting the result to be `Result`")
	}
}

// TestBuildReturnsNewInstance tests the building of a service
func TestBuildReturnsNewInstance(t *testing.T) {
	var fct map[string]Factory
	fct = make(map[string]Factory)
	fct["test"] = testfac

	services := NewServices(fct)
	a, _ := services.Build("test")
	b, _ := services.Build("test")

	if &a == &b {
		t.Error("Expecting the result to be different")
	}
}

// TestGetReturnsCorrectError tests the building of a service will return no error
func TestGetReturnsCorrectError(t *testing.T) {
	var fct map[string]Factory

	services := NewServices(fct)
	if _, err := services.Build("unknown"); err.Error() != "Service could not be found" {
		t.Error("Expected error of type ErrNotFound")
	}
}

// TestGetReturnsNoError tests the building of a service will return no error
func TestGetReturnsNoError(t *testing.T) {
	var fct map[string]Factory
	fct = make(map[string]Factory)
	fct["test"] = testfac

	services := NewServices(fct)
	if _, err := services.Get("test"); err != nil {
		t.Error("Expected no error")
	}
}

// TestGetReturnsService tests the building of a service
func TestGetReturnsService(t *testing.T) {
	var fct map[string]Factory
	fct = make(map[string]Factory)
	fct["test"] = testfac

	services := NewServices(fct)
	service, _ := services.Get("test")

	if service != "Result" {
		t.Error("Expecting the result to be `Result`")
	}
}

// TestGetReturnsSameInstance tests the building of a service
func TestGetReturnsSameInstance(t *testing.T) {
	var fct map[string]Factory
	fct = make(map[string]Factory)
	fct["test"] = testfac

	services := NewServices(fct)
	a, _ := services.Get("test")
	b, _ := services.Get("test")

	if a != b {
		t.Error("Expecting the result to be the same")
	}
}
