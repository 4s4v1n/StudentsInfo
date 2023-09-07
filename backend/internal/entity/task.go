package entity

type Task struct {
	Title      string  `json:"title"       csv:"title"`
	ParentTask *string `json:"parent_task" csv:"parent_task"`
	MaxXp      uint    `json:"max_xp"      csv:"max_xp"`
}
