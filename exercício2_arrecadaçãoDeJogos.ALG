Algoritmo "arrecadaçãoDeJogos"


Var
   i,nj,total,pop,ge,arq,cad: inteiro
   arrecadacao: vetor[1..10000000000000000000000000] de real



Inicio
      escreva("Quantos jogos você quer analisar? ")
      leia(nj)
      para i de 1 ate nj faca
           escrevaL("Responda os seguintes questionamentos, digitando linha por linha e pressionando enter para cada resposta" )
           escrevaL("Qual a quantidade TOTAL de pessoas que compraram igresso para o JOGO? ")
           escrevaL("Qual a porcentagem de pesssoas que compraram ingresso na categoria POPULAR? ")
           escrevaL("Qual a porcentagem de pesssoas que compraram ingresso na categoria GERAL? ")
           escrevaL("Qual a porcentagem de pesssoas que compraram ingresso na categoria ARQUIBANCADA? ")
           escrevaL("Qual a porcentagem de pesssoas que compraram ingresso na categoria CADEIRAS? ")
           leia(total,pop,ge,arq,cad)
           arrecadacao[i] <- ((total*pop/100)*1 + (total*ge/100)*5 + (total*arq/100)*10 + (total*cad/100)*20)
      fimpara
      para i de 1 ate nj faca
      escrevaL("A RENDA DO JOGO Nº", i," É", arrecadacao[i])
      fimpara
Fimalgoritmo