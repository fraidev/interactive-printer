package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	// r, w := io.Pipe()
	cmd := exec.Command("wget", "https://storage.googleapis.com/download/storage/v1/b/marigold-snapshot-bucket/o/2022.04.07%2FTEZOS_HANGZHOUNET_2021-11-04T15:00:00Z-BLraHczggjqm2wbSMotYCycuUwEWtKWKgNzw5voU98g5iBv2k4X-779008.rolling?generation=1649359511387050&alt=media")

	// cmd := exec.Command("myInteractiveCmd")
	// cmd.Stdin = r
	// go func() {
	// 	fmt.Fprintf(w, "John Doe\n")
	// 	fmt.Fprintf(w, "New York\n")
	// 	fmt.Fprintf(w, "35\n")
	// 	w.Close()
	// }()
	// cmd.Start()
	// cmd.Wait()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()
}
