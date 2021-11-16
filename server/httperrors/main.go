package httperrors

type Body struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type Error struct {
	Status int
	Body   Body
}
