package main

import (
    "fmt"
    "strings"
)

func numUniqueEmails(emails []string) int {
    tmp := make(map[string] int)
    for _,v := range emails {
        localDomain := strings.Split(v, "@")
        dotV := ""
        tmpDot := strings.Split(localDomain[0], ".")
        for _,tv := range tmpDot {
            dotV += tv
        }
        tmpPlus := strings.Split(dotV, "+")
        newEmail := tmpPlus[0] + localDomain[1]
        tmp[newEmail] = 1
    }
    return len(tmp)
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
