package models

type StartRequest struct {
	DeviceType string `json:"device_type"`
}

type StartResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type StopRequest struct {
	DeviceType string `json:"device_type"`
}

type StopResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
