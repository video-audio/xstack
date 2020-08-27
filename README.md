# xstack

[![GoDev][godev-image]][godev-url]
[![Build Status][build-image]][build-url]
[![Coverage Status][coverage-image]][coverage-url]
[![Go Report Card][goreport-image]][goreport-url]

Xstack is library for [xstack layouts][ffmpeg-wiki-xstack] generation written in [Go][go].

## Quick-start

```bash
$ sudo chmod 777 /var/run/smcroute
```

```go
package main

import (
  "bytes"
  "fmt"

  "github.com/video-audio/xstack"
)

// 0_0|w0_0|w0+w1_0|0_h0|w0_h0|w0+w1_h0|0_h0+h1|w0_h0+h1|w0+w1_h0+h1
func main() {
  fmt.Println(xstack.Layout(9))

  // or
  b := bytes.Buffer{}
  xstack.LayoutTo(&b, 9)

  fmt.Println(b.String())
}
```

[godev-image]: https://img.shields.io/badge/go.dev-reference-5272B4?logo=go&logoColor=white
[godev-url]: https://pkg.go.dev/github.com/video-audio/xstack

[build-image]: https://travis-ci.com/video-audio/xstack.svg?branch=master
[build-url]: https://travis-ci.com/video-audio/xstack

[coverage-image]: https://coveralls.io/repos/github/video-audio/xstack/badge.svg?branch=master
[coverage-url]: https://coveralls.io/github/video-audio/xstack?branch=master

[goreport-image]: https://goreportcard.com/badge/github.com/video-audio/xstack
[goreport-url]: https://goreportcard.com/report/github.com/video-audio/xstack

[go]: http://golang.org/
[ffmpeg-wiki-xstack]: https://trac.ffmpeg.org/wiki/Create%20a%20mosaic%20out%20of%20several%20input%20videos%20using%20xstack
