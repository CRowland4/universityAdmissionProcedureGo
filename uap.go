package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Applicant struct {
	firstName      string
	lastName       string
	pref1          string
	pref2          string
	pref3          string
	depAccepted    string
	physicsScore   float64
	chemistryScore float64
	mathScore      float64
	csScore        float64
}

func main() {
	departmentCapacity := readInt()
	applicants := getApplicants()

	admittedApplicants := executeAdmissionsProcess(applicants, departmentCapacity)
	printAdmitted(admittedApplicants)
	return
}

func executeAdmissionsProcess(apps []Applicant, departmentCapacity int) (admittedApplicants []Applicant) {
	departments := map[string]int{
		"Biotech":     0,
		"Chemistry":   0,
		"Engineering": 0,
		"Mathematics": 0,
		"Physics":     0,
	}

	apps = sortRound(apps, 1)
	for i, applicant := range apps {
		if (departments[applicant.pref1] < departmentCapacity) && (apps[i].depAccepted == "") {
			apps[i].depAccepted = applicant.pref1
			departments[applicant.pref1]++
			admittedApplicants = append(admittedApplicants, apps[i])
		}
	}

	apps = sortRound(apps, 2)
	for i, applicant := range apps {
		if (departments[applicant.pref2] < departmentCapacity) && (apps[i].depAccepted == "") {
			apps[i].depAccepted = applicant.pref2
			departments[applicant.pref2]++
			admittedApplicants = append(admittedApplicants, apps[i])
		}
	}

	apps = sortRound(apps, 3)
	for i, applicant := range apps {
		if (departments[applicant.pref3] < departmentCapacity) && (apps[i].depAccepted == "") {
			apps[i].depAccepted = applicant.pref3
			departments[applicant.pref3]++
			admittedApplicants = append(admittedApplicants, apps[i])
		}
	}

	return admittedApplicants
}

func sortRound(apps []Applicant, round int) (sortedApps []Applicant) {
	sort.Slice(apps, func(i, j int) bool {
		if examScore(apps[i], getPreference(apps[i], round)) != examScore(apps[j], getPreference(apps[j], round)) {
			return examScore(apps[i], getPreference(apps[i], round)) > examScore(apps[j], getPreference(apps[j], round))
		}
		if apps[i].firstName != apps[j].firstName {
			return apps[i].firstName < apps[j].firstName
		}

		return apps[i].lastName < apps[j].lastName
	})

	return apps
}

func getPreference(applicant Applicant, round int) (preference string) {
	switch round {
	case 1:
		return applicant.pref1
	case 2:
		return applicant.pref2
	default:
		return applicant.pref3
	}
}

func examScore(applicant Applicant, department string) (score float64) {
	switch department {
	case "Physics":
		return applicant.physicsScore
	case "Biotech", "Chemistry":
		return applicant.chemistryScore
	case "Mathematics":
		return applicant.mathScore
	case "Engineering":
		return applicant.csScore
	default:
		return 0.0
	}
}

func printAdmitted(apps []Applicant) {
	sort.Slice(apps, func(i, j int) bool {
		if apps[i].depAccepted != apps[j].depAccepted {
			return apps[i].depAccepted < apps[j].depAccepted
		}
		if examScore(apps[i], apps[i].depAccepted) != examScore(apps[j], apps[j].depAccepted) {
			return examScore(apps[i], apps[i].depAccepted) > examScore(apps[j], apps[j].depAccepted)
		}
		if apps[i].firstName != apps[j].firstName {
			return apps[i].firstName < apps[j].firstName
		}

		return apps[i].lastName < apps[j].lastName
	})

	for _, dep := range [5]string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"} {
		fmt.Println(dep)
		for _, applicant := range apps {
			if applicant.depAccepted == dep {
				fmt.Printf("%s %s %.1f\n",
					applicant.firstName,
					applicant.lastName,
					examScore(applicant, applicant.depAccepted))
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
		fmt.Sscanf(line, "%s %s %f %f %f %f %s %s %s",
			&newApplicant.firstName,
			&newApplicant.lastName,
			&newApplicant.physicsScore,
			&newApplicant.chemistryScore,
			&newApplicant.mathScore,
			&newApplicant.csScore,
			&newApplicant.pref1,
			&newApplicant.pref2,
			&newApplicant.pref3,
		)

		applicants = append(applicants, newApplicant)
	}

	return applicants
}

func readInt() (num int) {
	fmt.Scanln(&num)
	return num
}
