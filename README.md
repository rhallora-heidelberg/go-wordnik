# go-wordnik
## (Under Construction)
Unofficial Go library for the [Wordnik](https://www.wordnik.com/) API

[![Build Status](https://travis-ci.org/rhallora-heidelberg/go-wordnik.svg?branch=master)](https://travis-ci.org/rhallora-heidelberg/go-wordnik)

## Basic Usage
```golang
package main

import (
  "fmt"

  "github.com/rhallora-heidelberg/go-wordnik"
)

func main() {
  cl := wordnik.NewClient("<YOUR API KEY>")

  pochade, _ := cl.GetWordOfTheDay("2016-04-03")
  fmt.Println(pochade.Word, pochade.Definitions)
  // pochade [{A rough sketch. wiktionary  noun} {A slight, rough sketch which
  // can easily be erased for correction. century  noun}]

  phytop, _ := cl.Search("phytop")
  for _, res := range phytop.SearchResults {
    fmt.Println(res.Word, res.Count)
  }
  // phytop 0
  // phytoplankton 2192
  // phytophagous 25
  // phytophthora 23
  // phytoplasma 14

}
```

## Configuring Queries
For endpoints with many optional query parameters, such as Search, this project makes use of the [functional options](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html) pattern. This means that optional parameters can be set like this:
```golang
  //...
  // Only words with 6 or 7 characters:
  response, _ := cl.Search("fru", MinLength(6), MaxLength(7))

  // Disable case-sensitive search
  response, _ = cl.Search("ora", CaseSensitive(false))
  //...
```

If you need to set the same parameters on a regular basis, you can define your own functions which fit the QueryOption definition in [queryOptions.go](queryOptions.go):
```golang
  //  type QueryOption func(*url.Values)
  func LongListOfLongNouns() wordnik.QueryOption {
    return func(q *url.Values) {
      q.Set("includePartOfSpeech", "noun")
      q.Set("minLength", "7")
      q.Set("skip", "1")
      q.Set("limit", "100")
    }
  }

```

## Running The Tests
In order to run the included tests, you'll need to provide an API key via the [environment variable](https://www.twilio.com/blog/2017/01/how-to-set-environment-variables.html) WORDNIK_API_KEY. There are a number of ways to do this, but here's a simple one-off for the command line:
```sh
WORDNIK_API_KEY="YOUR_KEY_HERE" go test
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
