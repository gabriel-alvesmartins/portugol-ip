Algoritmo "Reajuste salarial"

Var
   sal,nvsal:real
Inicio
      escrevaL("Escreva o salário do funcionário: ")
      leia(sal)
      se (sal<=300) entao
         nvsal<- 1.5*sal
      senao
           nvsal<-1.3*sal
      fimse
      escreva("SALARIO COM REAJUSTE: ", nvsal:1:2)
Fimalgoritmo