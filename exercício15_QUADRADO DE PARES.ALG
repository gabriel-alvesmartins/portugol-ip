Algoritmo "Quadrado de pares"

Var
   quadrado: vetor[1..2000] de real
   x,i,n: inteiro

Inicio
      escrevaL("Digite um valor N, tal que N seja maior que 5 e menor que 2000: ")
      leia(n)
      x <- 2
      para i de 1 ate n-1 faca
           se x%2=0 entao
              quadrado[n] <- x^2
              escrevaL(x,"^2 =", quadrado[n])
           fimse
      x<- i+2
      fimpara
Fimalgoritmo