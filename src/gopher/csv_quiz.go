package main

import(
  "fmt"
  "os"
  "encoding/csv"
  "flag"
  "log"
  "strings"
  "time"
)

func main() {
  // get filename from flag
  filename := flag.String("file", "filename.csv", "enter csv file of questions")
  timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
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

  timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

  // could break this out into its own function to make it easier to test
  count := 0
  for i, p := range problems {
    fmt.Printf("Problem #%d: %s = ", i+1, p.q)
    answerChan := make(chan string)
    go func() {
      var answer string
      fmt.Scanf("%s\n", &answer)
      answerChan <- answer
    }()

    select {
      case <-timer.C:
        fmt.Printf("\nYou got %d out of %d problems\n", count, len(lines))
        return
      case answer := <-answerChan:
        if answer == p.a {
          fmt.Println("Correct!")
          count++
        } else {
          fmt.Println("Wrong!")
        }
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
