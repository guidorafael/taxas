/*
***************************************************************

	"TAXAS" is an application to help with the Annual
	Income Tax Return (Declaração Anual de Imposto sobre Renda)
	and also for tracking investments.

	Calculation of fees and taxes on brokerage note
	with multiple purchases and sales items.
	(Multiple buying and sailling assets).

	In this case, a proportional division (weighted) of the total
	rate of the note must be made by each transacted asset.

	The individual fees are printed and then incorporated into
	the ".pdf" of the corresponding Brokerage Note
	(document which the Federal Revenue can request).

	In the accounting view of the algorithm:
	Purchase of assets represents debit ("-" sign).
	Fees and taxes represent debits ("-" sign)
	Sale of assets represents credit ("+" sign).
	Although they do not appear here, remember:
	dividends and earnings represent credits ("+" sign).
*/
package main

import (
	"fmt"
	"math"
	"strings"
)

var sversion = "v0.04"  // 03-03-2023  string version  -  v0.03 include column qt. X unitary in final result
type LinhaNota struct { // linha de uma nota de corretagem
	cvlnota string  // "C" ou "V" compra ou venda na linha         - compra (débito) sera negativo  venda (crédito) sera positivo
	pplnota string  // ativo ("PaPel") transacionado ns linha
	qtlnota int     // quantidade do ativo na linha                - herda o sinal de C/V "-" na compra e "+" na venda
	unlnota float64 // preço medio da aquisição (unitario da linha)- sempre positivo
	epcnota float64 // coRRetagem específica deste papel na linha  - SEMPRE NEGATIVA conhece o valor - fora do calculo ponderado
	txlnota float64 // taxa desta transac. deste ativo             - SEMPRE NEGATIVA a PRÓPRIA INCOGNITA a obter através de ponderação
	ttlnota float64 // custo deste ativo na linha                  - positivo (credito) na venda   negativo (débito) na compra
}

