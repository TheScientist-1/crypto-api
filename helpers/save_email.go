package helpers

import (
	"fmt"
	"os"
)

const emailsFile = ".emails"

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
