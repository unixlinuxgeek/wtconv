### Weight converter


1. Create app folder :

```
mkdir myapp
cd myapp
```


2. Installing dependency:

```
$ go get github.com/unixlinuxgeek/wtconv
```

3. Create app.go file:

```go
package main

import(
	"fmt"
	"github.com/unixlinuxgeek/wtconv"
)

func main() {
	lb := wtconv.KgToLb(1)
	fmt.Printf("1 Kg == %v\n",lb)

	kg := wtconv.LbToKg(1)
	fmt.Printf("1 Lb == %v\n",kg)
} 
```

Run:
```
$ go run ./wtconv.go 1
1 Kg == 2.205 lb
1 Lb == 0.4535147392290249 kg
```