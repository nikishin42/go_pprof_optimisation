package hw3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const filePath2 string = "./data/u2.txt"

//easyjson:json
type myUser struct {
	Browsers []string `json:"browsers"`
	//Company  string   `json:"company"`
	//Country  string   `json:"country"`
	Email string `json:"email"`
	//Job      string   `json:"job"`
	Name string `json:"name"`
	//Phone string `json:"phone"`
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	var sb strings.Builder
	sb.WriteString("found users:\n")

	seenBrowsers := make([]string, 0, 1000)
	uniqueBrowsers := 0

	user := myUser{}
	i := -1

	reader := bufio.NewReader(file)

	for {
		i++
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		err = user.UnmarshalJSON(line)
		if err != nil {
			panic(err)
		}

		isAndroid := false
		isMSIE := false

		browsers := user.Browsers

		for _, browser := range browsers {
			if strings.Contains(browser, "Android") {
				isAndroid = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
						break
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		for _, browser := range browsers {
			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
						break
					}
				}
				if notSeenBefore {
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		emailNew := strings.Replace(user.Email, "@", " [at] ", -1)
		sb.WriteString("[")
		sb.WriteString(strconv.Itoa(i))
		sb.WriteString("] ")
		sb.WriteString(user.Name)
		sb.WriteString(" <")
		sb.WriteString(emailNew)
		sb.WriteString(">\n")
	}

	sb.WriteString("\nTotal unique browsers ")
	sb.WriteString(strconv.Itoa(uniqueBrowsers))
	fmt.Fprintln(out, sb.String())

}
