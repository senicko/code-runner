# Code runner

Code runner is a go program that takes code as an input and returs the output.

## Examples

Consider simple `Hello, World!` app.

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```

To execute it you need to write a following request to the `stdin`.

_For now `main.go` is an entry file._

```json
{
  "language": "golang",
  "files": [
    {
      "name": "main.go",
      "body": "package main\nimport \"fmt\"\nfunc main(){\nfmt.Println(\"Hello, World!\")\n}"
    }
  ]
}
```

You should get following response written to `stdout`.

```json
{
  "stdout": "Hello, World!\n",
  "stderr": "",
  "exitCode": 0
}
```

If error is caused by user's submitted code it is returned as a regular response written to `stdout`.

```go
package main

import "fmt"

func main() {
  fmt.Println"Hello, World!")
  //         ^
  // missing parenthese
}
```

```json
{
  "stdout": "",
  "stderr": "# command-line-arguments\nfiles/main.go:4:16: syntax error: unexpected jest, expecting comma or )\nfiles/main.go:4:27: newline in string\n",
  "exitCode": 2
}
```

Errors that occure because of invalid request or internal problems are written to `stderr`.

_Example with incorrect request_

```
Error: invalid character 'a' looking for beginning of value
```

Building images

```
docker build -t bee/<language> -f ./images/<image> .
```
