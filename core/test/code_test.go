package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	s := fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000000))
	fmt.Println(s)
}

func TestAdd(t *testing.T) {
	a := 1 + 1
	fmt.Printf("a: %v\n", a)
}
