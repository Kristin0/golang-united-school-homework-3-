package homework

func Average(input []float32) (result float32) {

	var sum float32 = 0
	var num int = len(input)
	for i := 0; i < num; i++ {
		sum += input[i]
	}
	return sum/float32(num)
}
