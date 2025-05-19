package services

import "edugo/storage"

func AddReview(data *storage.Data, courseID, userID, rating int, text string) error {
	review := models.Review{
		CourseID: courseID,
		UserID:   userID,
		Rating:   rating,
		Comment:  text,
	}

	data.Reviews = append(data.Reviews, review)

	func GetReviewsForCourse(data *storage.Data, courseID int) []models.Review {
		var result []models.Review
		for _, review := range data.Reviews {
			if review.CourseID == courseID {
				result = append(result, review)
			}
		}
		return result
	