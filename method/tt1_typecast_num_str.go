package main

import ( 
	"fmt"
	"strconv"
)

func main() {
  var price float64 = 15.00

  fmt.Printf("Price before: %+v\n", price)

  var number int = 10

  price = float64(number)

  fmt.Println("Price after: ", price)
  
  var price2 float64 = 100.00
  price2 = float64(price)
  fmt.Println("Price2: ", price2)

  fmt.Println("============ fmt.Scanf ============")

  fmt.Printf("corretora  total  quantos \n")
  corretora, total, quantos := " "," "," " 
  fmt.Scanf("%s %s %s\n", &corretora, &total, &quantos)  // despresou retornos de quantidade e erro
  floTotal , _ := strconv.ParseFloat(total, 64)
  intQuantos , _ := strconv.Atoi(quantos)
   
  //cabec.totalDaNotaCSinal , _ = strconv.ParseFloat(totalDaNotaCSinal, 64)	
  //cabec.quantosItens , _  = strconv.Atoi(quantosItens)
  // remenber syntax cabec.quantosItens , _ = strconv.ParseInt(quantosItens, 10, 64)										
  // remenber syntax vquantosItens, _ := strconv.Atoi(quantosItens)

  fmt.Printf("%#v %#v %#v %#v %#v \n", corretora, total, floTotal/2, quantos, intQuantos/2 )
}