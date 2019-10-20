package main

import (
	"fmt"
	"os"
	//"strings"
	"time"
	//"reflect"
)
var bigDigByte = map[byte][][]byte{
	'0': {[]byte{' ', '0','0',' '},
		  []byte{'0', ' ',' ','0'},
		  []byte{'0', ' ',' ','0'},
		  []byte{'0', ' ',' ','0'},
		  []byte{' ', '0','0',' '}},
	'1': {[]byte{' ', ' ','1',' '},
		  []byte{' ', '1','1',' '},
		  []byte{' ', ' ','1',' '},
		  []byte{' ', ' ','1',' '},
		  []byte{' ', '1','1','1'}},
	'2': {[]byte{' ', '2','2',' '},
		  []byte{'2', ' ',' ','2'},
		  []byte{' ', ' ','2',' '},
		  []byte{' ', '2',' ',' '},
		  []byte{'2', '2','2','2'}},
	'3': {[]byte{' ', '3','3',' '},
		  []byte{'3', ' ',' ','3'},
		  []byte{' ', ' ','3',' '},
		  []byte{'3', ' ',' ','3'},
		  []byte{' ', '3','3',' '}},
	'4': {[]byte{'4', ' ',' ','4'},
		  []byte{'4', ' ',' ','4'},
		  []byte{'4', '4','4','4'},
		  []byte{' ', ' ',' ','4'},
		  []byte{' ', ' ',' ','4'}},
	'5': {[]byte{'5', '5','5','5'},
		  []byte{'5', ' ',' ',' '},
		  []byte{'5', '5','5',' '},
		  []byte{' ', ' ',' ','5'},
		  []byte{'5', '5','5',' '}},
	'6': {[]byte{' ', '6','6',' '},
		  []byte{'6', ' ',' ',' '},
		  []byte{'6', '6','6',' '},
		  []byte{'6', ' ',' ','6'},
		  []byte{' ', '6','6',' '}},
	'7': {[]byte{'7', '7','7','7'},
		  []byte{' ', ' ',' ','7'},
		  []byte{' ', ' ','7',' '},
		  []byte{' ', '7',' ',' '},
		  []byte{'7', ' ',' ',' '}},
	'8': {[]byte{' ', '8','8',' '},
		  []byte{'8', ' ',' ','8'},
		  []byte{' ', '8','8',' '},
		  []byte{'8', ' ',' ','8'},
		  []byte{' ', '8','8',' '}},
	'9': {[]byte{' ', '9','9',' '},
		  []byte{'9', ' ',' ','9'},
		  []byte{' ', '9','9','9'},
		  []byte{' ', ' ',' ','9'},
		  []byte{' ', '9','9',' '}},
	'-': {[]byte{' ', ' ',' ',' '},
		  []byte{' ', ' ',' ',' '},
		  []byte{'-', '-','-','-'},
		  []byte{' ', ' ',' ',' '},
		  []byte{' ', ' ',' ',' '}},
}

func TimeTrack(start time.Time) {
	timeRun := time.Since(start)
	fmt.Print("run time: ", timeRun, "\n")

}

func main() {
	if len(os.Args) == 1 {
		os.Exit(1)
	}
	defer TimeTrack(time.Now())
	//countDigit := len(os.Args[1])
	//repetition := 1000000
	repetition := 50000
	var result [5][]byte
	//for j,_ := range result{
	//	result[j] = make([]byte, 4 * repetition * countDigit)
	//}
	for i := 0; i < repetition; i++ {
		for _, dig := range os.Args[1] {
		//for cmdIndex, dig := range os.Args[1] {
			if key, ok := bigDigByte[byte(dig)]; ok {
				for index, _ := range result {
					//copy(result[index][(i * countDigit *4) + (cmdIndex *4):], key[index])
					result[index] = append(result[index], key[index]...)
				}
			}
		}
	}
	//size := reflect.TypeOf(result)
	fmt.Println(len(result[0]))
	fmt.Println(cap(result[0]))
	fmt.Println(len(result[1]))
	fmt.Println(len(result[2]))
	fmt.Println(len(result[3]))
	fmt.Println(len(result[4]))
	//fmt.Println(result)
	//for i := 0; i <5; i++{
	//	fmt.Println(string(result[i]))
	//}
}
