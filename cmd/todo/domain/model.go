package domain

type Todo struct {
	Id        int64  `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}
