# go-open
> Cross-platform CLI application to open files, directories, URLs from the command-line, written in Go.

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Install

### API

```sh
go get github.com/petermbenjamin/go-open
```

### CLI

- Install pre-compiled binaries from [GitHub Releases](https://github.com/petermbenjamin/go-open/releases/latest/)
- Or build from source: 

  ```
  go get github.com/petermbenjamin/go-open/cmd/open
  ```

## Usage

### CLI

```sh
$ open file.txt
$ open https://google.com
$ open some_file.txt another_file.txt https://google.com https://github.com
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
