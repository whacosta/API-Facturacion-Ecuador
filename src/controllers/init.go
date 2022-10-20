package controllers

// Controller type
type Controller struct {
}

// NewController init controller
func NewController() *Controller {
	return &Controller{}
}

// Message example message
type Message struct {
	Message string `json:"message" example:"message"`
}
