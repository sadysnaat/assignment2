package property

// Property represents monopoly property
type Property struct {
	Name  string
	Cost  int
	Color string
}

// Data Mapper pattern where in memory struct is separate from persistence layer
// https://martinfowler.com/eaaCatalog/dataMapper.html
type Mapper interface {
	Save(p *Property) error
	GetProperties() ([]*Property, error)
	GetPropertyByName(name string) (*Property, error)
}
