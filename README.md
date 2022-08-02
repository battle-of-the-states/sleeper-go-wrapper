# sleeper-go-wrapper

Sleeper.app Fantasy Football read only API wrapper written in Go

## Simple usage

> Install the module

```bash
go get github.com/battle-of-the-states/sleeper-go-wrapper
```

> Import the module

```go
import sleeper "github.com/battle-of-the-states/sleeper-go-wrapper"

func main() {
    hClient := http.Client{Timeout: time.Duration(30) * time.Second}
	client, err := sleeper.NewClient(&hClient)
	if err != nil {
		panic(err)
	}

    user, err := client.GetUser("kylewithanr")

	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
```

## Sleeper.app API

> https://docs.sleeper.app
