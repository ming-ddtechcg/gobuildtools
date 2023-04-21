package main

import (
	"bufio"
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		os.Exit( 1 )
	}

	input = strings.TrimSuffix(input, "\n")

	defer func() {
		if err := recover(); err != nil {
			os.Exit( 2 )
		}
	}()

	quantity := resource.MustParse( input )

	fmt.Println( quantity.Value() )
}
