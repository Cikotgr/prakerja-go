package vote

type VoteResponse struct {
	ID         string `json:"id"`
	FullName   string `json:"full_name"`
	TotalVotes int    `json:"total_votes"`
}
