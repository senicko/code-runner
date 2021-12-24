# Bee

Bee is a go program running inside containers created by [Bee Hive](https://github.com/senicko/bee-hive). It's task is to
read the config, create source files, execute the program and return the output.

_Currently it supports only golang code execution._

## How does it work

For example consider simple `Hello, World!` app.

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```

To execute it inside a container that is running a bee you need to write a following request to it's stdin.

_For now `main.go` is an entry file._

```json
{
  "files": [
    {
      "name": "main.go",
      "body": "package main\nimport \"fmt\"\nfunc main(){\nfmt.Println(\"Hello, World!\")\n}"
    }
  ]
}
```

Bee should write following response to it's stdout.

```json
{
  "stdout": "Hello, World!\n",
  "stderr": "",
  "exitCode": 0
}
```
