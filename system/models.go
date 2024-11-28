package system

type ToDo struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  bool   `json:"status"`
}
