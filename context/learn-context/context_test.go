package learncontext

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)

	// membuat child dari context
	//args (parent , key, value)
	contextB := context.WithValue(background, "b", "B")
	contextC := context.WithValue(contextB, "c", "C")
	contextD := context.WithValue(contextB, "d", "D")
	fmt.Println(background)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)

	//untuk mencari value tidak bisa dari parent
	// tapi child bisa mencari value ke parent
	fmt.Println(background.Value("b"))
	fmt.Println(contextB.Value("b"))
	fmt.Println(contextC.Value("b"))

}
