package main

import (
	"fmt"
	"strings"
)

func main() {

	var nome string
	fmt.Print("Informe seu nome: ")
	fmt.Scanln(&nome)
	fmt.Println(nome)

	var email string
	fmt.Print("Informe seu email: ")
	fmt.Scanln(&email)
	fmt.Println(email)

}

type Contato struct {
	nome  string
	email string
}

func (c *Contato) alterarEmail(e string) {
	c.email = e
}

func adicionarContato(c *Contato,contatos [5]Contato){
	for i := 0; i < len(contatos); i++ {
		if (contatos[i] == Contato{}){
			contatos[i] = *c
		}
	}

}