package student

type CreateStudent struct {
	ID  string `json:"id" form:"id"`
	NIM int    `json:"nim" form:"nim"`
}

type UpdateStudent struct {
	ID  string `json:"id" form:"id"`
	NIM int    `json:"nim" form:"nim"`
}
