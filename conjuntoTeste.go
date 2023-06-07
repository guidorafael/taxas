	package conjuntoTeste
	/**************************************************************
	  casos de teste para o "taxas.go" 
	*/

	// ==> decimal point is "." <== +credito(venda)  -débito(compra)
	/*
	// *************************************
	// teste 01 tres compras sem corr.espec.
	// *************************************
	cabec := "NECTON - Nota de Corretagem: teste 01 - Pregão: " + substr(time.Now().String() , 0, 16) +" Tot.:"
	vlLiqFinalNota = -660.00 // colocar sinal + para credito(venda) e - para debito(compra)- decimal point is a "."

	// inicio lançamento dos papeis:
	linha.cvlnota = "c"
	linha.pplnota = "KNIP11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 0.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "c"
	linha.pplnota = "KNRI11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 0.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "c"
	linha.pplnota = "PPPP11"
	linha.qtlnota = 10
	linha.unlnota = 40.00
	linha.epcnota = 0 // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)
    */
	// ******************************************
	// teste 02 tres compras com duas corr.espec.
	// ******************************************
/*
	cabec := "NECTON - Nota de Corretagem: teste 02 - Pregão: " + substr(time.Now().String() , 0, 16) +" Tot.:"
	vlLiqFinalNota = -670.00 // colocar sinal + para credito(venda) e - para debito(compra)- decimal point is a "."

	// inicio lançamento dos papeis:
	linha.cvlnota = "c"
	linha.pplnota = "KNIP11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 3.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "c"
	linha.pplnota = "KNRI11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 7.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "c"
	linha.pplnota = "PPPP11"
	linha.qtlnota = 10
	linha.unlnota = 40.00
	linha.epcnota = 0 // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)
/*
	// *************************************
	// teste 03 tres vendas sem corr.espec.
	// *************************************
	cabec := "NECTON - Nota de Corretagem: teste 03 - Pregão: " + substr(time.Now().String() , 0, 16) +" Tot.:"
	vlLiqFinalNota = +540.00 // colocar sinal + para credito(venda) e - para debito(compra)- decimal point is a "."

	// inicio lançamento dos papeis:
	linha.cvlnota = "v"
	linha.pplnota = "KNIP11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 0.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "v"
	linha.pplnota = "KNRI11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 0.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "v"
	linha.pplnota = "PPPP11"
	linha.qtlnota = 10
	linha.unlnota = 40.00
	linha.epcnota = 0 // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)
*/
/*
	// *************************************
	// teste 04 tres vendas com corr.espec.    - encontrou erro - Este teste foi bem sucedido ! 530 <> 510 (subtraiu 2 a epcnota)
	// *************************************
	cabec := "NECTON - Nota de Corretagem: teste 04 - Pregão: " + substr(time.Now().String(), 0, 16) + " Tot.:"
	vlLiqFinalNota = +530.00 // colocar sinal + para credito(venda) e - para debito(compra)- decimal point is a "."
	// racional:
	//    +100 -10 -10    +100 -10 -0    +400 -40 -0

	// inicio lançamento dos papeis:
	linha.cvlnota = "v"
	linha.pplnota = "KNIP11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 10. // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "v"
	linha.pplnota = "KNRI11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 0. // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "v"
	linha.pplnota = "PPPP11"
	linha.qtlnota = 10
	linha.unlnota = 40.00
	linha.epcnota = 0 // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)
    
	// RESULTADO

└$ go run taxas.go
{KNIP11 V 10 10 -10 0 0} -10
{KNRI11 V 10 10 0 0 0} -10
{PPPP11 V 10 40 0 0 0} -10

   Compras:      0.00     Vendas:    600.00 
sum alg cv:    600.00 sum mod cv:    600.00 
 vlLiqFinalNota:    530.00 tttaxas:      -60.00 
fatorTx:  0.100000 vlttepc:    -10.00 
=================== taxa de cada papel: ==========v0.03====
NECTON - Nota de Corretagem: teste 04 - Pregão: 2023-03-03 11:54 Tot.: 530.00 

lin ativo    C/V  qtd  unitario  qtd X unt especifica    taxa      preço
 1  KNIP11    V    10     10.00    100.00    -10.00    -10.00      80.00
 2  KNRI11    V    10     10.00    100.00      0.00    -10.00      90.00
 3  PPPP11    V    10     40.00    400.00      0.00    -40.00     360.00
   tot.taxas: -60.00   sumtx: -60.00   vl.liq.operac.:    530.00


*/
/*	// *************************************
	// teste 05 tres compras com corr.espec.  não deu erro !
	// *************************************  ############## 
	cabec := "NECTON - Nota de Corretagem: teste 05 - Pregão: " + substr(time.Now().String() , 0, 16) +" Tot.:"
	vlLiqFinalNota = -670.00 // colocar sinal + para credito(venda) e - para debito(compra)- decimal point is a "."
    // -616 = -100-10-10  -100-10  -400-40  (cada 100 cobra 10 de fee e uma cobra 10 corretagem específica)  
	// inicio lançamento dos papeis:
	linha.cvlnota = "c"
	linha.pplnota = "KNIP11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 10.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "c"
	linha.pplnota = "KNRI11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 0.
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "c"
	linha.pplnota = "PPPP11"
	linha.qtlnota = 10
	linha.unlnota = 40.00
	linha.epcnota = 0 
	linha.txlnota = 0.
	nota = append(nota, linha)
    
	// RESULTADO: 
└$ go run taxas.go
{KNIP11 C -10 10 -10 0 0} -10
{KNRI11 C -10 10 0 0 0} -10
{PPPP11 C -10 40 0 0 0} -10

   Compras:   -600.00     Vendas:      0.00 
sum alg cv:   -600.00 sum mod cv:    600.00 
 vlLiqFinalNota:   -670.00 tttaxas:       60.00 
fatorTx:  0.100000 vlttepc:    -10.00 
=================== taxa de cada papel: ==========v0.01====
NECTON - Nota de Corretagem: teste 05 - Pregão: 2023-03-03 11:31 Tot.: -670.00 

lin ativo    C/V  qtd  unitario  especifica    taxa      preço
 1  KNIP11    C   -10     10.00    -10.00    -10.00    -120.00
 2  KNRI11    C   -10     10.00      0.00    -10.00    -110.00
 3  PPPP11    C   -10     40.00      0.00    -40.00    -440.00
   tot.taxas: 60.00   sumtx: -60.00   vl.liq.operac.:   -670.00



*/
/*
	// ***********************************************
	// teste 06 duas compras uma venda sem corr.espec.   sem erro ! 
	// ***********************************************   ##########
	cabec := "NECTON - Nota de Corretagem: teste 06 - Pregão: " + substr(time.Now().String() , 0, 16) +" Tot.:"
	vlLiqFinalNota = -460.00 // colocar sinal + para credito(venda) e - para debito(compra)- decimal point is a "."
    // -460 = v+100-10  c-100-10  c-400-40  0 + 0 + 0 sem específica   
	// inicio lançamento dos papeis:
	linha.cvlnota = "v"
	linha.pplnota = "KNIP11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 0.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "c"
	linha.pplnota = "KNRI11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 0.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "c"
	linha.pplnota = "PPPP11"
	linha.qtlnota = 10
	linha.unlnota = 40.00
	linha.epcnota = 0 // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)
*/

