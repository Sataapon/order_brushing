package shop

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Record struct {
	orderId  string
	shopid   string
	userid   string
	unixTime int64
}

type Mapping struct {
	data map[string][]Record
}

func New(path string) Mapping {
	file, err := os.Open(path)
	check(err)
	return Mapping{data: newMapping(file)}
}

func newMapping(in io.Reader) map[string][]Record {
	data := make(map[string][]Record)
	r := csv.NewReader(in)
	transc, err := r.Read()
	if err != io.EOF {
		check(err)
	}
	validateHeader(transc)
	for {
		transc, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)
		record, ok := newRecord(transc)
		if !ok {
			continue
		}
		data[transc[1]] = append(data[transc[1]], record)
	}
	return data
}

func newRecord(transc []string) (Record, bool) {
	if !validateLength(transc) {
		return Record{}, false
	}
	unixTime, ok := unixTime(transc[3])
	if !ok {
		return Record{}, false
	}
	record := Record{
		orderId:  transc[0],
		shopid:   transc[1],
		userid:   transc[2],
		unixTime: unixTime,
	}
	return record, true
}

func validateHeader(header []string) {
	if !validateLength(header) {
		log.Fatal()
	}
	validHeader := []string{"orderid", "shopid", "userid", "event_time"}
	for i, v := range header {
		if v != validHeader[i] {
			log.Fatalf("invalid header %s != %s\n", v, header[i])
		}
	}
}

func validateLength(transc []string) bool {
	if len(transc) != 4 {
		log.Printf("invalid length: %v\n", transc)
		return false
	}
	return true
}

func unixTime(data string) (int64, bool) {
	eventTime, err := time.Parse("2006-01-02 15:04:05", data)
	if err != nil {
		log.Println(err)
		return 0, false
	}
	return eventTime.Unix(), true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (s Mapping) String() string {
	return fmt.Sprintf("%v", s.data)
}

func (s Mapping) Length() int {
	return len(s.data)
}
