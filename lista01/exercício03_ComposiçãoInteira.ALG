Algoritmo "algarismos"

Var
n1,n2,n3,concN:inteiro
Inicio
      escrevaL("Digite os 3 algorismos do número desejado, separados por ENTER: ")
      leia(n1,n2,n3)
      se (n1>=10) ou (n2>=10) ou (n3>=10) entao
         escrevaL ("DIGITO INVALIDO")
      senao
         se n1<>0 entao
            concN <- n1*100 + n2*10 + n3
         senao
               se n2<>0 entao
                  concN <- n2*10 + n3
               senao
                    concN <- n3
                    se n3 = 0 entao
                       concN <- 0
                    fimse
               fimse
         fimse
      escrevaL ("O seu número é: ", concN)
      escreva ("O seu número ao quadrado é: ", concN^2)
      fimse
Fimalgoritmo
