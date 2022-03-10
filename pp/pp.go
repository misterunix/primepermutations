package main

import "fmt"

func main() {

	for i := 1; i <= 1019; i++ {
		fmt.Printf("echo 'scale=30000; %d/1019' | bc | awk 'BEGIN{FS=\".\"}{ print $2}' > %d.txt\n", i, i)
	}
}
