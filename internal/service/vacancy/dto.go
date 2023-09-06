package vacancy

type Create struct {
	FullName    string `form:"full_name"`
	PhoneNumber string `form:"phone_number"`
	Email       string `form:"email"`
	CV          string
}

type SetRating struct {
	Rating int `form:"rating"`
}
