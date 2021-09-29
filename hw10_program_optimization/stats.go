package hw10programoptimization

import (
	"bufio"
	"io"
	"regexp"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	return countDomains(r, domain)
}

func countDomains(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)

	sc := bufio.NewScanner(r)

	var user User
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	re, err := regexp.Compile("\\." + domain)
	if err != nil {
		return nil, err
	}

	for sc.Scan() {
		if err = json.Unmarshal(sc.Bytes(), &user); err != nil {
			return nil, err
		}

		if re.Match([]byte(user.Email)) {
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]++
		}
	}
	return result, nil
}
