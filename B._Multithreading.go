package main
import (
    "bufio"
    "fmt"
    "os"
)
func main() {
    var n int 
    fmt.Scan(&n)
	z := bufio.NewReader(os.Stdin)
	v:=make([]int ,n)
	for i:=0;i<n;i++{
		fmt.Fscan(z,&v[i])
	}
	var i int 
	for i=n-2;i>=0;i--{
		if v[i]>v[i+1]{
			break
		}
	}
	fmt.Println(i+1)
  }

