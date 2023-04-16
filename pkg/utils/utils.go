package utils

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func ParseBody(c *gin.Context, x interface{}) error {
	if reader, err := io.ReadAll(c.Request.Body); err == nil {
		if err := json.Unmarshal([]byte(reader), x); err != nil {
			return fmt.Errorf("could not umarshal request body: %v", err)
		}
	}
	return nil
}
