Algoritmo 
Var
   n1,n2,n3:inteiro
Inicio
      escreval("Digite 3 números inteiros: ")
      leia(n1,n2,n3)
      se n1>n2>n3 entao
        escreval(n1,n2,n3)
        senao
          se n1>n3>n2 entao
            escreval(n1,n3,n2)
              senao
                se n2>n3>n1 entao
                  escreval(n2,n3,n1)
                    senao
                     se n2>n1>n3 entao
                      escreval(n2,n1,n3)
                        senao
                          se n3>n1>n2 entao
                            escreval(n3,n1,n2)
                              senao
                                se n3>n2>n1 entao
                                  escreval(n3,n2,n1)
                                fimse
                          fimse
                     fimse
                 fimse
           fimse
      fimse  
Fimalgoritmo
