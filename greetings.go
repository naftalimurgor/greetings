package greetings

import (
	"fmt"
	"errors"
)

func Hello(name string) (string, error)  {
  message := fmt.Sprintf("Hi %v. Welcome!", name)
	if name == "" {
		return name, errors.New("Empty name")
	}
	return message, nil	
}
