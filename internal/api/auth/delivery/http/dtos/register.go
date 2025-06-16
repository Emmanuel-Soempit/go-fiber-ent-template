package dtos

type RegisterUserPayload struct {
	Firstname string `json:"firstname" xml:"firstname" form:"firstname"`
	Lastname  string `json:"lastname" xml:"lastname" form:"lastname"`
	Email     string `json:"Email" xml:"Email" form:"Email"`
	Password  string `json:"Password" xml:"Password" form:"Password"`
}
