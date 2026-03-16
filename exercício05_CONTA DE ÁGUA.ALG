Algoritmo "Conta de Água"

Var
   conta,consumo,valor: real
   tipo_cons: caractere

Inicio
      escreval("Qual é o número da sua conta? ")
      escreval("Qual o valor do seu consumo de água (em m³)? ")
      escreval("Qual o seu tipo de consumidor? RESIDENCIAL (R), COMERCIAL (C) ou INDUSTRIAL (I) ")
      leia(conta, consumo, tipo_cons)
      se (tipo_cons = "R") entao
         valor <- 5 + 0.05*consumo
      senao
           se (tipo_cons = "C") e (consumo <= 80) entao
              valor <- 500
           senao
                se (tipo_cons= "C") e (consumo>80) entao
                      valor <- 500 + 0.25*(consumo-80)
                senao
                     se (tipo_cons= "I") e (consumo<=100) entao
                        valor <- 800
                     senao
                          valor <- 800+0,04*(consumo-100)
                     fimse
                fimse
           fimse
      fimse
      escreval("Conta: ", conta)
      escreval("Valor da conta: ", valor)

Fimalgoritmo
