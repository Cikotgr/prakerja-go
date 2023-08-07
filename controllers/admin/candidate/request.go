package candidate

type CreateCandidate struct {
	ID       string `json:"id" form:"id" validate:"required"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
	Batch    int    `json:"batch" form:"batch" validate:"required"`
	Vision   string `json:"vision" form:"vision" validate:"required"`
	Mission  string `json:"mission" form:"mission" validate:"required"`
}

type UpdateCandidate struct {
	ID       string `json:"id" form:"id" validate:"required"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
	Batch    int    `json:"batch" form:"batch" validate:"required"`
	Vision   string `json:"vision" form:"vision" validate:"required"`
	Mission  string `json:"mission" form:"mission" validate:"required"`
}
