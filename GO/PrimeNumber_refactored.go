
package main

import ( 
    "fmt" 
    "os" 
    "strconv"
	"bufio"
    "math"
 )

func main() {  
readFile, errorOutput := os.Open("GO/numbers.txt");
if errorOutput != nil {
    fmt.Println(errorOutput)
}
fileScanner := bufio.NewScanner(readFile)
for fileScanner.Scan() {
    numberOutput , errorOutput := strconv.Atoi(fileScanner.Text())
    if (errorOutput != nil) {
        fmt.Println(errorOutput)
    }
    fmt.Println(numberOutput, isPrimeNumber(numberOutput))

}
readFile.Close();
}
func isPrimeNumber(input int) string {
    
  if(input<2){
    return "=> NOT PRIME"
  }
  sq_root := int(math.Sqrt(float64(input)))
	for i := 2; i <= sq_root; i++ {
		if input%i == 0 {
			return "=> NOT PRIME"
		}
	}
    return "=>PRIME"

} 