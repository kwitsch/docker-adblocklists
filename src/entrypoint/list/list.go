package list

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

type List map[string]struct{}

func New() List {
	return make(List)
}

func (l List) Add(val string) {
	l[val] = struct{}{}
}

func (l List) Has(val string) bool {
	if _, ok := l[val]; ok {
		return true
	} else {
		return false
	}
}

func (l List) AddMap(m map[int]string) {
	for ak := range m {
		l.Add(m[ak])
	}
}

func (l List) AddOnlineMap(m map[int]string, verbose bool) {
	for ak := range m {
		if verbose {
			fmt.Println("List load from:", m[ak])
		}
		resp, err := http.Get(m[ak])
		if err == nil {
			defer resp.Body.Close()
			scanner := bufio.NewScanner(resp.Body)
			for scanner.Scan() {
				row := scanner.Text()
				if !strings.HasPrefix("#", row) {
					l.Add(row)
				}
			}

		}
	}
}

func (l List) ToString() (string, int) {
	res := ""
	entries := 0
	for row := range l {
		res += row + "\n"
		entries++
	}
	return strings.TrimSpace(res), entries
}
