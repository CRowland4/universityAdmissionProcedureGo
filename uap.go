package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Applicant struct {
	firstName, lastName, pref1, pref2, pref3, acceptedTo             string
	physicsScore, chemistryScore, mathScore, csScore, admissionScore float64
}

func main() {
	departmentCapacity := readInt()
	applicants := getApplicants()
	admittedApplicants := executeAdmissionsProcess(applicants, departmentCapacity)
	storeAdmittedApplicants(admittedApplicants)
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
		if bestScore(apps[i], getPreference(apps[i], round)) != bestScore(apps[j], getPreference(apps[j], round)) {
			return bestScore(apps[i], getPreference(apps[i], round)) > bestScore(apps[j], getPreference(apps[j], round))
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

func bestScore(app Applicant, department string) (score float64) {
	switch department {
	case "Biotech":
		return math.Max(average(app.physicsScore, app.chemistryScore), app.admissionScore)
	case "Physics":
		return math.Max(average(app.physicsScore, app.mathScore), app.admissionScore)
	case "Mathematics":
		return math.Max(app.mathScore, app.admissionScore)
	case "Engineering":
		return math.Max(average(app.csScore, app.mathScore), app.admissionScore)
	default:
		return math.Max(app.chemistryScore, app.admissionScore)
	}
}

func average(nums ...float64) (result float64) {
	for _, num := range nums {
		result += num
	}
	result /= float64(len(nums))
	return result
}

func storeAdmittedApplicants(apps []Applicant) {
	apps = sortAdmittedApplicants(apps)
	for _, dep := range [5]string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"} {
		storeDepApplicants(dep, apps)
	}

	fmt.Println("Admitted applicants have been stored in one text file for each department.")
	return
}

func storeDepApplicants(dep string, apps []Applicant) {
	file, _ := os.Create(strings.ToLower(dep) + ".txt")
	defer file.Close()

	for _, app := range apps {
		if app.acceptedTo == dep {
			fmt.Fprintf(file, "%s %s %.1f\n", app.firstName, app.lastName, bestScore(app, app.acceptedTo))
		}
	}

	return
}

func sortAdmittedApplicants(apps []Applicant) (sortedApps []Applicant) {
	sort.Slice(apps, func(i, j int) bool {
		if apps[i].acceptedTo != apps[j].acceptedTo {
			return apps[i].acceptedTo < apps[j].acceptedTo
		}
		if bestScore(apps[i], apps[i].acceptedTo) != bestScore(apps[j], apps[j].acceptedTo) {
			return bestScore(apps[i], apps[i].acceptedTo) > bestScore(apps[j], apps[j].acceptedTo)
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
		fmt.Sscanf(scanner.Text(), "%s %s %f %f %f %f %f %s %s %s",
			&app.firstName,
			&app.lastName,
			&app.physicsScore,
			&app.chemistryScore,
			&app.mathScore,
			&app.csScore,
			&app.admissionScore,
			&app.pref1,
			&app.pref2,
			&app.pref3,
		)
		apps = append(apps, app)
	}

	return apps
}

func readInt() (num int) {
	fmt.Println("Enter the capacity of the departments: ")
	fmt.Scanln(&num)
	return num
}
