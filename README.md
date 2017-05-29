# go-open
> Cross-platform CLI application to open files, directories, URLs from the command-line, written in Go.

[![Build Status](https://travis-ci.org/petermbenjamin/go-open.svg?branch=master)](https://travis-ci.org/petermbenjamin/go-open)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Install

### API

```sh
go get github.com/petermbenjamin/go-open
```

### CLI

```sh
go get github.com/petermbenjamin/go-open/cmd/gopen
```

## Usage

### CLI

```sh
$ gopen file.txt
$ gopen https://google.com
$ gopen some_file.txt another_file.txt https://google.com https://github.com
```

### API

```go
import open "github.com/petermbenjamin/go-open"

open.Open([]string{"file.txt"})
open.Open([]string{"https://google.com"})
open.Open([]string{"some_file.txt", "another_file.txt", "https://google.com", "https://github.com"})
```

## License

MIT &copy; [Peter Benjamin](https://petermbenjamin.github.io)
