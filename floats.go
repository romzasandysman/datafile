package datafile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error){
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File)  {
	fmt.Println("Closing file")
	file.Close()
}

//GetFloats read value float64 from each row of file
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil{
		return nil, err
	}
	defer CloseFile(file)

	var number float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		number, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil{
			return nil, err
		}
		numbers = append(numbers, number)
	}

	if scanner.Err() != nil{
		return nil, err
	}

	return numbers, nil
}
