package entities

import "time"

type Address struct {
	Address_uid string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User        User
	User_uid    string `gorm:"size:36;index"`
	Name        string `gorm:"size:100"`
	IsPrimary   bool
	CityID      string `gorm:"size:100"`
	ProvinceID  string `gorm:"size:100"`
	Address1    string `gorm:"size:255"`
	Address2    string `gorm:"size:255"`
	Phone       string `gorm:"size:100"`
	Email       string `gorm:"size:100"`
	PostCode    string `gorm:"size:100"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
