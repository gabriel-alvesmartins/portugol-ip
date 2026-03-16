Algoritmo "Conversão de Nota em Conceito"

Var
   n: real


Inicio
      escreval("Digite a nota do aluno: ")
      leia (n)
      se (n>=9) e (n<=10) entao
         escreva("Nota: ", n," Conceito: A")
      senao
           se (n >= 7.5) e (n<9) entao
              escreva("Nota: ", n," Conceito: B")
           senao
                se (n>=6) e (n < 7.5) entao
                   escreva("Nota: ", n," Conceito: C")
                senao
                     escreva("Nota: ", n," Conceito: D")
                fimse
           fimse
      fimse



Fimalgoritmo