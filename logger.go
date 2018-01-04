package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type parser struct {
	advance int
	err     error
}

func (p *parser) getToken(line string, start string, stop string) string {
	var retVal string
	if p.err == nil {
		tStart := strings.Index(line[p.advance:], start) + len(start) + p.advance
		tStop := strings.Index(line[tStart:], stop) + tStart
		//fmt.Printf("tStart = %d tStop = %d\n", tStart+0, tStop+0)
		if tStart == -1 || tStop == -1 {
			p.err = errors.New("item not found")
			return ""
		}
		p.advance = tStop
		retVal = line[tStart:tStop]
	}
	return retVal
}

func (p *parser) parseLine(line string) string {
	timeStamp := p.getToken(line, "[", "]")
	logType := p.getToken(line, ":", ":")
	cpuID := p.getToken(line, "cpu_id = ", " }")
	fileName := p.getToken(line, "file = \"", "\"")
	lineNo := p.getToken(line, "line = ", " ,")
	msg := p.getToken(line, "msg = \"", "\" }")
	p.advance = 0

	if p.err == nil {
		//fmt.Println(timeStamp, logType, cpuID, fileName, lineNo, msg)
		return fmt.Sprintf("%s %s %s %s %s %s\n", timeStamp, logType, cpuID, fileName, lineNo, msg)
	}
	return (line)
}

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	parse := parser{}
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(parse.parseLine(line))

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
