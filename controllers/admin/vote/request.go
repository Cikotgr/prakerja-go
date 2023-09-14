package vote

type CreateVote struct {
	ID          string `json:"id" form:"id" validate:"required"`
	CandidateId string `json:"candidate_id" form:"candidate_id" validate:"required"`
	StudentId   string `json:"student_id" form:"student_id" validate:"required"`
}
