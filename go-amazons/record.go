package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Record struct {
	FromX  rune
	FromY  int
	ToX    rune
	ToY    int
	ArrowX rune
	ArrowY int
}

var (
	recordSlice []Record
)

func AddRecord(fromX, fromY, toX, toY, arrowX, arrowY int) {
	recordSlice = append(recordSlice, Record{
		FromX:  rune(fromX),
		FromY:  fromY,
		ToX:    rune(toX),
		ToY:    toY,
		ArrowX: rune(arrowX),
		ArrowY: arrowY,
	})
}

func Save() {
	filename := fmt.Sprintf("E:/record/AM-先手队B vs 后手队W-先(后)手胜-%v 重庆-2020 CCGC.txt",
		time.Now().Format("2006年01月02日 15时04分"))
	file, _ := os.Create(filename)
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("#[AM][先手参赛队][后手参赛队][后手胜][" +
		time.Now().Format("2006.01.02 15:04") + " 重庆][2020 CCGC];\r\n")
	for i, record := range recordSlice {
		if i%2 == 0 {
			writer.WriteString(fmt.Sprintf("%v ", i/2+1))
		}
		writer.WriteString(fmt.Sprintf("%c%d%c%d(%c%d)", record.FromX, record.FromY, record.ToX,
			record.ToY, record.ArrowX, record.ArrowY))
		if i%2 == 1 {
			writer.WriteString("\r\n")
		} else {
			writer.WriteString(" ")
		}
	}
	writer.Flush()
	recordSlice = recordSlice[:0]
}
