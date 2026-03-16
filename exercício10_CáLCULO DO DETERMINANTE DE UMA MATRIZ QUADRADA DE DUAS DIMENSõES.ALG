Algoritmo "Cálculo do Determinante de uma Matriz Quadrada de Duas Dimensões"

Var
   matr: vetor[1..2,1..2] de real
   i,j: inteiro
   determinante: real
Inicio
      Escreval("Este é um calculador de matriz 2x2")
      para i de 1 ate 2 faca
           para j de 1 ate 2 faca
                Escreval("Digite o ", j,"º número da matriz da linha",i)
                leia(matr[i,j])
           fimpara
      fimpara
      determinante<- matr[1,1]*matr[2,2] - matr[1,2]*matr[2,1]
      escreva("O VALOR DO DETERMINANTE É: ",determinante)
Fimalgoritmo