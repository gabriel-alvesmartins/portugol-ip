Algoritmo "mercadoria"
Var
  preco,pdesc,vmerc,desc: real

Inicio
      escreval("Digite o valor da mercadoria e o percentual de desconto para aplicar")
      leia(preco,pdesc)
      desc <- (pdesc/100)*preco
      vmerc <- (1-pdesc/100)*preco
      escreval("O valor do desconto é: ", desc," e o valor da mercadoria após o desconto é: ", vmerc)
Fimalgoritmo
