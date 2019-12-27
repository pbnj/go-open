# go-open

> Cross-platform CLI application to open files, directories, URLs from the command-line, written in Go.

[![Build Status](https://travis-ci.org/petermbenjamin/go-open.svg?branch=master)](https://travis-ci.org/petermbenjamin/go-open)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Install

### API

```sh
go get github.com/pbnj/go-open
```

### CLI

```sh
go get github.com/pbnj/go-open/cmd/gopen
```

## Usage

### API

```go
import open "github.com/pbnj/go-open"

open.Open("file.txt")             // will open file in default text editor
open.Open("https://google.com")   // will open url in default browser
open.Open("image.png")            // will open image in default image editor
```

### CLI

```sh
gopen file.txt
gopen https://google.com
gopen some_file.txt another_file.txt https://google.com https://github.com
```

## License

MIT &copy; 2017 [Peter Benjamin](https://pbnj.github.io)