func main() {
	var nota []LinhaNota
	var linha LinhaNota
	var ttCompras, // soma modular de todas as compras
		ttVendas, // soma modular de todas as vendas
		tttaxas,
		vlLiqFinalNota, // valor final da Nota de Corretagem com sinal (Crédito "+" e Débito "-") !
		fatorTx, // (difAPond / vlBruto) - a multiplicar por cada "qtd. X unit." resulta a taxa ponderada da linha
		vlttepc, // sum of the asset-specific brokerage fees (out of the total to be weighted)
		sumAlgCV, // algebraic sum of purchases and sales ("profit or loss")
		sumModCV float64 // modular sum of purchases and sales ("traded amount")

	/************************************************************************************************************
	*                             total de nota só de compras é negativo
	*                   A T E N Ç Ã O === A T E N Ç Ã O === A T E N Ç Ã O === A T E N Ç Ã O ===
	*                            liquido para daqui a dois dias úteis COM SINAL ! ! ! ! ! ! ! ! !
	*
	*             CREDITO(+venda) ou DEBITO(-compra) líquido a ser lançado na conta de investimento -
	*
	*         vlLiqFinalNota é o ultimo valor da nota canto inferior direito - é o que entra ou sai da conta.
	*                                   trecho para carga do slice nota
	*
	*************************************************************************************************************/
	// ==> decimal point is "." <== +credito(venda)  -débito(compra)

	//cabec := "NECTON - Nota de Corretagem: teste 11 - Pregão: " + substr(time.Now().String(), 0, 16) + " Tot.:"
	//cabec := "NECTON - Nota de Corretagem: 4769400"+" - "+"Pregão:10/03/2022" + " Tot.:"
	cabec := "INTER - Nota de Corretagem: 20183469" + " - " + "Pregão: 20/03/2023" + " Tot.:"
	vlLiqFinalNota = 80.00 // colocar sinal + para credito(venda) e - para debito(compra)- decimal point is a "."

	// inicio de lançamento dos papeis:

	linha.cvlnota = "V"
	linha.pplnota = "KLBN11"
	linha.qtlnota = 5
	linha.unlnota = 9.00
	linha.epcnota = 0. // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "V"
	linha.pplnota = "SANB4"
	linha.qtlnota = 5
	linha.unlnota = 8.00
	linha.txlnota = 0.
	nota = append(nota, linha)
	/*
	   	linha.cvlnota = "V"
	   	linha.pplnota = "RECT11"
	   	linha.qtlnota = 20
	   	linha.unlnota = 50.2265
	   	linha.epcnota = 0            // "clearing" corretagem específica ações 4.50 NEC. ate 99 papeis
	   	linha.txlnota = 0.
	   	nota = append(nota, linha)

	   /*
	   	linha.cvlnota = "V"
	   	linha.pplnota = "HGRU11"
	   	linha.qtlnota = 9
	   	linha.unlnota = 126.75
	   	linha.epcnota = 0.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	   	linha.txlnota = 0.
	   	nota = append(nota, linha)

	   	linha.cvlnota = "V"
	   	linha.pplnota = "MXRF11"
	   	linha.qtlnota = 100
	   	linha.unlnota = 10.31
	   	linha.epcnota = 0.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	   	linha.txlnota = 0.
	   	nota = append(nota, linha)

	   	linha.cvlnota = "V"
	   	linha.pplnota = "MXRF11"
	   	linha.qtlnota = 200
	   	linha.unlnota = 10.30
	   	linha.epcnota = 0.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	   	linha.txlnota = 0.
	   	nota = append(nota, linha)

	   	linha.cvlnota = "V"
	   	linha.pplnota = "MXRF11"
	   	linha.qtlnota = 100
	   	linha.unlnota = 10.31
	   	linha.epcnota = 0.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	   	linha.txlnota = 0.
	   	nota = append(nota, linha)

	   	linha.cvlnota = "V"
	   	linha.pplnota = "TGAR11"
	   	linha.qtlnota = 8
	   	linha.unlnota = 124.37
	   	linha.epcnota = 0.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	   	linha.txlnota = 0.
	   	nota = append(nota, linha)
	*/
	/**/
	/**/

	// algoritmo:
	// soma algebrica de compras e vendas [qtd X vl.unit] = sumAlgCV
	// modulo de sumAlgCV menos modulo de vlLiqFinalNota (desconsidera clearing especifico de cada ativo)
	// é o valor das taxas a serem divididos proporcionalmente
	// fim do trecho para carga do slice nota ==================================================================

	sumModCV = 0 // sum   modulos compra e venda = "valor da operação"
	sumAlgCV = 0 // sum algebrica compra e venda = "liquido da operação sem corretagem e taxas"
	ttCompras = 0
	ttVendas = 0
	vlttepc = 0
	for i := range nota { // modo "simplificado" de for i , _ := range ...
		// normalizando os parametros:
		// .... nomes de papel em maiúsculo e também "C ou V" maiúsculo
		nota[i].pplnota = strings.ToUpper(nota[i].pplnota)
		nota[i].cvlnota = strings.ToUpper(nota[i].cvlnota)
		// .... corretagem específica é negativa sempre
		if nota[i].epcnota != 0 {
			nota[i].epcnota = -1 * nota[i].epcnota // corretagem específica é negativa se existir
		}
		// .... valor unitário sempre positivo
		if nota[i].unlnota < 0 {
			nota[i].unlnota = -1 * nota[i].unlnota // valor unitario é maior que zero
		}
		// .... quantidade leva o sinal algébrico de "C" "-" ou "V" "+"
		if nota[i].cvlnota == "V" {
			if nota[i].qtlnota < 0 {
				nota[i].qtlnota = -1 * nota[i].qtlnota // se venda a quantidade é maior que zero
			}
		} else if nota[i].cvlnota == "C" {
			if nota[i].qtlnota > 0 {
				nota[i].qtlnota = -1 * nota[i].qtlnota // se compra a quantidade é menor que zero
			}
		} else {
			//log.Fatal("****** ERRO !  Tem que ser 'c' ou 'v' ", nota[i].pplnota )
		}
		// fim da normalização de parametros

		// acumulação em ttCompras e ttVendas
		if strings.ToUpper(nota[i].cvlnota) == "C" {
			ttCompras += float64(nota[i].qtlnota) * nota[i].unlnota
		} else if strings.ToUpper(nota[i].cvlnota) == "V" {
			ttVendas += float64(nota[i].qtlnota) * nota[i].unlnota
		} else {
			//fmt.Println("valor não esperado para o campo C/V:", nota[i].pplnota, " - Adeus! -")
			panic("***** ERRO: valor não esperado para o\n             campo C/V: " + nota[i].cvlnota + " em " + nota[i].pplnota + " - Adeus! *****")
		} //                                           12345678901234

		// acumula a corretagem específica a substrair do total a ser ponderado para taxas normais
		vlttepc += nota[i].epcnota

		fmt.Println(nota[i], vlttepc)
	}
	fmt.Printf("\n")
	sumModCV = math.Abs(ttCompras) + math.Abs(ttVendas) // soma do módulo compra e venda = "total de operações" da N.C.
	sumAlgCV = ttCompras + ttVendas                     // sum algebrica compra e venda

	//tttaxas = math.Abs(vlLiqFinalNota) - math.Abs(sumAlgCV) - math.Abs(vlttepc)
	if sumAlgCV <= 0 {
		tttaxas = math.Abs(vlLiqFinalNota) - math.Abs(sumAlgCV) - math.Abs(vlttepc)
	} else { // sumAlgCV > 0
		tttaxas = math.Abs(vlLiqFinalNota) - math.Abs(sumAlgCV) + math.Abs(vlttepc)
	}
	fatorTx = math.Abs(tttaxas) / sumModCV

	fmt.Printf("   Compras: %9.2f     Vendas: %9.2f \nsum alg cv: %9.2f sum mod cv: %9.2f \n vlLiqFinalNota: %9.2f tttaxas:   %9.2f \nfatorTx: %9.6f vlttepc: %9.2f \n",
		ttCompras, ttVendas, sumAlgCV, sumModCV, vlLiqFinalNota, tttaxas, fatorTx, vlttepc)

	fmt.Println("========================= taxa de cada papel: ==============" + sversion + "=======") // a versão

	fmt.Printf("%s %1.2f \n", cabec, vlLiqFinalNota)

	fmt.Printf("\n%s", "lin ativo    C/V  qtd  unitario  qtd X unt especifica    taxa      preço")
	for i := range nota {
		nota[i].txlnota = -1 * fatorTx * math.Abs(float64(nota[i].qtlnota)) * nota[i].unlnota
		nota[i].ttlnota = fttlinnota((&nota[i]))
		fmt.Printf("\n%2d %7s %4s %5d %9.2f %9.2f %9.2f %9.2f %10.2f",
			(i + 1),
			nota[i].pplnota,
			nota[i].cvlnota,
			nota[i].qtlnota,
			nota[i].unlnota,
			float64(nota[i].qtlnota)*nota[i].unlnota,
			nota[i].epcnota,
			nota[i].txlnota,
			nota[i].ttlnota,
		)
	}
	fmt.Printf("\n")

	somataxas := 0.0
	sumttlnota := 0.0
	for _, umaLinha := range nota {
		somataxas += umaLinha.txlnota
		sumttlnota += umaLinha.ttlnota
		//fmt.Println(i, somataxas , umaLinha.txlnota)
	}
	fmt.Printf("   tot.taxas: %5.2f   sumtx: %5.2f   vl.liq.operac.: %9.2f\n", tttaxas, somataxas, sumttlnota)
	if math.Abs(sumttlnota-vlLiqFinalNota) > 0.005 {
		fmt.Println("* * * * * algo errado", vlLiqFinalNota, sumttlnota, " * * * * * ")
	}
}

func fttlinnota(aa *LinhaNota) float64 {
	resu := float64(aa.qtlnota)*(aa.unlnota) + aa.txlnota + aa.epcnota
	//fmt.Println(" ",resu)
	return (resu)
}

//	 NOTE: this isn't multi-Unicode-codepoint aware, like specifying skintone or
//		gender of an emoji: https://unicode.org/emoji/charts/full-emoji-modifiers.html
func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
