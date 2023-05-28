package helpers

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
)

const emailsFile = "emails"

func SaveEmail(email string) error {
	file, err := os.OpenFile(emailsFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(email + "\n")
	if err != nil {
		return err
	}

	fmt.Println(email)

	return nil
}
func ExtractEmailsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	emailRegex := regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`)
	emails := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		matches := emailRegex.FindAllString(line, -1)
		emails = append(emails, matches...)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return emails, nil
}