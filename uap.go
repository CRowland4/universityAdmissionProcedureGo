package main

import (
	"fmt"
	"sort"
)

type Applicant struct {
	firstName string
	lastName  string
	gpa       float32
}

func main() {
	nApplicants := readInt()
	capacity := readInt()

	applicants := getApplicants(nApplicants)
	sort.Slice(applicants, func(i, j int) bool {
		if applicants[i].gpa != applicants[j].gpa {
			return applicants[i].gpa > applicants[j].gpa
		}

		return (applicants[i].firstName + applicants[i].lastName) < (applicants[j].firstName + applicants[j].lastName)
	})

	printAccepted(applicants, capacity)
	return
}

func printAccepted(applicants []Applicant, capacity int) {
	fmt.Println("Successful applicants:")
	for i := 0; i < capacity; i++ {
		fmt.Println(applicants[i].firstName, applicants[i].lastName)

	}

	return
}

func getApplicants(nApplicants int) (applicants []Applicant) {
	for i := 0; i < nApplicants; i++ {
		var applicant Applicant
		fmt.Scan(&applicant.firstName, &applicant.lastName, &applicant.gpa)
		applicants = append(applicants, applicant)
	}

	return applicants
}

func readInt() (num int) {
	fmt.Scanln(&num)
	return num
}
