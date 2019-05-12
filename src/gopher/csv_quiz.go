package main

import(
  "fmt"
  "os"
  "encoding/csv"
  "flag"
  "log"
  "io"
  "bufio"
  "strings"
)

func main() {
  reader := bufio.NewReader(os.Stdin)

  // get filename from flag
  filename := flag.String("file", "filename.csv", "enter csv file of questions")
  flag.Parse()
  fmt.Println("filename: " + *filename)
  // open file
  file, err := os.Open(*filename)
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

    fmt.Printf("%s: ", record[0])
    input, _ := reader.ReadString('\n')

    if strings.TrimRight(input, "\n") == record[1] {
      fmt.Println("Correct!")
    } else {
      fmt.Println("Wrong!")
    }
  }
  //   ask user question
  //   get input
  //   check input
}
