package domain

type IService interface {
	FindAll(resources *Resources) (interface{}, error)
	FindByID(resources *Resources, id int) (interface{}, error)
}
