package belajar_golang_context

import (
	"context"
	"fmt"
	"testing"
)

// Membuat Context
func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

// Parent dan Child Context
// Context With Value
var key = map[string]interface{}{
	"b": "B",
	"c": "C",
	"d": "D",
	"e": "E",
	"f": "F",
	"g": "G",
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()
	// Child of Context A
	contextB := context.WithValue(contextA, key["b"], key["b"])
	contextC := context.WithValue(contextA, key["c"], key["c"])
	// Child of Context B
	contextD := context.WithValue(contextB, key["d"], key["d"])
	contextE := context.WithValue(contextB, key["e"], key["e"])
	// Child of Context C
	contextF := context.WithValue(contextC, key["f"], key["f"])
	contextG := context.WithValue(contextC, key["g"], key["g"])

	// fmt.Println(contextA)
	// fmt.Println(contextB)
	// fmt.Println(contextC)
	// fmt.Println(contextD)
	// fmt.Println(contextE)
	// fmt.Println(contextF)
	// fmt.Println(contextG)
	// Context Get Value
	fmt.Println(contextF.Value(key["f"])) // dapat value dari contextF
	fmt.Println(contextG.Value(key["c"])) // dapat value dari parent
	fmt.Println(contextD.Value(key["c"])) // tidak dapat, karna beda parent
	fmt.Println(contextE.Value(key["c"])) // tidak dapat, karna beda parent
	fmt.Println(contextA.Value(key["c"])) // tidak dapat, karna parent tidak bisa mengambil data dari child
}
