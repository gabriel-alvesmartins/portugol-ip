Algoritmo "Soma de progressão aritmética"

Var
    soma,a1,r,n,an:real
Inicio
      escreval("Digite, respectivamente, o valor do primeiro termo, a razão e o número de elementos da progressão: ")
      leia(a1,r,n)
      an<- a1+(n-1)*r
      soma<- (a1+an)*n/2
      escreval (soma)
Fimalgoritmo