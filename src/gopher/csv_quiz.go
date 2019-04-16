package main

import(
  "fmt"
  "os"
  "encoding/csv"
  "flag"
  "log"
)

func main() {

  // get filename from flag
  filename := flag.String("f", "", "enter csv filename of questions")

  // open file
  file, err := os.Open(&filename)
  if err != nil {
    log.Fatal(err)
  }

  // parse file with csv package
  r := csv.NewReader(file)

  // for question in file
  for {
    record, err := r.Read()
    if err == io.EOF {
      break;
    }
    if err != nil {
      log.Fatal(err)
    }

    fmt.Println(record)
  }
  //   ask user question
  //   get input
  //   check input
}
