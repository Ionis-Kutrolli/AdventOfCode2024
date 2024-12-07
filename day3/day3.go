package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

func MulBasicHandler(str string) error {
	re, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		return err
	}

	regMatch := re.FindAllStringSubmatch(str, -1)
	var sum int
	for _, r := range regMatch {
		num1, err := strconv.ParseInt(r[1], 10, 0)
		if err != nil {
			return err
		}
		num2, err := strconv.ParseInt(r[2], 10, 0)
		if err != nil {
			return err
		}
		sum += int(num1) * int(num2)
	}
	fmt.Println(sum)
	return nil
}

func MulAdvancedHandler(str string) error {
	re, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)|((do\(\)))|((don't\(\)))`)
	if err != nil {
		return err
	}

	regMatch := re.FindAllStringSubmatch(str, -1)
	var sum int
	do := true
	for _, r := range regMatch {
		switch r[0] {
		case "do()":
			do = true
			continue
		case "don't()":
			do = false
			continue
		}
		if do {
			num1, err := strconv.ParseInt(r[1], 10, 0)
			if err != nil {
				return err
			}
			num2, err := strconv.ParseInt(r[2], 10, 0)
			if err != nil {
				return err
			}
			sum += int(num1) * int(num2)
		}
	}
	fmt.Println(sum)
	return nil
}
