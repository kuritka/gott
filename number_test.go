package gott

import (
	"fmt"
	"testing"
)

func TestXyz (t *testing.T){

	for key, _ := range cultures {
		fmt.Println("xx.Add(\""+key+"\");")
	}

}
