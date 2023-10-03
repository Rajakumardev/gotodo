package models

type Todo struct {
	ID          int64  `json:"id,omitempty" gorm:"primaryKey`
	Description string `json:"description,omitempty"`
	Complted    bool   `json:"complted,omitempty" gorm:"default:false"`
}
