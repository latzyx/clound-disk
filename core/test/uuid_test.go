package test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestUUid(t *testing.T) {
	u := uuid.New().String()
	u = "0"
	fmt.Print(u)
}
