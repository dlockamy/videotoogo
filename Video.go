package videotogo

type Video struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	BlockHash   string `json:"blockhash"`
}
