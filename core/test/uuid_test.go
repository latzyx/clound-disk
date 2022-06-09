package test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestUUid(t *testing.T) {
	uuid := uuid.New().String()
	uuid = "0"
	fmt.Print(uuid)
}
