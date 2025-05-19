package services

import (
	"edugo/models"
	"fmt"
)

var currentUser *models.User

func RegisterUser(username, password, email string) (*models.User, error) {
	user := &models.User{
		ID:       len(users) + 1,
		Username: username,
		Password: password,
		Email:    email,
		Courses:  []int{},
	}
	data.Users = append(dataUsers, user)
	return &user, nil
}

func LoginUser(username, password string) (*models.User, error) {
	for _, user := range data.Users {
		if user.Username == username && user.Password == password {
			currentUser = &user
			return &user, nil
		}
	}
	return nil, fmt.Errorf("invalid username or password")
}
func BuyCourse(user *models.User, courseID int) {
	user.Courses = append(user.Courses, courseID)
}
