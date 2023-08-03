package student

type GetRequestParam struct {
	NIM string `query:"nim" validate:"omitempty"`
}

type CreateStudent struct {
	ID     string `json:"id" form:"id" validate:"required"`
	NIM    string `json:"nim" form:"nim" validate:"required,number,min=12"`
	RoleId int    `json:"role_id" form:"role_id" validate:"required,number"`
}

type UpdateStudent struct {
	ID     string `json:"id" form:"id"`
	NIM    string `json:"nim" form:"nim"`
	RoleId int    `json:"role_id" form:"role_id" validate:"required,number"`
}
