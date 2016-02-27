

####short url usage example
```
package main

import (
	"fmt"
	"os"

	"./short_url"
)

func main() {
	//fmt.Println(short_url.IntToBase62(10))
	if len(os.Args) != 2 {
		fmt.Printf("%+v\n", os.Args)
		fmt.Println("use as go run " + os.Args[0] + " url")
		os.Exit(0)
	}
	/*id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("str to int failed")
	}
	*/
	fmt.Println(short_url.LongToShort(os.Args[1]))
}
```
