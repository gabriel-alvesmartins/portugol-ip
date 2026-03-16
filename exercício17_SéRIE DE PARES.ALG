Algoritmo "Série de pares"

Var
   n,q,i: inteiro
   valor: vetor[0..999] de inteiro


Inicio
      escrevaL("Digite um número PAR e em sequência a QUANTIDADE de números que pares que o sucedam: ")
      leia(n,q)
      se (n%2=0) entao
         para i de 0 ate (q*2)-2 faca
              se i%2=0 entao
                 valor[i] <- n+i
                 escreval(valor[i])
              fimse
         fimpara
      senao
          escreva("O PRIMEIRO NÚMERO NÃO É PAR")
      fimse




Fimalgoritmo