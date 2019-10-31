package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var bigDigit = map[string][]string{
	"0": {" 00 ",
		"0  0",
		"0  0",
		"0  0",
		" 00 "},
	"1": {"  1 ",
		" 11 ",
		"  1 ",
		"  1 ",
		" 111"},
	"2": {" 22 ",
		"2  2",
		"  2 ",
		"2   ",
		"2222"},
	"3": {" 33 ",
		"3  3",
		"  3 ",
		"3  3",
		" 33 "},
	"4": {"4  4",
		"4  4",
		"4444",
		"   4",
		"   4"},
	"5": {"5555",
		"5   ",
		"555 ",
		"   5",
		"555 "},
	"6": {" 66 ",
		"6   ",
		"666 ",
		"6  6",
		" 66 "},
	"7": {"7777",
		"   7",
		"  7 ",
		" 7  ",
		"7   "},
	"8": {" 88 ",
		"8  8",
		" 88 ",
		"8  8",
		" 88 "},
	"9": {" 99 ",
		"9  9",
		" 999",
		"   9",
		" 99 "},
	"-": {"    ",
		"    ",
		"----",
		"    ",
		"    "},
}

func TimeTrack(start time.Time)time.Duration {
	timeRun := time.Since(start)
	return timeRun + 0
	// fmt.Print("run time: ", timeRun, "\n")
}

func main() {
	if len(os.Args) == 1 {
		os.Exit(1)
	}
	defer TimeTrack(time.Now())

	var result [5]string
	repetition := 1
	for i := 0; i < repetition; i++ {
		//result = [5]string{"", "", "", "", ""}
		for _, dig := range os.Args[1] {
			// fmt.Print(string(dig) + "\n")
			if key, ok := bigDigit[string(dig)]; ok {
				// fmt.Print(key[0] + " \n")
				//for index, _ := range result {
				//	result[index] += key[index]
				//}
				result[0] += key[0]
				result[1] += key[1]
				result[2] += key[2]
				result[3] += key[3]
				result[4] += key[4]
			}
		}
	}

	fmt.Println(strings.Join(result[:], "\n"))

	//for _, line :=range result{
	//	fmt.Println(len(line))
	//}
}
