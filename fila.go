package main

import "fmt"

type No struct {
	valor int
	prox  *No
}

type Fila struct {
	tamanho int
	inicio  *No
	fim     *No
}

func (f *Fila) percorre() {
	if f.inicio == nil {
		fmt.Println("Fila vazia")
	} else {
		no := f.inicio
		for no.prox != nil {
			fmt.Print(no.valor, " -> ")
			no = no.prox
		}
		fmt.Println(no.valor)
	}
}

func (f *Fila) insere(novoValor int) {
	novoNo := &No{valor: novoValor}

	if f.inicio == nil {
		f.inicio = novoNo
	} else {
		novoNo.prox = f.fim
		f.fim = novoNo
	}
	f.tamanho++
}

func (f *Fila) remove() error {
	if f.inicio == nil{
		return fmt.Errorf("Fila vazia")
	}

	if f.tamanho==1{
		f.inicio = nil
	}else{
		aux := f.inicio
		f.inicio = f.inicio.prox
		aux.prox = nil
	}

	f.tamanho--
	return nil

}

func main() {
	fila := Fila{}

	fila.insere(10)
	fmt.Println(fila.tamanho)
	fila.percorre()
	fmt.Println(fila.tamanho)
	fila.remove()
	fmt.Println(fila.tamanho)

}
