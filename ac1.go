// package main

// import "fmt"

// func calcMedia(a, b float64) float64 {
// 	media := (a + b) / 2
// 	return media
// }

// func main() {
// 	// Declaracao das duas variaveis
// 	var a, b float64
// 	var media float64

// 	fmt.Println("Digite o primeiro valor:")
// 	fmt.Scanln(&a)
// 	fmt.Println("Digte o segundo valor:")
// 	fmt.Scanln(&b)

// 	//funcao para calcular media
// 	media = calcMedia(a, b)

// 	fmt.Println("A média é:", media)
// }

// package main

// import "fmt"

// func testePar(n int64) string {
// 	if n%2 == 0 {
// 		return "Par"
// 	}

// 	return "Impar"
// }

// func main() {
// 	var n int64

// 	fmt.Println("Digite o valor")
// 	fmt.Scanln(&n)

// 	answer := testePar(n)
// 	fmt.Println("Este numero e:", answer)

// }

// package main

// import "fmt"

// func calcExpoente(base int64, expoente int64) int64 { //o int64 fora do parenteses determinar qual tipo de resultado a função dara retorno
// 	// a é a variavel que esta sendo usada temporaritamente pela função para calcular o resultado
// 	a := base
// 	for i := 0; i < int(expoente-1); i++ {
// 		a *= base
// 	}

// 	return a
// }

// func main() {
// 	var base, expoente int64
// 	var resultado int64

// 	fmt.Println("Digite o valor da base:")
// 	fmt.Scanln(&base)
// 	fmt.Println("Digite o valor do expoente:")
// 	fmt.Scanln(&expoente)

// 	resultado = calcExpoente(base, expoente)

// 	fmt.Println("O resultado é:", resultado)

// }

// package main

// import "fmt"

// func converteCelsiusParaFahrenheit(temp float64) float64 {
// 	result := (temp * 1.8) + 32
// 	return result
// }

// func main() {
// 	var temp float64

// 	fmt.Println("Digite a temperatura em Celsius:")
// 	fmt.Scanln(&temp)

// 	resposta := converteCelsiusParaFahrenheit(temp)

// 	fmt.Println("Resultado da convercao", resposta)

// }

