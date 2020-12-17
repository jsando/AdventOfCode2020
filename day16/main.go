package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type fieldSpec struct {
	name       string
	min1, max1 int
	min2, max2 int
}

func main() {
	text, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	sections := strings.Split(strings.TrimSpace(string(text)), "\n\n")
	fields, err := parseFields(sections[0])
	if err != nil {
		panic(err)
	}
	tickets, err := parseTickets(sections[1])
	if err != nil {
		panic(err)
	}
	myTicket := tickets[0]
	tickets, err = parseTickets(sections[2])
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 1: %d\n", scanForErrors(tickets, fields))

	fieldsInOrder := mapFields(tickets, fields)
	multiple := 1
	for fieldNumber, field := range fieldsInOrder {
		if strings.HasPrefix(field.name, "departure") {
			multiple *= myTicket[fieldNumber]
		}
	}
	fmt.Printf("part 2: %d\n", multiple)
}

// find the sum of the invalid fields
func scanForErrors(tickets [][]int, fields []fieldSpec) int {
	errors := 0
	for _, ticket := range tickets {
		for _, value := range ticket {
			matches := matchField(value, fields)
			if len(matches) == 0 {
				errors += value
			}
		}
	}
	return errors
}

func matchField(value int, fields []fieldSpec) []fieldSpec {
	matches := []fieldSpec{}
	for _, field := range fields {
		if (value >= field.min1 && value <= field.max1) || (value >= field.min2 && value <= field.max2) {
			matches = append(matches, field)
		}
	}
	return matches
}

func mapFields(tickets [][]int, fields []fieldSpec) []fieldSpec {
	mapping := make([]map[string]bool, len(fields))
	for _, ticket := range tickets {
		if !isValid(ticket, fields) {
			continue
		}
		for fieldNumber, value := range ticket {
			matches := matchField(value, fields)
			names := mapping[fieldNumber]
			if names == nil {
				names = map[string]bool{}
				mapping[fieldNumber] = names
				for _, match := range matches {
					names[match.name] = true
				}
			} else {
				// remove any names that aren't in the current list
				for k := range names {
					found := false
					for _, f := range matches {
						if f.name == k {
							found = true
							break
						}
					}
					if !found {
						delete(names, k)
					}
				}
			}
		}
	}
	// iterate until narrow down 1 possiblity per
	for {
		changed := false
		for i := 0; i < len(mapping); i++ {
			if len(mapping[i]) == 1 {
				// remove from other fields
				name := getMapValue(mapping[i])
				for j := 0; j < len(mapping); j++ {
					if i == j {
						continue
					}
					if _, ok := mapping[j][name]; ok {
						delete(mapping[j], name)
						changed = true
					}
				}
			}
		}
		if !changed {
			break
		}
	}
	assignments := []fieldSpec{}
	for _, m := range mapping {
		name := getMapValue(m)
		for _, f := range fields {
			if f.name == name {
				assignments = append(assignments, f)
			}
		}
	}
	return assignments
}

func getMapValue(m map[string]bool) string {
	if len(m) != 1 {
		panic("map has zero or more than one entry")
	}
	for k := range m {
		return k
	}
	panic("cant happen")
}

func isValid(ticket []int, fields []fieldSpec) bool {
	for _, value := range ticket {
		matches := matchField(value, fields)
		if len(matches) == 0 {
			return false
		}
	}
	return true
}

// ex: "departure location: 39-715 or 734-949"
func parseFields(text string) ([]fieldSpec, error) {
	fields := []fieldSpec{}
	lines := strings.Split(text, "\n")
	var err error
	for _, line := range lines {
		field := fieldSpec{}
		field.name = strings.Split(line, ":")[0]
		ranges := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), " or ")
		r := strings.Split(ranges[0], "-")
		field.min1, err = strconv.Atoi(r[0])
		if err != nil {
			return nil, err
		}
		field.max1, err = strconv.Atoi(r[1])
		if err != nil {
			return nil, err
		}
		r = strings.Split(ranges[1], "-")
		field.min2, err = strconv.Atoi(r[0])
		if err != nil {
			return nil, err
		}
		field.max2, err = strconv.Atoi(r[1])
		if err != nil {
			return nil, err
		}
		fields = append(fields, field)
	}
	return fields, nil
}

func parseTickets(text string) ([][]int, error) {
	tickets := [][]int{}
	lines := strings.Split(text, "\n")
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		ticket := []int{}
		numbers := strings.Split(line, ",")
		for _, number := range numbers {
			val, err := strconv.Atoi(number)
			if err != nil {
				return nil, err
			}
			ticket = append(ticket, val)
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}
