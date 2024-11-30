package database

import "restAPI/app/queries"

// Queries struct for collect all app queries.
type Queries struct {
	*queries.BookQueries
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) { 
	db, err := PostgreSQLConnection()
	if err != nil { 
		return nil, err
	}

	return &Queries{
		BookQueries: &queries.BookQueries{DB: db},
	}, nil	
}