package models

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Courses []int  `json:"courses"`
}

func (u *User) Register(name string) {

}
func (u *User) Login() bool {
	return true

}
func (u *User) BuyCourse(courseID int) {
	u.Courses = append(u.Courses, courseID)
}
