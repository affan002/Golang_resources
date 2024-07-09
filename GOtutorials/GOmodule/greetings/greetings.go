package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}


    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	message_map := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		message_map[name] = message
	}
		
	return message_map, nil
	

}

func randomFormat() string {
	slice := []string{
		"Hi %v, welcome!",
		"Great to see you! %v", 
		"Hello %v",
	}

	return slice[rand.Intn(len(slice))]
}
