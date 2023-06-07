package main

import (
	"fmt"
	"strconv"
	"strings"
)

type LinhaNota struct { // linha de uma nota de corretagem
	cvlnota string  // "C" ou "V" compra ou venda na linha         - compra (débito) sera negativo  venda (crédito) sera positivo
	pplnota string  // ativo ("PaPel") transacionado ns linha
	qtlnota int     // quantidade do ativo na linha                - herda o sinal de C/V "-" na compra e "+" na venda
	unlnota float64 // preço medio da aquisição (unitario da linha)- sempre positivo
	epcnota float64 // coRRetagem específica deste papel na linha  - SEMPRE NEGATIVA conhece o valor - fora do calculo ponderado
	txlnota float64 // taxa desta transac. deste ativo             - SEMPRE NEGATIVA a PRÓPRIA INCOGNITA a obter através de ponderação
	ttlnota float64 // custo deste ativo na linha                  - positivo (credito) na venda   negativo (débito) na compra
}

// func Scanf(format string, a ...any) (n int, err error)
// n ==> It returns the number of items successfully scanned
// n ==> If that is less than the number of arguments, err will report why.
func main() {
	var nota []LinhaNota
	var linha LinhaNota
	var erro error
	fmt.Printf("prompt to enter, use space as delimiter\n")
	corretora := "NECTON/INTER"
	numNotaCorr := "123456789"
	dataPregao := "23/04/2023"
	totalDaNotaCSinal := "+/-10000.00"
	quantosItens := "2" // how much itens

	fmt.Printf("corretora nt.Corret dt.Pregao +/-valTotal itens\n")
	//"NECTON/INTER" "123456789" "23/04/2023" "+/-10000.00"  "2"
	n, err := fmt.Scanf("%s %s %s %s %s\n", &corretora, &numNotaCorr,
		&dataPregao, &totalDaNotaCSinal, &quantosItens)
	if err != nil {
		panic(err)
	}
	if n != 5 {
		panic("must read 5 variables")
	}

	fmt.Println(corretora, numNotaCorr, dataPregao, totalDaNotaCSinal, quantosItens)

	vquantosItens, _ := strconv.Atoi(quantosItens)
	fmt.Println("Numero de itens:", vquantosItens)
	// taking each item:
	cv1 := "V"
	pp2 := "KNIP11"
	qt3 := "10"
	un4 := "100.00"
	epc5 := "4.50"
	for i := 1; i <= vquantosItens; i++ {
		fmt.Println("===== cada item ===== entre cada item - ", i, " em ", vquantosItens)
		cv1 = "'C'ou'V'"
		pp2 = "PPPPPP"
		qt3 = "999"
		un4 = "999.99"
		epc5 = "4.50"
		fmt.Printf("%s %s %s %s %s\n", cv1, pp2, qt3, un4, epc5)
		n, err = fmt.Scanf("%s %s %s %s %s", &cv1, &pp2, &qt3, &un4, &epc5) // by reference!
		if err != nil {
			panic(err)
		}
		if n != 5 {
			panic("must read 5 variables")
		}
		cv1 = strings.ToUpper(cv1)
		if cv1 != "V" && cv1 != "C" {
			panic("Compra ou Venda deve ser 'C' ou 'V'")
		}
		linha.cvlnota = cv1
		linha.pplnota = pp2
		linha.qtlnota, erro = strconv.Atoi(qt3)
		if erro != nil {
			panic(erro)
		}
		linha.unlnota, erro = strconv.ParseFloat(un4, 64)
		if erro != nil {
			panic(erro)
		}
		linha.epcnota, erro = strconv.ParseFloat(epc5, 64)
		if erro != nil {
			panic(erro)
		}

		linha.txlnota = 0.
		linha.ttlnota = 0.
		nota = append(nota, linha)

		fmt.Println(linha)
	}
	fmt.Println(nota)
	fmt.Printf("%v\n", nota)
	//fmt.Printf("%q\n",nota)
	for ii := 0; ii < len(nota); ii++ {

		fmt.Printf(" %5v %5v %5v %5v %5v %5v \n",
			nota[ii].cvlnota,
			nota[ii].pplnota,
			nota[ii].qtlnota,
			nota[ii].unlnota,
			nota[ii].epcnota,
			nota[ii].txlnota,
		)

	}

}
