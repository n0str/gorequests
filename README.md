GoRequests: HTTP for Humans
==================================================

GoRequests is the library that simplifies HTTP requests. It was developed  
in the image of famous Python requests library.

It is as simple as:
```go
package main

import (
	"github.com/n0str/gorequests"
	"log"
)

func main() {
	http := http_session.New()
	r := http.EasyRequest("GET", "https://httpbin.org/ip")
	log.Printf("Response: %v", r.EasyJson())
	log.Printf("Response: %v", r.EasyString())
}

```

Results:

```go
>>> Response: map[origin:178.70.140.98, 178.70.140.98]
>>> Response: {
  "origin": "178.70.140.98, 178.70.140.98"
}

```