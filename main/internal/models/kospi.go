package models

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

var kospiData = *NewKospi()

type Kospi struct {
	Date     map[time.Time]int
	IndexVal []float32
}

func NewKospi() *Kospi {
	//file, err := os.Open("internal/models/kospi/kospi.csv")
	file, err := os.Open("main/internal/models/kospi/kospi.csv")
	if err != nil {
		fmt.Println(err.Error())
	}

	rdr := csv.NewReader(bufio.NewReader(file))
	rows, err := rdr.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	}

	k := Kospi{
		make(map[time.Time]int),
		[]float32{},
	}

	for i := range rows {
		t, err := time.Parse("2006-01-02T15:04:05.000Z", rows[i][0])
		if err != nil {
			fmt.Println(err.Error())
		}
		f, err := strconv.ParseFloat(rows[i][1], 32)
		if err != nil {
			fmt.Println(err.Error())
		}
		k.Date[t] = i
		k.IndexVal = append(k.IndexVal, float32(f))
	}
	return &k
}
