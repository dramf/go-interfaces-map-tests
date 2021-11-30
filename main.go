package main

import "fmt"

type myfunc func() string

func main() {
	fmt.Println("Go Interfaces map tests")

	m := map[string]interface{}{
		"str": "str1",
		"funcWrong": func() string {
			return "funcWrongInstance"
		},
		"funcCorrect": myfunc(func() string {
			return "funcCorrectInstance"
		}),
	}

	for k, v := range m {
		if value, ok := v.(string); ok {
			fmt.Printf("%q is string %q\n", k, value)
		}

		if value, ok := v.(myfunc); ok {
			fmt.Printf("%q is myfunc %q\n", k, value())
		}
	}
}
