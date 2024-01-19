package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Applicant struct {
	firstName   string
	lastName    string
	gpa         float64
	choice1     string
	choice2     string
	choice3     string
	depAccepted string
}

func main() {
	applicants := getApplicants()
	departmentCapacity := readInt()

	admittedApplicants := executeAdmissionsProcess(&applicants, departmentCapacity)
	printAdmitted(admittedApplicants)
	return
}

func executeAdmissionsProcess(applicants *[]Applicant, departmentCapacity int) (admittedApplicants []Applicant) {
	departments := map[string]int{
		"Physics":     0,
		"Mathematics": 0,
		"Engineering": 0,
		"Chemistry":   0,
		"Biotech":     0,
	}

	// Round 1
	for i, applicant := range *applicants {
		if (departments[applicant.choice1] < departmentCapacity) && ((*applicants)[i].depAccepted == "") {
			(*applicants)[i].depAccepted = applicant.choice1
			departments[applicant.choice1]++
			admittedApplicants = append(admittedApplicants, (*applicants)[i])
		}
	}

	// Round 2
	for i, applicant := range *applicants {
		if (departments[applicant.choice2] < departmentCapacity) && ((*applicants)[i].depAccepted == "") {
			(*applicants)[i].depAccepted = applicant.choice2
			departments[applicant.choice2]++
			admittedApplicants = append(admittedApplicants, (*applicants)[i])
		}
	}

	// Round 3
	for i, applicant := range *applicants {
		if (departments[applicant.choice3] < departmentCapacity) && ((*applicants)[i].depAccepted == "") {
			(*applicants)[i].depAccepted = applicant.choice3
			departments[applicant.choice3]++
			admittedApplicants = append(admittedApplicants, (*applicants)[i])
		}
	}

	return admittedApplicants
}

func printAdmitted(applicants []Applicant) {
	sort.Slice(applicants, func(i, j int) bool {
		if applicants[i].depAccepted != applicants[j].depAccepted {
			return applicants[i].depAccepted < applicants[j].depAccepted
		}
		if applicants[i].gpa != applicants[j].gpa {
			return applicants[i].gpa > applicants[j].gpa
		}
		if applicants[i].firstName != applicants[j].firstName {
			return applicants[i].firstName < applicants[j].firstName
		}

		return applicants[i].lastName < applicants[j].lastName
	})

	for _, dep := range [5]string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"} {
		fmt.Println(dep)
		for _, applicant := range applicants {
			if applicant.depAccepted == dep {
				fmt.Printf("%s %s %.2f\n", applicant.firstName, applicant.lastName, applicant.gpa)
			}
		}
		fmt.Println()
	}

	return
}

func getApplicants() (applicants []Applicant) {
	file, _ := os.Open("applicants.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var newApplicant Applicant
		line := scanner.Text()
		fmt.Sscanf(line, "%s %s %f %s %s %s",
			&newApplicant.firstName,
			&newApplicant.lastName,
			&newApplicant.gpa,
			&newApplicant.choice1,
			&newApplicant.choice2,
			&newApplicant.choice3)

		applicants = append(applicants, newApplicant)
	}

	sort.Slice(applicants, func(i, j int) bool {
		if applicants[i].gpa != applicants[j].gpa {
			return applicants[i].gpa > applicants[j].gpa
		}

		return (applicants[i].firstName + applicants[i].lastName) < (applicants[j].firstName + applicants[j].lastName)
	})

	return applicants
}

func readInt() (num int) {
	fmt.Scanln(&num)
	return num
}
