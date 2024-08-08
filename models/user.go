package models

type User struct {
	Id       int    `json:"-"`
	HiddenId int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	PlanId   int    `json:"-"`
}

func (u *User) HideId() {
	u.HiddenId = u.Id + 12345
}

func (u *User) UnhideId() {
	u.Id = u.HiddenId - 12345
}
