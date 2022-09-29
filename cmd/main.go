package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yudgxe/cut"
	"github.com/yudgxe/cut/value"
)

var (
	f value.MasValue
	d string
	s bool
)

func init() {
	flag.Var(&f, "f", "выбрать поля (колонки)")
	flag.StringVar(&d, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&s, "s", false, "только строки с разделителем")
}

func main() {
	flag.Parse()

	path := os.Args[len(os.Args)-1]

	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	r, err := cut.Read(file, f, d, s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(r)

}
