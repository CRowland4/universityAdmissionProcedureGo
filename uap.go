package main

import "fmt"

func main() {
	exam1, exam2, exam3 := readExamScores()
	fmt.Println(calculateAverage(exam1, exam2, exam3))
	fmt.Println("Congratulations, you are accepted!")

	return
}

func calculateAverage(scores ...float64) (averageScore float64) {
	for _, score := range scores {
		averageScore += score
	}

	averageScore /= float64(len(scores))
	return averageScore
}

func readExamScores() (score1, score2, score3 float64) {
	fmt.Scan(&score1, &score2, &score3)
	return score1, score2, score3
}
