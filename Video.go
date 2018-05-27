package videotogo

type Video struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	BlockHash   string `json:"blockhash"`
}
