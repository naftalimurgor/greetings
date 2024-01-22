## Shared Go package
1. mkdir `greetings`
2. cd `/greetings`
3. Run:
```sh
go mod init github.com/greetings # should match github repo for the package
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

## Setup with Git- has to have a package version and or release
```sh
git tag --version
git push origin v0.1.0 # can be any version number, see https://go.dev/blog/publishing-go-modules
```

That's how you publish go modules for reuse on other projects!