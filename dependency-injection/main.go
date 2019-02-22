package main

import (
	"go.uber.org/dig"
)

//Dependency Injection is the idea that your components (usually structs in go)
//should receive their dependencies when being created

//“container” is often used in DI frameworks to describe the thing into which you add “providers”
//and out of which you ask for fully-build objects.

//dig library gives us the Provide function for adding providers and
//the Invoke function for retrieving fully-built objects out of the container

//The container recognizes that we’re asking for a *sql.DB
//It determines that our function ConnectDatabase provides that type
//It next determines that our ConnectDatabase function has a dependency of Config
//It finds the provider for Config, the NewConfig function NewConfig doesn’t have any dependencies, so it is called
//The result of NewConfig is a Config that is passed to ConnectDatabase
//The result of ConnectionDatabase is a *sql.DB that is passed back to the caller of Invoke

func BuildContainer() *dig.Container {
	//build a new container
	container := dig.New()

	//add new providers
	//I provide a Config type to the container. In order to build it, I don’t need anything else.
	container.Provide(NewConfig)
	container.Provide(ConnectDatabase)
	container.Provide(NewPersonRepository)
	container.Provide(NewPersonService)
	container.Provide(NewServer)

	return container
}

func main() {
	container := BuildContainer()

	//Now, we can ask the container to give us a fully-built component
	//for any of the types we’ve provided. We do so using the Invoke function
	err := container.Invoke(func(server *Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}
