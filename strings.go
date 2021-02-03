package datafile

import (
	"bufio"
	"os"
)

//GetStrings read value string from each row of file
func GetStrings(fileName string) ([]string, error){
	var lines []string
	file, err := os.Open(fileName)
	if err != nil{
		return nil, err
	}

	var line string

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line = scanner.Text()
		lines = append(lines, line)
	}

	err = file.Close()
	if err != nil{
		return lines, err
	}
	if scanner.Err() != nil{
		return nil, err
	}

	return lines, nil
}
