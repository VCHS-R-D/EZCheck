package log

import (
	"fmt"
	"os"
)

func Log(message string) error {

	if _, err := os.Stat("log.txt"); os.IsNotExist(err) {
		_, err := os.Create("log.txt")
		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	
	_, err = fmt.Fprintln(f, message)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
