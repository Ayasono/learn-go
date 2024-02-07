package main

import "fmt"

func main() {
  cards := newDeck()

  cards.shuffle()

  err := cards.saveToFile("cards_cache")
  if err != nil {
    return
  }
  cards.print()

  play()
}

func play() {
  type person struct {
    name string
    age  int
  }

  p1 := person{name: "Alex", age: 20}

  p2 := person{name: "Bob", age: 30}

  fmt.Println(p1, &p2)

  //p := map[string]person{
  //  "Alex": p1,
  //  "Bob":  person{name: "Bob", age: 30},
  //}

  p := make(map[string]string)
  p["Alex"] = "20"
  p["Bob"] = "30"

  for k, v := range p {
    fmt.Println(k, v)
  }

  type bot interface {
    getGreeting() string
  }
}
