package services

import (
	"edugo/models"
	"encoding/json"
	"fmt"
)

func GetAllCourses(data *storage.Data) {
	for _, course := range data.Courses {
		b, _ := json.Marshal(course)
		if _, ok := course["price"]; ok {

			var c models.PaidCourse
			json.Unmarshal(b, &c)
			fmt.Printf("[%d] %s\n", c.ID, c.GetCourseInfo())
		} else {
			var c models.FreeCourse
			json.Unmarshal(b, &c)
			fmt.Printf("[%d] %s\n", c.ID, c.GetCourseInfo())
		}
	}
}

func EnrollCourse(user *models.User, courseID int) error {
	for _, course := range data.Courses {
		if course.GetID() == courseID {
			if err := course.EnrollStudent(user.ID); err != nil {
				return err
			}
			user.Courses = append(user.Courses, courseID)
			return nil
		}
	}
	return fmt.Errorf("course not found")
}
