package main

import "fmt"

func main() {
    // Criando um vetor de inteiros com 10 elementos
    vetor := make([]int, 10)

    // Inicializando o vetor com números quaisquer
    vetor[0] = 10
    vetor[1] = 20
    vetor[2] = 30
    vetor[3] = 40
    vetor[4] = 50
    vetor[5] = 60
    vetor[6] = 70
    vetor[7] = 80
    vetor[8] = 90
    vetor[9] = 100

    // Imprimindo os elementos do vetor um abaixo do outro
    for _, elemento := range vetor {
        fmt.Println(elemento)
    }
}


package main

import (
    "fmt"
)

// Função para inverter uma string
func inverterString(input string) string {
    // Converte a string para um slice de runes
    caracteres := []rune(input)

    // Inverte o slice de runes
    for i, j := 0, len(caracteres)-1; i < j; i, j = i+1, j-1 {
        caracteres[i], caracteres[j] = caracteres[j], caracteres[i]
    }

    // Converte o slice de runes de volta para uma string e retorna
    return string(caracteres)
}

func main() {
    // Solicita ao usuário que digite uma string
    fmt.Print("Digite uma string: ")
    var entrada string
    fmt.Scanln(&entrada)

    // Chama a função para inverter a string e imprime o resultado
    resultado := inverterString(entrada)
    fmt.Println("String invertida:", resultado)
}


package main

import (
    "fmt"
    "math"
)

// Definição do tipo Ponto
type Ponto struct {
    X float64
    Y float64
}

// Método para calcular a distância do ponto até a origem
func (p *Ponto) DistanciaOrigem() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
    // Criando um ponto
    ponto := Ponto{3, 4}

    // Calculando a distância até a origem
    distancia := ponto.DistanciaOrigem()

    // Imprimindo a distância
    fmt.Printf("A distância do ponto (%.2f, %.2f) até a origem é %.2f\n", ponto.X, ponto.Y, distancia)
}


package geometria

// Função para calcular a área do retângulo
func AreaRetangulo(base, altura float64) float64 {
    return base * altura
}

// Função para calcular o perímetro do retângulo
func PerimetroRetangulo(base, altura float64) float64 {
    return 2 * (base + altura)
}

package main

import (
    "fmt"
    "geometria"
)

func main() {
    // Solicita ao usuário que forneça as dimensões do retângulo
    var base, altura float64
    fmt.Println("Digite a base do retângulo:")
    fmt.Scanln(&base)
    fmt.Println("Digite a altura do retângulo:")
    fmt.Scanln(&altura)

    // Calcula a área e o perímetro do retângulo usando o pacote "geometria"
    area := geometria.AreaRetangulo(base, altura)
    perimetro := geometria.PerimetroRetangulo(base, altura)

    // Imprime os resultados
    fmt.Printf("Área do retângulo: %.2f\n", area)
    fmt.Printf("Perímetro do retângulo: %.2f\n", perimetro)
}

