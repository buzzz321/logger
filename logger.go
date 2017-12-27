package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getToken(line string, start string, stop string) (int, string) {
	tStart := strings.Index(line, start) + len(start)
	tStop := strings.Index(line[tStart:], stop) + tStart
	token := line[tStart:tStop]

	//fmt.Printf("tStart = %d tStop = %d\n", tStart+0, tStop+0)
	return tStop, token
}

func main() {
	fmt.Println("vim-go")
	file, err := os.Open("../../../c++/logger/log.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		advance := 0
		line := scanner.Text()

		tStop, timeStamp := getToken(line, "[", "]")
		advance += tStop
		tStop, logType := getToken(line[advance:], ":", ":")
		advance += tStop
		tStop, cpuId := getToken(line[advance:], "cpu_id = ", " }")
		advance += tStop
		tStop, fileName := getToken(line[advance:], "file = \"", "\"")
		advance += tStop
		tStop, lineNo := getToken(line[advance:], "line = ", " ,")
		advance += tStop
		tStop, msg := getToken(line[advance:], "msg = \"", "\" }")
		advance += tStop

		fmt.Println(timeStamp, logType, cpuId, fileName, lineNo, msg)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
