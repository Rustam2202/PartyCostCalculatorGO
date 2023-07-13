package handlers

type ErrorResponce struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
