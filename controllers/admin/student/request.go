package student

type GetRequestParam struct {
	NIM string `query:"nim" validate:"omitempty"`
}

type CreateStudent struct {
	ID     string `json:"id" form:"id" validate:"required"`
	NIM    int    `json:"nim" form:"nim" validate:"required,min=202410102000,max=232410102100"`
	RoleId int    `json:"role_id" form:"role_id" validate:"required,number"`
}

type UpdateStudent struct {
	ID  string `json:"id" form:"id"`
	NIM int    `json:"nim" form:"nim" validate:"required,min=202410102000,max=232410102100"`
}
