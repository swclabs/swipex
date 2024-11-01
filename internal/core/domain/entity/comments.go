package entity

// Comment Table
type Comment struct {
	ID        int64  `json:"id" db:"id"`
	Level     int64  `json:"level" db:"level"`
	Content   string `json:"content" db:"content"`
	UserID    int64  `json:"user_id" db:"user_id"`
	ProductID int64  `json:"product_id" db:"product_id"`
	ParentID  int64  `json:"parent_id" db:"parent_id"`
	Rating    int64  `json:"rating" db:"rating"`
	Liked     int64  `json:"liked" db:"liked"`
	Disliked  int64  `json:"disliked" db:"disliked"`
	// Created   time.Time `json:"created" db:"created"`
}
