package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("./evngen.sh")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(stdout)
	scanner.Scan()
	difference := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The difference between the timestamps is %s milliseconds.\n", difference)
}
