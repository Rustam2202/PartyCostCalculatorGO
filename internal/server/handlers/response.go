package handlers

type UpdateDeleteResponse struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

type ErrorResponce struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
