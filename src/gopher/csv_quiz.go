package main

import(
  "fmt"
  "os"
  "encoding/csv"
  "flag"
  "log"
  "strings"
)

func main() {
  // get filename from flag
  filename := flag.String("file", "filename.csv", "enter csv file of questions")
  flag.Parse()

  // open file
  file, err := os.Open(*filename)
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  // parse file with csv package
  r := csv.NewReader(file)
  lines, err := r.ReadAll()

  if err  != nil {
    exit("Failed to parse the provided CSV file.\n")
  }

  problems := parseLines(lines)

  // could break this out into its own function to make it easier to test
  count := 0
  for i, p := range problems {
    fmt.Printf("Problem #%d: %s = ", i+1, p.q)
    var answer string
    fmt.Scanf("%s\n", &answer)
    if answer == p.a {
      fmt.Println("Correct!\n")
      count++
    } else {
      fmt.Println("Wrong!\n")
    }
  }

  fmt.Printf("You got %d out of %d problems.\n", count, len(lines))
}

type problem struct {
  q string
  a string
}

func parseLines(lines [][]string) []problem {
  ret := make([]problem, len(lines))
  for i, line := range lines {
    ret[i] = problem{
      q: line[0],
      a: strings.TrimSpace(line[1]),
    }
  }
  return ret
}

func exit(msg string) {
  fmt.Println(msg)
  os.Exit(1)
}
