package main

import (
  "fmt"
  "net/http"
)

func main() {
  links := []string{
    "https://www.google.com",
    "https://www.facebook.com",
    "https://www.stackoverflow.com",
    "https://www.golang.org",
  }

  var channel = make(chan string)

  for _, link := range links {
    go checkLink(link, channel)
  }

  for i := 0; i < len(links); i++ {
    fmt.Println(<-channel)
  }
}

func checkLink(link string, c chan string) {
  _, err := http.Get(link)
  if err != nil {
    c <- fmt.Sprintf("%v Might be down!", link)
    return
  }
  c <- fmt.Sprintf("%v is up!", link)
}
