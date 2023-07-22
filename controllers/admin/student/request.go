package student

type GetRequestParam struct {
	NIM string `query:"nim" validate:"omitempty"`
}

type CreateStudent struct {
	ID  string `json:"id" form:"id"`
	NIM int    `json:"nim" form:"nim"`
}

type UpdateStudent struct {
	ID  string `json:"id" form:"id"`
	NIM int    `json:"nim" form:"nim"`
}
