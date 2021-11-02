package main

import (
	"fmt"
	"strings"
)

func main() {
	/*	validatePen(2, compute_penalty("1 0 1 0", 2))
		validatePen(3, compute_penalty("0 0 0 1 1 1 1", 0))
		validatePen(4, compute_penalty("0 0 0 1 1 1 1", 7))
		validatePen(0, compute_penalty("0 0 0 1 1 1 1", 3))
		validatePen(0, compute_penalty("", 0))
		validatePen(1, compute_penalty("0 1 0 1 1 1 1", 3))
	*/
	//find_best_removal_time("0 0 1 1") should return 2

	validate(2, find_best_removal_time("0 0 1 1"))

	validate(3, find_best_removal_time("0 0 0 1 1 1 1"))
	validate(0, find_best_removal_time(""))
	validate(0, find_best_removal_time("1 1 1 1"))

	validate(4, find_best_removal_time("0 0 0 0"))
	validate(5, find_best_removal_time("1 0 0 0 0 1 1 1 0 1 1 0 0 1 1 1 1 0 0 1 1 0 1 1 1"))
	validate(0, find_best_removal_time("1 1 1 1 1 0 0 0 1 1 1 1 0 0 0 1 1 1 0 1 0 0 1 0 1"))
	validate(25, find_best_removal_time("0 0 1 1 1 0 0 1 0 0 1 1 1 0 0 1 1 0 0 0 1 0 1 0 0"))

}

func compute_penalty(logs string, remove_at int) int {
	//compute_penalty("1 0 1 0", 2) should return 2
	log := parseLog(logs)
	var penalty int
	for i, status := range log {
		if i < remove_at && status == "1" {
			penalty++
		} else if i >= remove_at && status == "0" {
			penalty++
		}
	}
	return penalty
}

func find_best_removal_time(logs string) int {
	log := parseLog(logs)
	smallest := len(log)
	remove_at := len(log)
	for tmp_remove_at := 0; tmp_remove_at <= len(log); tmp_remove_at++ {
		tmp := compute_penalty(logs, tmp_remove_at)
		if tmp < smallest {
			smallest = tmp
			remove_at = tmp_remove_at
		}
	}
	return remove_at
}

func parseLog(logs string) []string {
	return strings.Split(logs, " ")
}

func validate(want int, got int) {
	fmt.Printf("want %d, got %d\n", want, got)
}
