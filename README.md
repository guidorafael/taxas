# taxas
# The problem and the purpose of the application:
A brokerage note ("nota de corretagem") corresponds to all transactions with assets carried out on a given date at a securities brokerage ("Corretora de Valores Mobiliarios").

It is important in the Federal Revenue ("Receita Federal") Income Tax return to discriminate for each asset what the rate was.
This fee is proportional to the amount transacted.
(Fees should not be considered in the collection of taxes).

# Application operation:
A slice "note" ("nota") of type "lineNote" ("linhaNota") describes the transactions carried out in a brokerage note. The expense called "fee" ("taxa") is calculated for each transaction by the weighted average between each portion of the amount transacted within the total value of the note.

Each transaction is written in source code (corresponds to a "linhaNota") added to the "nota" slice.

After the weighted average algorithm to calculate the fees for each transaction, STDOUT (screen) is used to display the results.



Important:
"TAXAS" is an application to help with the Annual Income Tax Return 
(Declaração Anual de Imposto sobre Renda) and also for tracking investments.

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