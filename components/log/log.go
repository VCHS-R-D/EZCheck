package log

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Log(message string) error {

	if _, err := os.Stat("ezcheck-log.txt"); os.IsNotExist(err) {
		_, err := os.Create("ezcheck-log.txt")
		if err != nil {
			return err
		}
	}

	err := NewRecord("ezcheck-log.txt").Prepend(time.Now().Local().Format("01-02-2006 15:04:05")+" "+message)
	
	return err
}

func Read() (string, error) {
	data, err := os.ReadFile("ezcheck-log.txt")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

type Record struct {
	Filename string
	Contents []string
}

func NewRecord(filename string) *Record {
	return &Record{
		Filename: filename,
		Contents: make([]string, 0),
	}
}

func (r *Record) readLines() error {
	if _, err := os.Stat(r.Filename); err != nil {
		return nil
	}

	f, err := os.OpenFile(r.Filename, os.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if tmp := scanner.Text(); len(tmp) != 0 {
			r.Contents = append(r.Contents, tmp)
		}
	}

	return nil
}

func (r *Record) Prepend(content string) error {
	err := r.readLines()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(r.Filename, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	writer.WriteString(fmt.Sprintf("%s\n", content))
	for _, line := range r.Contents {
		_, err := writer.WriteString(fmt.Sprintf("%s\n", line))
		if err != nil {
			return err
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}
