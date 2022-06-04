package test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestUUid(t *testing.T) {
	uuid := uuid.New().String()
	fmt.Print(uuid)
}
