package fileOperations

import (
	"errors"
	"fmt"
	"os"
)

const fileName = "output.txt"


func WriteOutputToFile (name string, place string , time float64 ) {

	info := fmt.Sprintf("Name: %s\nPlace of Work: %s\nTime of Employment: %f\n", name, place, time)
	os.WriteFile(fileName, []byte(info), 0644)
}

func ReadFromFile () (float64, error)  {
	data, err := os.ReadFile(fileName)
	if(err != nil) {
		fmt.Println("Error reading file", err)
		return 1000, errors.New("Error reading file")
	}
	fmt.Println(string(data))
	return 0, nil


}