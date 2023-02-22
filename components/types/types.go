package types

import ( 
	"time"
	"github.com/lib/pq"
)


type User struct {
	ID        string              `json:"id" gorm:"primaryKey"`
	Username  string              `json:"username" gorm:"uniqueIndex"`
	Password  string              `json:"password"`
	FirstName string              `json:"first" gorm:"index"`
	LastName  string              `json:"last" gorm:"index"`
	Grade     string              `json:"grade" gorm:"index"`
	Code      string              `json:"code" gorm:"uniqueIndex"`
	Machines  []*Machine `gorm:"many2many:users_machines"`
	CreatedAt time.Time           `json:"-" gorm:"index"`
	UpdatedAt time.Time           `json:"-" gorm:"index"`
}

type Admin struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex"`
	Password  string    `json:"password"`
	FirstName string    `json:"first" gorm:"index"`
	LastName  string    `json:"last" gorm:"index"`
	Code      string    `json:"code" gorm:"uniqueIndex"`
	CreatedAt time.Time `json:"-" gorm:"index"`
	UpdatedAt time.Time `json:"-" gorm:"index"`
}

type Machine struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	InUSE     bool      `json:"in_use" gorm:"index"`
	Users  []*User `gorm:"many2many:users_machines"`
	Actions   pq.StringArray  `json:"actions" gorm:"type:text[]"`
	CreatedAt time.Time `json:"-" gorm:"index"`
	UpdatedAt time.Time `json:"-" gorm:"index"`
}