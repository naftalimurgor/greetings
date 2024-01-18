## Shared Go package
1. mkdir `greetings`
2. cd `/greetings`
3. Run:
```sh
go mod init github.com/naftalimurgor.com/greetings # should match github repo for the package
```
4. Add package code and push

## Using the package in another codebase/package:

```go
import (
  "github.com/naftalimurgor.com/greetings"
  "fmt"
)

func main() {
  msg := greetings.Hello("Ruth")
  fmt.Println(msg)
}
```