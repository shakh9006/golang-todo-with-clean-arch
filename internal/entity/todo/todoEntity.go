package entity

type Todo struct {
	ID        int    `json:"id,omitempty" db:"id"`
	Title     string `json:"title" db:"title"`
	Completed bool   `json:"completed" db:"completed"`
}
