package models

import "github.com/lib/pq"

type User struct {
	Id       string `json:"id" gorm:"type:uuid;primaryKey:default:gen_random_uuid()"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null; unique"`
	Password string `json:"password" gorm:"not null"`
	State    string `json:"state" gorm:"not null"`
	Cpf      string `json:"cpf" gorm:"not null; unique"`
}

type Proposal struct {
	Id        string  `json:"id" gorm:"type:uuid;primaryKey:default:gen_random_uuid()"`
	Value     float32 `json:"value" gorm:"not null"`
	User      User    `gorm:"foreignKey:UserId; not null"`
	UserId    string  `json:"user" gorm:"not null"`
	AuctionId string  `json:"auctionId" gorm:"not null"`
	Auction   Auction `gorm:"foreignKey:AuctionId"`
}

type Auction struct {
	Id          string         `json:"id" gorm:"type:uuid;primaryKey:default:gen_random_uuid()"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	Category    Category       `gorm:"foreignKey:CategoryId; not null"`
	CategoryId  string         `json:"category" gorm:"not null"`
	Owner       User           `gorm:"foreignKey:OwnerId; not null"`
	OwnerId     string         `json:"owner" gorm:"not null"`
	Files       pq.StringArray `gorm:"type:text[]"`
	Proposals   []Proposal     `json:"proposals"`
}

type Category struct {
	Id          string `json:"id" gorm:"type:uuid;primaryKey:default:gen_random_uuid()"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}
