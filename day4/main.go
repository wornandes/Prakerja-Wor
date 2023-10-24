package main

import "fmt"

// type Car struct {
// 	// Properti type dan fuelIn sesuai dengan spesifikasi
// 	Type   string
// 	FuelIn float64
// }

// // Method untuk menghitung perkiraan jarak yang bisa ditempuh
// func (car *Car) CalculateEstimatedDistance() float64 {
// 	// 1.5 L / mill adalah konversi dari 1 L / km
// 	return car.FuelIn / 1.5
// }

// func main() {
// 	myCar := Car{
// 		Type:   "Sedan",
// 		FuelIn: 60.0, // Misalnya, terdapat 60 liter bahan bakar
// 	}

// 	estimatedDistance := myCar.CalculateEstimatedDistance()

// 	fmt.Printf("Mobil %s bisa menempuh perkiraan jarak %.2f km\n", myCar.Type, estimatedDistance)
// }

// Function untuk mencari nilai maksimum dan minimum
// func findMaxMin(numbers []int) (max, min int) {
// 	max = numbers[0]
// 	min = numbers[0]

// 	for _, num := range numbers {
// 		if num > max {
// 			max = num
// 		}
// 		if num < min {
// 			min = num
// 		}
// 	}

// 	return max, min
// }

// func main() {
// 	var numbers [6]int

// 	fmt.Println("Masukkan 6 angka:")
// 	for i := 0; i < 6; i++ {
// 		fmt.Scan(&numbers[i])
// 	}

// 	max, min := findMaxMin(numbers[:])

// 	fmt.Printf("Nilai maksimum: %d\n", max)
// 	fmt.Printf("Nilai minimum: %d\n", min)
// }

type Student struct {
	Name  string
	Score int
}

type Students []Student

func (s Students) CalculateAverage() float64 {
	total := 0
	for _, student := range s {
		total += student.Score
	}
	return float64(total) / float64(len(s))
}

func (s Students) FindMin() Student {
	minStudent := s[0]
	for _, student := range s {
		if student.Score < minStudent.Score {
			minStudent = student
		}
	}
	return minStudent
}

func (s Students) FindMax() Student {
	maxStudent := s[0]
	for _, student := range s {
		if student.Score > maxStudent.Score {
			maxStudent = student
		}
	}
	return maxStudent
}

func main() {
	var students Students

	for i := 0; i < 5; i++ {
		var name string
		var score int

		fmt.Printf("Input %d Student's Name: ", i+1)
		fmt.Scan(&name)

		fmt.Printf("Input %d Student's Score: ", i+1)
		fmt.Scan(&score)

		students = append(students, Student{Name: name, Score: score})
	}

	average := students.CalculateAverage()
	minStudent := students.FindMin()
	maxStudent := students.FindMax()

	fmt.Printf("\nAverage Score: %.2f\n", average)
	fmt.Printf("Min Score of Students: %s (%d)\n", minStudent.Name, minStudent.Score)
	fmt.Printf("Max Score of Students: %s (%d)\n", maxStudent.Name, maxStudent.Score)
}
