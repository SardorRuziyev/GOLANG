package models

type Review struct {
	CourseID int    `json:"course_id"`
	UserID   int    `json:"user_id"`
	Rating   int    `json:"rating"`
	Comment  string `json:"comment"`
}
