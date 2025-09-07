package dtos

type UpsertUserProfileRequestDTO struct {
	Fullname    string `json:"fullname"`
	Bio         string `json:"bio"`
	DateOfBirth int64  `json:"date_of_birth"`
	Gender      bool   `json:"gender"`
	Sentence    string `json:"sentence"`
	Author      string `json:"author"`
}
