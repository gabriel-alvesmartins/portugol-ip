Algoritmo "consumoDeEnergia"

Var
salario,energia,custo,porKw,desconto:real


Inicio
      escreva("Qual o valor do salário mínimo? ")
      leia (salario)
      escreva ("Quanto que é gasto de energia na sua casa(em kW)? ")
      leia (energia)
      // a cada 1kW, é gasto 0.7% do salário mínimo
      custo <- energia*0.007*salario
      porKw <- custo/energia
      desconto <- custo*0.9
      escrevaL("Custo por kW: R$", porKw:1:2)
      escrevaL("Custo do consumo: R$", custo:1:2)
      escrevaL("Custo com desconto: R$", desconto:1:2)
      



Fimalgoritmo