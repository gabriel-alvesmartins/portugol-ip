Algoritmo "médiaAluno"
Var
   n1,n2,n3,media: real
Inicio
      escreval("Quais foram as notas do aluno? ")
      leia(n1, n2, n3)
      media <- (n1+n2+n3)/3
      se (media>=6) entao
         escrevaL("Média = ", media)
         escrevaL("APROVADO")
      senao
           escrevaL("Média = ", media)
           escrevaL("REPROVADO")
      fimse
Fimalgoritmo