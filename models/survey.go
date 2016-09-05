package models

type Survey struct {
	BaseModel
	TransactionWith string `json:"transaction_with" form:"transaction_with" binding:"required"`
	Division string `json:"division" form:"division" binding:"required"`
	ServiceRequested string `json:"service_requested" form:"service_requested"`
	Comments string `json:"comments" form:"comments"`
	Name string `json:"name" form:"name"`
	Agency string `json:"agency" form:"agency"`
	Address string `json:"address" form:"address"`
	ContactNo string `json:"contact_no" form:"contact_no"`
	Email string `json:"emai" form:"email"`
	Rating int `json:"rating" form:"rating"`
	Status string `json:"status"`
}

func (i *Survey) BeforeCreate() (err error) {
	i.Status = "active"
	return
}