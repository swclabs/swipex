package entity

// Supplier table
type Supplier struct {
	ID    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}
