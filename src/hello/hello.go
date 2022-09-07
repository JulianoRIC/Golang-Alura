//primeiro programa em Go
//para compilar: go code.go
//executar: ./code
//ou: go run code.go

package main

import "fmt"
import  "os"

func main(){
    
   exibeIntroducao()
   exibeMenu()

   //lendo a opcao do usuario
   comando := leComando()
    
  //controle de fluxo com if-else
  /************************************ 
   if comando == 1{
	  fmt.Println("Monitorando...")
   }else if comando == 2{
      fmt.Println("Exibindo logs...")
   }else if comando == 0{
       fmt.Println("Saindo do programa")
   }else{
    fmt.Println("Yo no lo conozco senor")
   }
}
********************************************/

//controle de fluxo com switch
switch comando{
    case 1: 
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa..")
		os.Exit(0)
	default:
		fmt.Println("Yo no lo conozco senor")	
		os.Exit(-1)
}

}

//trabalhando com funcoes

func exibeIntroducao(){
    //nao eh preciso tipar a variavel, o Go infere sozinho	
    //outra forma de declarar variavel é  :=   
	var versao float32 = 1.1
	nome := "Juliano"
	fmt.Println("Hello Mister", nome)
	fmt.Println("Este programa está  na versão: ", versao)
}

func leComando() int{
	var comandoLido int
	//fmt.Scanf("%d", &comandoLido)
    //de uma forma melhor (sem o modificador)
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func exibeMenu(){
   fmt.Println("1-Iniciar monitoramento")
   fmt.Println("2-Exibir Logs")
   fmt.Println("0-Sair do programa")
}