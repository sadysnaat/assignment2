package property

// Property represents monopoly property
type Property struct {
	Name  string
	Cost  int
	Color string
}

type Mapper interface {
	Save(p *Property) error
	GetProperties() ([]*Property, error)
	GetPropertyByName(name string) (*Property, error)
}
