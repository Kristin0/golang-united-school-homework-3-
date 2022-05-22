package homework



func Average(input [15]float32) (result float32) {

    var count int
	var sum float32 = 0
	var num int = len(input)
	for i := 0; i < num; i++ {
		if input[i] != 0 {
			sum += input[i]
			count++
		}
	}
	return sum/float32(count)
}
