package platforms

import (
	"github.com/gin-gonic/gin"
)

type WindowsService interface {
	Configure() error
	GetStatus() string
}

type windowsService struct {
	router *gin.Engine
}

func NewWindowsService(router *gin.Engine) WindowsService {
	return &windowsService{
		router: router,
	}
}

func (ws *windowsService) Configure() error {
	// Windows-specific configuration
	return nil
}

func (ws *windowsService) GetStatus() string {
	return "Windows service running"
}
