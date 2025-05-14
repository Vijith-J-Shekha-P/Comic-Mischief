package main

import "fmt"

// Define the function outside main
func printComicDetails(title, writer, artist, publisher string, year, pageNumber, age int, grade float32, genre string) {
  fmt.Println(title, "written by", writer, "drawn by", artist, "Published by", publisher, "Year of publication", year, "Number of pages", pageNumber, "Condition grade", grade, "Genre", genre, "Book age", age)
}

func main() {
  var publisher, writer, artist, title, genre string
  var year, pageNumber, age int
  var grade float32
  const currentYear = 2025

  // First comic
  title = "Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  grade = 6.5
  genre = "Horror"
  age = currentYear - year

  printComicDetails(title, writer, artist, publisher, year, pageNumber, age, grade, genre)
  fmt.Println()

  // Second comic
  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  year = 2013
  pageNumber = 160
  grade = 9.0
  genre = "Fantasy"
  age = currentYear - year

  printComicDetails(title, writer, artist, publisher, year, pageNumber, age, grade, genre)
}
