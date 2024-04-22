package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

// instance create a single singleton instance
var instance *singleton

// GetInstance will always return the same instance
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

// Solo tenga en cuenta que el patrón Singleton
// le dará el poder de tener una instancia única de alguna estructura en su aplicación
// y que ningún paquete puede crear ningún clon de esta estructura.
