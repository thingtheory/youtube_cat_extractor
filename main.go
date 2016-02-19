package main

import (
  "fmt"
  "log"

  "github.com/PuerkitoBio/goquery"
)

func main() {
  channels := []string{}
  baseUrl := "https://www.youtube.com"
  // arbitrary limit until some end check is implemented
  pageLimit := 10

  for i := 1; i <= pageLimit; i++ {
    url := fmt.Sprintf("%s/channels?q=cats&page=%d", baseUrl, i)
    doc, err := goquery.NewDocument(url)
    if err != nil {
      log.Fatal(err)
    }

    channels = append(getChannels(doc))
  }

  fmt.Println(len(channels))

  videos := getVideos(channels, baseUrl)

  fmt.Println(len(videos))
}

func getChannels(categoryPage *goquery.Document) []string {
  channels := []string{}
  categoryPage.Find(".yt-gb-shelf-hero >a").Each(func(i int, s *goquery.Selection) {
    href, found := s.Attr("href")
    if found != true {
      log.Fatal(found)
    }
    fmt.Println(href)
    channels = append(channels, href)
  })
  return channels
}

func getVideos(channels []string, baseUrl string) []string {
  videos := []string{}
  for _, channelPath := range channels {
    url := fmt.Sprintf("%s%s", baseUrl, channelPath)
    doc, err := goquery.NewDocument(url)
    if err != nil {
      log.Fatal(err)
    }

    doc.Find(".channels-content-item .spf-link >a").Each(func(i int, s *goquery.Selection) {
      href, found := s.Attr("href")
      if found != true {
        log.Fatal(found)
      }
      fmt.Println(href)

      videos = append(videos, href)
    })
  }

  return videos
}