/*
	// ***********************************************
	// teste 07 venda com corr.espec. uma venda sem e uma compra sem  -  sem erro ! 
	// ***********************************************                   ##########
	cabec := "NECTON - Nota de Corretagem: teste 07 - Pregão: " + substr(time.Now().String() , 0, 16) +" Tot.:"
	vlLiqFinalNota = -262.00 // colocar sinal + para credito(venda) e - para debito(compra)- decimal point is a "."
    // -460 = v+100-10  v+100-10  c-400-40  v-2 + 0 + 0 sem específica   
	// inicio lançamento dos papeis:
	linha.cvlnota = "v"
	linha.pplnota = "KNIP11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 2.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "v"           
	linha.pplnota = "KNRI11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 0.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "c"
	linha.pplnota = "PPPP11"
	linha.qtlnota = 10
	linha.unlnota = 40.00
	linha.epcnota = 0 // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)
*/
/*
	// ***********************************************
	// teste 08 venda com corr.espec. duas vendas uma com -  erro ! 
	// ***********************************************       ######
		cabec := "NECTON - Nota de Corretagem: teste 08 - Pregão: " + substr(time.Now().String(), 0, 16) + " Tot.:"
		vlLiqFinalNota = 178.00 // colocar sinal + para credito(venda) e - para debito(compra)- decimal point is a "."
	         // RACIONAL
		     // +178,00 = v+100-10-2 com especifica  v+100-10  sem especifica  
		
		// inicio lançamento dos papeis:
	   	linha.cvlnota = "v"
	   	linha.pplnota = "KNIP11"
	   	linha.qtlnota = 10
	   	linha.unlnota = 10.00
	   	linha.epcnota = 2.       // corretagem específica ações 4.50 NEC. ate 100 papeis
	   	linha.txlnota = 0.
	   	nota = append(nota, linha)

	   	linha.cvlnota = "v"
	   	linha.pplnota = "KNRI11"
	   	linha.qtlnota = 10
	   	linha.unlnota = 10.00
	   	linha.epcnota = 0.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	   	linha.txlnota = 0.
	   	nota = append(nota, linha)
//      RESULTADO
└$ go run taxas.go
{KNIP11 V 10 10 -2 0 0} -2
{KNRI11 V 10 10 0 0 0} -2

   Compras:      0.00     Vendas:    200.00 
sum alg cv:    200.00 sum mod cv:    200.00 
 vlLiqFinalNota:    178.00 tttaxas:      -20.00 
fatorTx:  0.100000 vlttepc:     -2.00 
=================== taxa de cada papel: ==========v0.03====
NECTON - Nota de Corretagem: teste 08 - Pregão: 2023-03-03 12:05 Tot.: 178.00 

lin ativo    C/V  qtd  unitario  qtd X unt especifica    taxa      preço
 1  KNIP11    V    10     10.00    100.00     -2.00    -10.00      88.00
 2  KNRI11    V    10     10.00    100.00      0.00    -10.00      90.00
   tot.taxas: -20.00   sumtx: -20.00   vl.liq.operac.:    178.00
*/

