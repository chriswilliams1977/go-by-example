package main

//relies on both the Config and the PersonRepository
type PersonService struct {
	config     *Config
	repository *PersonRepository
  }
  
  //calls the PersonRepository if the application is enabled
  func (service *PersonService) FindAll() []*Person {
	if service.config.Enabled {
	  return service.repository.FindAll()
	}
  
	return []*Person{}
  }
  
  func NewPersonService(config *Config, repository *PersonRepository)
*PersonService {
	return &PersonService{config: config, repository: repository}
  }