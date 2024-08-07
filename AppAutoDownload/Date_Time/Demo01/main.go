// https://vocus.cc/article/65041827fd8978000174d093
// https://www.dotnetperls.com/time-go
// https://juejin.cn/post/7114511786388226085
// https://www.jb51.net/article/275527.htm
package main

import (
	"fmt"
	"time"
)

// Use layout string for time format.
const layout = "200601021504"

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)

	// Place now in the string.
	t := time.Now()
	filename := "file-" + t.Format(layout) + ".txt"
	fmt.Println("Name:", filename)
}
