package models

import "fmt"

type ICourse interface {
	GetCourseID() string

	EnrollStudent(userID int) error
	GetID() int
	GetName() string
}

type PaidCourse struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Enrolled []int   `json:"enrolled"`
}
type FreeCourse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Enrolled []int  `json:"enrolled"`
}

func (p PaidCourse) GetCourseID() string {
	return fmt.Sprintf("%d", p.ID)
}
func (p PaidCourse) EnrollStudent(userID int) error {
	p.Enrolled = append(p.Enrolled, userID)
	return nil
}
func (p PaidCourse) GetID() int {
	return p.ID
}

func (f FreeCourse) GetCourseID() string {
	return f.Title
}
func (f FreeCourse) EnrollStudent(userID int) error {
	f.Enrolled = append(f.Enrolled, userID)
	return nil
}
func (f FreeCourse) GetID() int {
	return f.ID
}
