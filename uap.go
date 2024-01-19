package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Applicant struct {
	firstName, lastName, pref1, pref2, pref3, acceptedTo string
	physicsScore, chemistryScore, mathScore, csScore     float64
}

func main() {
	departmentCapacity := readInt()
	applicants := getApplicants()
	admittedApplicants := executeAdmissionsProcess(applicants, departmentCapacity)
	printAdmitted(admittedApplicants)
	return
}

func executeAdmissionsProcess(apps []Applicant, max int) (admittedApplicants []Applicant) {
	deps := map[string]int{"Biotech": 0, "Chemistry": 0, "Engineering": 0, "Mathematics": 0, "Physics": 0, "max": max}
	apps, deps = doAdmissionRound(apps, deps, 1)
	apps, deps = doAdmissionRound(apps, deps, 2)
	apps, deps = doAdmissionRound(apps, deps, 3)
	return apps
}

func doAdmissionRound(apps []Applicant, deps map[string]int, round int) (appsNew []Applicant, depsNew map[string]int) {
	apps = sortRound(apps, round)
	for i, app := range apps {
		if (deps[getPreference(app, round)] < deps["max"]) && (apps[i].acceptedTo == "") {
			apps[i].acceptedTo = getPreference(app, round)
			deps[getPreference(app, round)]++
		}
	}

	return apps, deps
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
	case "Mathematics":
		return applicant.mathScore
	case "Engineering":
		return applicant.csScore
	default:
		return applicant.chemistryScore
	}
}

func printAdmitted(apps []Applicant) {
	apps = sortAdmittedApplicants(apps)
	for _, dep := range [5]string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"} {
		fmt.Println(dep)
		for _, app := range apps {
			if app.acceptedTo == dep {
				fmt.Printf("%s %s %.1f\n", app.firstName, app.lastName, examScore(app, app.acceptedTo))
			}
		}
		fmt.Println()
	}

	return
}

func sortAdmittedApplicants(apps []Applicant) (sortedApps []Applicant) {
	sort.Slice(apps, func(i, j int) bool {
		if apps[i].acceptedTo != apps[j].acceptedTo {
			return apps[i].acceptedTo < apps[j].acceptedTo
		}
		if examScore(apps[i], apps[i].acceptedTo) != examScore(apps[j], apps[j].acceptedTo) {
			return examScore(apps[i], apps[i].acceptedTo) > examScore(apps[j], apps[j].acceptedTo)
		}
		if apps[i].firstName != apps[j].firstName {
			return apps[i].firstName < apps[j].firstName
		}
		return apps[i].lastName < apps[j].lastName
	})

	return apps
}

func getApplicants() (apps []Applicant) {
	file, _ := os.Open("applicants.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var app Applicant
		fmt.Sscanf(scanner.Text(), "%s %s %f %f %f %f %s %s %s",
			&app.firstName,
			&app.lastName,
			&app.physicsScore,
			&app.chemistryScore,
			&app.mathScore,
			&app.csScore,
			&app.pref1,
			&app.pref2,
			&app.pref3,
		)
		apps = append(apps, app)
	}

	return apps
}

func readInt() (num int) {
	fmt.Scanln(&num)
	return num
}