/*
	// ***********************************************
	// teste 09 - uma venda com corr.espec. 
	// ***********************************************                   ##########
	vlLiqFinalNota = -24.00 // colocar sinal + para credito(venda) e - para debito(compra)- decimal point is a "."
    // +88,00 = v+100-10  0v+100-10  v-2 -0 sem específica   
	// inicio lançamento dos papeis:
	linha.cvlnota = "v"
	linha.pplnota = "KNIP11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 2.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)

	linha.cvlnota = "c"
	linha.pplnota = "KNRI11"
	linha.qtlnota = 10
	linha.unlnota = 10.00
	linha.epcnota = 2.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	linha.txlnota = 0.
	nota = append(nota, linha)
    
	// resultado: 
	└$ go run taxas.go
	{KNRI11 C -10 10 -2 0 0} -4
	{KNIP11 V 10 10 -2 0 0} -2

	sum alg cv:      0.00 sum mod cv:    200.00 
	Compras:   -100.00     Vendas:    100.00 
	fatorTx:  0.100000 vlttepc:     -4.00 
	=================== taxa de cada papel: ==========v0.01====
	NECTON - Numero da Nota e Data do Pregão: teste3182967 teste09/02/2023     -24.00 
	
	lin ativo    C/V  qtd  unitario  especifica    taxa      preço
	1  KNIP11    V    10     10.00     -2.00    -10.00      88.00
	2  KNRI11    C   -10     10.00     -2.00    -10.00    -112.00
	tot.taxas: 20.00   sumtx: -20.00   vl.liq.operac.:    -24.00
	vlLiqFinalNota:    -24.00 tttaxas:       20.00 


	   	// *************************************
	   	// teste 11 tres vendas com corr.espec.
	   	// *************************************
        // RACIONAL:  tres vendas  -10 -10 -40 taxa proporcional  -1 -2 taxa específica 
		   //  v +100 -10 -1  v +100 -10 -0    v +400 -40 -2    
		cabec := "NECTON - Nota de Corretagem: teste 11 - Pregão: " + substr(time.Now().String(), 0, 16) + " Tot.:"
	   	vlLiqFinalNota = +537.00 // colocar sinal + para credito(venda) e - para debito(compra)- decimal point is a "."

	   	// inicio lançamento dos papeis:
	   	linha.cvlnota = "v"
	   	linha.pplnota = "KNIP11"
	   	linha.qtlnota = 10
	   	linha.unlnota = 10.00
	   	linha.epcnota = 1.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	   	linha.txlnota = 0.
	   	nota = append(nota, linha)

	   	linha.cvlnota = "v"
	   	linha.pplnota = "KNRI11"
	   	linha.qtlnota = 10
	   	linha.unlnota = 10.00
	   	linha.epcnota = 0.           // corretagem específica ações 4.50 NEC. ate 99 papeis
	   	linha.txlnota = 0.
	   	nota = append(nota, linha)

	   	linha.cvlnota = "v"
	   	linha.pplnota = "PPPP11"
	   	linha.qtlnota = 10
	   	linha.unlnota = 40.00
	   	linha.epcnota = 2 // corretagem específica ações 4.50 NEC. ate 99 papeis
	   	linha.txlnota = 0.
	   	nota = append(nota, linha)	
// RESULTADO
└$ go run taxas.go
{KNIP11 V 10 10 -1 0 0} -1
{KNRI11 V 10 10 0 0 0} -1
{PPPP11 V 10 40 -2 0 0} -3

   Compras:      0.00     Vendas:    600.00 
sum alg cv:    600.00 sum mod cv:    600.00 
 vlLiqFinalNota:    537.00 tttaxas:      -60.00 
fatorTx:  0.100000 vlttepc:     -3.00 
========================= taxa de cada papel: ==============v0.03=======
NECTON - Nota de Corretagem: teste 11 - Pregão: 2023-03-03 12:21 Tot.: 537.00 

lin ativo    C/V  qtd  unitario  qtd X unt especifica    taxa      preço
 1  KNIP11    V    10     10.00    100.00     -1.00    -10.00      89.00
 2  KNRI11    V    10     10.00    100.00      0.00    -10.00      90.00
 3  PPPP11    V    10     40.00    400.00     -2.00    -40.00     358.00
   tot.taxas: -60.00   sumtx: -60.00   vl.liq.operac.:    537.00

*/