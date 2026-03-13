Algoritmo "Custo da Lata de Cerveja"

Var
     areaT,custo,r,h:real


Inicio
      escreval("Qual o raio e a altura da lata, respectivamente? (em metros)")
      leia(r,h)
      areaT <- 2*3.14159*r^2 + 2*3.14159*r*h
      custo <- areaT*100
      escreval("O VALOR DO CUSTO É: ", custo:1:2)
Fimalgoritmo