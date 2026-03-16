Algoritmo "maior"
Var
   n1,n2,n3:inteiro
Inicio
      escreval("Digite 3 números inteiros: ")
      leia(n1,n2,n3)
      se n1>n2 e n2>n3 e n1>n3 entao
        escreval(n1)
        senao
          se n1>n3 e n1>n2 e n3>n2 entao
            escreval(n1)
              senao
                se n2>n3 e n2>n1 e n3>n1 entao
                  escreval(n2)
                    senao
                     se n2>n1 e n2>n3 n1>n3 entao
                      escreval(n2)
                        senao
                          se n3>n1 e n3>n2 e n1>n2 entao
                            escreval(n3)
                              senao
                                se n3>n2 e n3>n1 e n2>n1 entao
                                  escreval(n3)
                                fimse
                          fimse
                     fimse
                 fimse
           fimse
      fimse  
Fimalgoritmo
