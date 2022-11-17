package hello

// Greet Greets GitHub Actions
func Greet() string {
	return "Hello GitHub Actions. Dev.to is awesome"
}

func Soma[T int64 | float32] (array []T) T {
	var soma T 
	for _, v := range array {
		soma = soma + v
	}
	return soma
}