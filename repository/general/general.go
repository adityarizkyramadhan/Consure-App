package general

type GeneralRepository interface {
	Create(interface{}) error
	FindById(int, interface{}) error
	FindAll(interface{}) error
	Delete(int, interface{}) error
	Update(int, interface{}) error
}
