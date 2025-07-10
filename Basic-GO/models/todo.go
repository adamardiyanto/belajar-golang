package models

type Todo struct {
	Id    int    `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Done  bool   `json:"done" gorm:"default:false"`
}
