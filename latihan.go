package main
import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	hasil := alternateDigitSum(n)
	fmt.Print(hasil)
}
func alternateDigitSum(n int) int {
    var i, c, jum int
    i = 1
    for n > 0{
        c = n%10
        if i % 2 != 0{
            jum += c
        }else{
            c = -c
            jum += c
        }
        i++
        n = n/10
    }
    return jum
}
