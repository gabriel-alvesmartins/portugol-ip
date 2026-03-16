Algoritmo "Locadora de charretes"

Var
   valor: real
   horas: inteiro
Inicio
      escreva("Quantas horas o cliente usou a charrete? ")
      leia(horas)
      // a cada 3 horas, o valor pago é de 10 reais. Para horas avulsas o valor pago é de 5 reais.
      valor<- ((horas-horas%3)/3)*10 + (horas%3)*5
      escreval("O VALOR A PAGAR É: ", valor:1:2)
Fimalgoritmo