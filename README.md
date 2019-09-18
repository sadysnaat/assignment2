#### Installing
run 
```
go get github.com/sadysnaat/assignment2
```

#### Running 
Build the project and run 
```
// If you want to use mysql as store
./assignment2 --mysql 

// If you want to use postgres as store
./assignment2 --postgres
```

Please change the uri in the manager initialization in main.go
currently the values are 
* MySQL "root:my-secret-pw@tcp(127.0.0.1:32768)/monopoly"
* PostgreSQL "postgres://postgres:docker@127.0.0.1:5432/monpoly?sslmode=disable"

Please make sure you have relevant table 'properties' created in the database
```
create table properties (
    name varchar(100),
    cost int,
    color varchar(6)
);
``` 

#### Design
This project uses datamapper pattern 
Models have an Mapper interface which contain all the utility functions required
A Store Mananger implements these functions 