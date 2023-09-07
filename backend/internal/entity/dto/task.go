package dto

type Task struct {
	Title      string  `json:"title"       csv:"title"       db:"title"`
	ParentTask *string `json:"parent_task" csv:"parent_task" db:"parent_task"`
	MaxXp      uint    `json:"max_xp"      csv:"max_xp"      db:"max_xp"`
}
