package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sadysnaat/assignment2/models/property"
)

var (
	getProperties      = "SELECT * FROM properties"
	insertProperties   = `INSERT INTO properties(name, cost, color) VALUES ('%s', %d, '%s')`
	findByNameProperty = `SELECT * FROM properties where name = '%s'`
)

type Manager struct {
	db *sql.DB
}

func NewManager(uri string) (*Manager, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return &Manager{}, err
	}
	m := &Manager{db: db}
	return m, nil
}

func (m *Manager) Save(p *property.Property) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	fmt.Println(p.Name, p.Color, p.Cost)
	_, err = tx.Exec(fmt.Sprintf(insertProperties, p.Name, p.Cost, p.Color))
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) GetProperties() ([]*property.Property, error) {
	var res []*property.Property
	rows, err := m.db.Query(getProperties)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		p := &property.Property{}
		err := rows.Scan(&p.Name, &p.Cost, &p.Color)
		if err != nil {
			fmt.Println(err)
		}
		res = append(res, p)
	}

	return res, nil
}

func (m *Manager) GetPropertyByName(name string) (*property.Property, error) {
	rows, err := m.db.Query(fmt.Sprintf(findByNameProperty, name))
	if err != nil {
		return &property.Property{}, err
	}
	defer rows.Close()

	for rows.Next() {
		p := &property.Property{}
		err := rows.Scan(&p.Name, &p.Cost, &p.Color)
		if err != nil {
			fmt.Println(err)
		}
		return p, nil
	}
	return &property.Property{}, nil
}
