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

func (l List) AddOnlineMap(client *http.Client, m map[int]string, verbose bool) {
	var resp *http.Response
	var err error

	for ak := range m {
		if verbose {
			fmt.Println("List load from:", m[ak])
		}

		if resp, err = client.Get(m[ak]); err == nil {
			if resp.StatusCode == http.StatusOK {
				defer resp.Body.Close()
				entries := 0
				scanner := bufio.NewScanner(resp.Body)
				for scanner.Scan() {
					row := scanner.Text()
					if !strings.HasPrefix("#", row) {
						l.Add(row)
						entries++
					}
				}
				if verbose {
					fmt.Println("Entries added:", entries)
				}
			} else {
				_ = resp.Body.Close()
			}
		} else {
			fmt.Println("List load failed:", m[ak])
			fmt.Println(err.Error())
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
