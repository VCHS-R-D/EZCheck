package log

import (
	"fmt"
	"os"
	"time"
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
	
	_, err = fmt.Fprintln(f, time.Now().Local().Format("2006-01-02 15:04:05") + " " + message)
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

func Read() (string, error) {
	data, err := os.ReadFile("log.txt")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
