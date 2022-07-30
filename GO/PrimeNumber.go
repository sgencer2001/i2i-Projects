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
if (numberOutput<=2){
    fmt.Println(numberOutput , "=> NOT PRIME");
}else {
    x:=0;
    _=x;
    sq_root := int(math.Sqrt(float64(numberOutput)))
    for i := 2; i <= sq_root; i++ {
        
    if numberOutput%i == 0 {
        x=0;
    }else{
        x=1;
    }
    
}
if x==1{
    fmt.Println(numberOutput , "=> PRIME");  
}
if x==0{
    fmt.Println(numberOutput , "=> NOT PRIME");
}
       

    }
    }
    readFile.Close();
}


