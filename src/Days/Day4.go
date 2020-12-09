package Days

import (
	"../Constants"
	"../Input"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day4 struct{}

type passport struct {
	byr string // (Birth Year)
	iyr string //(Issue Year)
	eyr string //(Expiration Year)
	hgt string //(Height)
	hcl string //(Hair Color)
	ecl string //(Eye Color)
	pid string //(Passport ID)
	cid string //(Country ID)
}

func (d Day4) Calc() string {
	input := Input.ReadInputFile(4)
	var passports []passport

	for _, passportInput := range strings.Split(input, Constants.LineSeparator+Constants.LineSeparator) {
		r_byr := regexp.MustCompile("byr:([0-9]{4})")
		r_iyr := regexp.MustCompile("iyr:([0-9]{4})")
		r_eyr := regexp.MustCompile("eyr:([0-9]{4})")
		r_hgt := regexp.MustCompile("hgt:([0-9]+(?:cm|in))")
		r_hcl := regexp.MustCompile("hcl:(#?[0-9a-z]+)")
		r_ecl := regexp.MustCompile("ecl:(#?[0-9a-z]+)")
		r_pid := regexp.MustCompile("pid:#?([0-9a-z]+)")
		r_cid := regexp.MustCompile("cid:#?([0-9a-z]+)")

		newP := passport{
			byr: atIndexOrEmpty(r_byr.FindStringSubmatch(passportInput), 1),
			iyr: atIndexOrEmpty(r_iyr.FindStringSubmatch(passportInput), 1),
			eyr: atIndexOrEmpty(r_eyr.FindStringSubmatch(passportInput), 1),
			hgt: atIndexOrEmpty(r_hgt.FindStringSubmatch(passportInput), 1),
			hcl: atIndexOrEmpty(r_hcl.FindStringSubmatch(passportInput), 1),
			ecl: atIndexOrEmpty(r_ecl.FindStringSubmatch(passportInput), 1),
			pid: atIndexOrEmpty(r_pid.FindStringSubmatch(passportInput), 1),
			cid: atIndexOrEmpty(r_cid.FindStringSubmatch(passportInput), 1),
		}

		passports = append(passports, newP)
	}

	return fmt.Sprintf("1: %d\n2: %d\n", calc1(passports), validate(passports))
}

func calc1(s []passport) int {
	count := 0

	for _, p := range s {
		if p.byr != "" &&
			p.iyr != "" &&
			p.eyr != "" &&
			p.hgt != "" &&
			p.hcl != "" &&
			p.ecl != "" &&
			p.pid != "" {
			count++
		}
	}

	return count
}

func validate(passports []passport) int {
	validCount := 0

	for _, passport := range passports {
		fmt.Println(passport)

		if passport.byr == "" ||
			passport.iyr == "" ||
			passport.eyr == "" ||
			passport.hgt == "" ||
			passport.hcl == "" ||
			passport.ecl == "" ||
			passport.pid == "" {
			continue
		}

		birthYear, _ := strconv.Atoi(passport.byr)
		if len(passport.byr) != 4 || birthYear < 1920 || birthYear > 2002 {
			continue
		}

		issueYear, _ := strconv.Atoi(passport.iyr)
		if len(passport.iyr) != 4 || issueYear < 2010 || issueYear > 2020 {
			continue
		}

		expirationYear, _ := strconv.Atoi(passport.eyr)
		if len(passport.iyr) != 4 || expirationYear < 2020 || expirationYear > 2030 {
			continue
		}

		if len(passport.hgt) > 2 {
			heightUnit := passport.hgt[len(passport.hgt)-2:]
			height, err := strconv.Atoi(passport.hgt[:len(passport.hgt)-2])
			if err != nil {
				panic(err)
			}
			switch heightUnit {
			case "cm":
				if height < 150 || height > 193 {
					continue
				}
			case "in":
				if height < 59 || height > 76 {
					continue
				}
			}
		}

		if !regexp.MustCompile("^#[0-9a-f]{6}$").MatchString(passport.hcl) {
			continue
		}

		if !regexp.MustCompile(passport.ecl).MatchString("amb blu brn gry grn hzl oth") {
			continue
		}

		if !regexp.MustCompile("^[0-9]{9}$").MatchString(passport.pid) {
			continue
		}

		validCount++
	}

	return validCount
}

func atIndexOrEmpty(s []string, index int) string {
	if index < len(s) {
		return s[index]
	}

	return ""
}
