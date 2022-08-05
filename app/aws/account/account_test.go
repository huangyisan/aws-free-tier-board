package account

import (
	"fmt"
	"testing"
)

func TestAccount(t *testing.T) {
	a := BuildAWSAccount("aaaa", "cccc", "ap")
	fmt.Println(a.getAccessKey())
}
