package common

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var (
	Validate = validator.New()
)

type GenericResponse struct {
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
}

func WriteResp(c *gin.Context, data interface{}, msg string, statusCode int) {
	resp := GenericResponse{
		Data:       data,
		Message:    msg,
		StatusCode: statusCode,
	}
	c.JSON(statusCode, &resp)
}

// Use this for response struct so the time format is proper
type CustomTime time.Time

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, time.Time(ct).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}
