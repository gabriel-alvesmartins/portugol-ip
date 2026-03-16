Algoritmo "Somatório simples"

Var
   n,i:inteiro
   valor: vetor[1..999] de real
   soma: real


Inicio
      escreva("Digite um algarismo inteiro maior que 1: ")
      leia(n)
      soma<-0
      se n>1 entao
         para i de 1 ate n faca
              valor[i] <- 1/i
              soma<- soma+valor[i]
         fimpara
      senao
      escreva("Numero invalido!")
      fimse
      escreva(soma:1:6)
Fimalgoritmo
