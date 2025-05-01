package migrations

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"size:255;index:idx_email,unique"`
	Password  string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type Usergroup struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
type UsergroupUser struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID;references:ID"`
	UsergroupID uint      `gorm:"not null"`
	Usergroup   Usergroup `gorm:"foreignKey:UsergroupID;references:ID"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type Role struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type UsergroupRole struct {
	ID          uint      `gorm:"primaryKey"`
	UsergroupID uint      `gorm:"not null"`
	Usergroup   Usergroup `gorm:"foreignKey:UsergroupID;references:ID"`
	RoleID      uint      `gorm:"not null"`
	Role        Role      `gorm:"foreignKey:RoleID;references:ID"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type Permission struct {
	ID        uint      `gorm:"primaryKey"`
	ParentID  uint      `gorm:"default:0"`
	Name      string    `gorm:"size:100;not null"`
	Key       string    `gorm:"size:100;default:null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type RolePermission struct {
	ID           uint       `gorm:"primaryKey"`
	RoleID       uint       `gorm:"not null"`
	Role         Role       `gorm:"foreignKey:RoleID;references:ID"`
	PermissionID uint       `gorm:"not null"`
	Permission   Permission `gorm:"foreignKey:PermissionID;references:ID"`
	IsIndex      bool       `gorm:"default:false"`
	CreatedAt    time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

var Entities = []interface{}{
	&User{},
	&Usergroup{},
	&UsergroupUser{},
	&Role{},
	&UsergroupRole{},
	&Permission{},
	&RolePermission{},
}
