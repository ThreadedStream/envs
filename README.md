# envs

Very simple library for setting values associated with environment variables.  

# Usage

Here's a simple example 

```go
package main

import(
  "github.com/ThreadedStream/envs"
  "fmt"
)

type Config struct{
  Host   string `env:"HOST"    fallback:"127.0.0.1"`
  Port   int    `env:"PORT"    fallback:"3000"`
  UseSSL bool   `env:"USE_SSL" fallback:"true"`
}

func main() { 
  var(
    c = &Config{}
    err = envs.Parse(c)
  )
  
  if err != nil{
    panic(err)
  }
  
  fmt.Printf("Host: %s, Port: %d, UseSSL: %t", c.Host, c.Port, c.UseSSL)
  
  // Host: 127.0.0.1, Port: 3000, UseSSL: true 
}
```

# TODOs

[] Add more tests 
  

