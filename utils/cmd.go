package utils

import (
	"bufio"
	"fmt"
	"os/exec"
	"sync"
)

func DoCmd(cmdStr string, args ...string) error {
	cmd := exec.Command(cmdStr, args...)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	// Execute
	if err := cmd.Start(); err != nil {
		return err
	}
	//
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			fmt.Println("Stderr:", scanner.Text())
		}
	}()

	if err := cmd.Wait(); err != nil {
		fmt.Println("El comando no se ejecutó con éxito:", err)
	}

	wg.Wait()

	return nil
}
