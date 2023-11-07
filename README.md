# launchbox-go-sdk

## Installation 

```shell
go get github.com/launchboxio/launchbox-go-sdk 
```

## Usage

```go 
package main

import (
  "github.com/launchboxio/launchbox-go-sdk/config"
  "github.com/launchboxio/launchbox-go-sdk/service/project"
)

func main() {
  conf, err := config.Default()
  if err != nil {
    log.Fatal(err)
  }

  projectSdk := project.New(conf)
  out, err := projectSdk.Create(&project.CreateProjectInput{
	  
  })
  
  if err != nil {
	  log.Fatal(err)
  }
  fmt.Println(out)
}
```
