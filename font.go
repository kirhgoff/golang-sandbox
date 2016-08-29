package main

import (
  "bufio"
  "fmt"
  //"log"
  "os"
  "strings"
  "strconv"
  "bytes"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func processLine(line string) {
  numbers := strings.Split(line, ",")
  fmt.Printf("line is %q\n", numbers)
  for _, numberString := range numbers {
    trimmed := strings.Trim(numberString, " ")
    if len(trimmed) > 0 {
      number, err := strconv.ParseInt(trimmed, 0, 64)
      if err != nil {
        panic(err.Error())
      }
      bitsToChars(int(number), ' ', '*')
    }
  }
}

func bitsToChars(number int, zeroChar rune, oneChar rune) {
  var buffer bytes.Buffer
  var bit = 256
  var symbol rune
  for i := 0; i < 8; i ++ {
    if symbol = zeroChar; number & bit != 0 {
      symbol = oneChar
    }
    buffer.WriteRune(symbol)
    bit /= 2
  }
  fmt.Printf("%s\n", buffer.String()) //, strconv.FormatInt(int64(number), 2)
}

func main() {
  fmt.Println("Reading font")
  fmt.Println("========================")

  printAsciiChars()

  lines, err := readLines("font8.c")
  if err != nil {
    panic(err.Error())
  }

  for _, line := range lines {
    processLine(line)
  }
}

func printAsciiChars() {
  var buffer bytes.Buffer
  for i := 0; i < 256; i ++ {
    buffer.WriteRune(rune(i))
  }
  fmt.Printf("%s\n", buffer.String())
}