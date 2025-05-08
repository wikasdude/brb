package domain

type Service struct {
	ID          int64
	Name        string
	Description string
	Active      bool
}

// CREATE TABLE services (
//     id SERIAL PRIMARY KEY,
//     name TEXT NOT NULL,
//     description TEXT,
//     active BOOLEAN DEFAULT TRUE
// );
