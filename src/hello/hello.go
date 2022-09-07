//primeiro programa em Go
//para compilar: go code.go
//executar: ./code
//ou: go run code.go

package main

import "fmt"

func main(){
   //nao eh preciso tipar a variavel, o Go infere sozinho	
   var versao float32 = 1.1
   //outra forma de declarar variavel é  :=   
   nome := "Juliano"
   
   fmt.Println("Hello Mister", nome)
   fmt.Println("Este programa está  na versão: ", versao)
   
   fmt.Println("1-Iniciar monitorament")
   fmt.Println("2-Exibir Logs")
   fmt.Println("0-Sair do programa")

   //lendo a opcao do usuario
   var comando int
   //fmt.Scanf("%d", &comando)

   //de uma forma melhor (sem o modificador)
   fmt.Scan(&comando)

   fmt.Println("O comando escolhido foi", comando)
}