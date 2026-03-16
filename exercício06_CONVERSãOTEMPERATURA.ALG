Algoritmo "conversãoTemperatura"

Var
   tempF: vetor[1..10] de real
   tempC: vetor[1..10] de real
   n,i: inteiro

Inicio
      escreval("Quantas temperaturas você quer converter de Farenheit para Celcius? ")
      leia(n)
      para i de 1 ate n faca
           escreva("Qual o valor da", i, "ª temperatura? ")
           leia(tempF[i])
           tempC[i] <- 5*(tempF[i]-32)/9
      fimpara
      para i de 1 ate n faca
      escreval (tempF[i], "º FAHRENHEIT EQUIVALE A", tempC[i], "º CELSIUS")
      fimpara
           



Fimalgoritmo
