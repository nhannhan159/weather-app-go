package domain

type IService interface {
	FindAll(resources *Resources) (interface{}, error)
	FindByID(resources *Resources, id int) (interface{}, error)
	Create(resources *Resources, entity interface{}) (interface{}, error)
	Update(resources *Resources, entity interface{}) (interface{}, error)
	Delete(resources *Resources, id int) (interface{}, error)
}
