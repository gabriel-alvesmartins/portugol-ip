Algoritmo "Volume da Pirâmide de Base Hexagonal"

Var
   h,a,volume:real

Inicio
      escreval("Qual o valor da altura e da aresta da pirâmide, respectivamente? (em metros) ")
      leia(h,a)
      volume <- 3*a^2*RaizQ(3)*h/6
      escreva("O VOLUME DA PIRAMIDE É: ", volume:1:2," METROS CUBICOS")
      
Fimalgoritmo
