package models

type Roles struct {
	ID       uint   `json:"role_id" gorm:"primaryKey"`
	RoleName string `json:"role_name" gorm:"not null"`
}

func (r *Roles) TableName() string {
	return "roles"
}
