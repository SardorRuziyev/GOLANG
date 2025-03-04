/*ICourse interface – Barcha kurslar uchun umumiy metodlar
• PaidCourse, FreeCourse struct – Pullik va bepul kurslar
• GetCourseInfo() – Kurs haqida ma’lumot olish
• Enroll() – Kursga yozilish
*/

package models

import "fmt"

type ICourse interface {
	GetCourseInfo() string
	Enroll()
}
type Course struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsFree      bool    `json:"is_free"`
}
type PaidCourse struct {
	Course
}

type FreeCourse struct {
	Course
}

func (c *PaidCourse) GetCourseInfo() string {
	return c.Title + " - $" + fmt.Sprintf("%.2f", c.Price)
}

func (c *PaidCourse) Enroll(userID int) error {
	// Simulate payment logic
	return nil
}

func (c *FreeCourse) GetCourseInfo() string {
	return c.Title + " - Free"
}

func (c *FreeCourse) Enroll(userID int) error {
	// No payment required
	return nil
}
