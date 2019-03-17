package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	tmp := make(map[string] int)
	for _, v := range emails {
		var newEmail []rune
		domainFlag := false
		plusFlag := false
		for _, cv := range []rune(v) {
			if domainFlag {
				newEmail = append(newEmail, cv)
				continue
			} else {
				if cv == '@' {
					domainFlag = true
					newEmail = append(newEmail, cv)
				}
				if !plusFlag {

					if cv == '.' {
						continue
					} else if cv == '+' {
						plusFlag = true
						continue
					} else {
						newEmail = append(newEmail, cv)
					}
				}
			}
		}
		tmp[string(newEmail)] = 1
	}
	return len(tmp)
}

func search(emails []string, email string) bool {
	for _, v := range emails {
		if strings.Compare(v, email) == 0 {
			return true
		}
	}
	return false
}

func main() {
	ems := []string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}
	i := numUniqueEmails(ems)
	fmt.Printf("num", i)
}
