package utils

import (
	"fmt"
	"testing"
)

func TestStringToInt(t *testing.T) {
	msg := "123"
	msgint := StringToInt(msg)
	fmt.Println(msgint)
}
