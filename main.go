package main

import (
	"flag"
	"fmt"
	"github.com/sadysnaat/assignment2/models/property"
	"github.com/sadysnaat/assignment2/store/mysql"
	"github.com/sadysnaat/assignment2/store/postgres"
	"os"
)

var (
	mysqlflag    = flag.Bool("mysql", false, "if you want to use mysql")
	postgresflag = flag.Bool("postgres", false, "if you want to use postgres")
)

func main() {
	flag.Parse()

	// as both store mysql and postgres manager fulfil the property mapper
	// this can be initialised according to the datasource switch and passed around
	// and use to manage properties via the interface methods
	var pm property.Mapper

	if *mysqlflag == *postgresflag {
		if *mysqlflag {
			fmt.Println("you can use only one data store at a time")
		} else {
			fmt.Println("at least one database should be selected")
		}
		os.Exit(2)
	}

	if *mysqlflag {
		m, err := mysql.NewManager("root:my-secret-pw@tcp(127.0.0.1:32768)/monopoly")
		if err != nil {
			fmt.Println(err)
		}

		pm = m
	}

	if *postgresflag {
		m, err := postgres.NewManager("postgres://postgres:docker@127.0.0.1:5432/monpoly?sslmode=disable")
		if err != nil {
			fmt.Println(err)
		}

		pm = m
	}

	p := &property.Property{
		Name:  "MayFair",
		Cost:  300,
		Color: "Blue",
	}

	err := pm.Save(p)
	if err != nil {
		fmt.Println(err)
	}

	ps, err := pm.GetProperties()
	if err != nil {
		fmt.Println(err)
	}

	for key, prop := range ps {
		fmt.Printf("%d name=%s cost=%d color=%s\n", key, prop.Color, prop.Cost, prop.Color)
	}
}
