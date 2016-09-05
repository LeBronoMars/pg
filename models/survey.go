package models

type Survey struct {
	BaseModel
	LocalId int `json:"local_id" form:"local_id"`
	Division string `json:"division" form:"division" binding:"required"`
	ServiceRequested string `json:"service_requested" form:"service_requested"`
	Comments string `json:"comments" form:"comments"`
	Name string `json:"name" form:"name"`
	Agency string `json:"agency" form:"agency"`
	Address string `json:"address" form:"address"`
	ContactNo string `json:"contact_no" form:"contact_no"`
	Email string `json:"email" form:"email"`
	Rating int `json:"rating" form:"rating"`
	Status string `json:"-"`
}

func (i *Survey) BeforeCreate() (err error) {
	i.Status = "active"
	return
}