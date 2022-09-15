package models

import "time"

type Book struct {
	Id            int
	Title         string
	Author        string
	PublisherYear time.Time
	//Id            int       `json:"id"`
	//Title         string    `json:"title"`
	//Author        string    `json:"author"`
	//PublisherYear time.Time `json:"publisher_year"`
}
