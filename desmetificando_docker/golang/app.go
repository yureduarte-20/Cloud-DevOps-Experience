package main
import ("fmt")

func main(){
	fmt.Println("Qual seu nome?:")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Oi %s!, seja bem vindo", name)
}

