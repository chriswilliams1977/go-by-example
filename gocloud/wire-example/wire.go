//+build wireinject

package main

import (
	"github.com/google/wire"
)

//you can group providers in a set
var Superset = wire.NewSet(NewEvent, NewGreeter, NewMessage)

//InitializeEvent :
//single call to wire.Build passing in the initializers we want to use
//initializers = providers in wire
//InitializeEvent = injector
//zero value for Event as a return value to satisfy the compiler. Does not accept values
func InitializeEvent(phrase string) (Event, error) {
	wire.Build(Superset)
	return Event{}, nil
}
