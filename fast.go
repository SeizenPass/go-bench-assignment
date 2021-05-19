package main

import (
	"fmt"
	user2 "hw3_bench/user"
	"io"
	"io/ioutil"
	"strings"
)

func FastSearch(out io.Writer) {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Amiran: превратил слайс в мапу, хотя перформанс мог измениться в худшую сторону на основе бенчмарков
	seenBrowsers := make(map[string]interface{})
	// Amiran: удалил лишнюю переменную
	lines := strings.Split(string(fileContents), "\n")
	// Amiran: храню только одну структуру, которую перезаписываю через easyjson
	user := &user2.User{}
	fmt.Fprintln(out, "found users:")
	// Amiran: совместил все в один цикл, для большей скорости
	for i, line := range lines {
		// fmt.Printf("%v %v\n", err, line)
		// Amiran: использую easyjson
		err := user.UnmarshalJSON([]byte(line))
		if err != nil {
			panic(err)
		}
		isAndroid := false
		isMSIE := false

		browsers := user.Browsers
		// Amiran: Вместо двух циклов сделал один, хоть это и не сильно влияет
		for _, browserRaw := range browsers {
			browser := browserRaw
			// Amiran: Поменял regexp.MatchString на strings.Contains, так как мы особо тяжелых регулярок не делаем
			if strings.Contains(browser, "Android") {
				isAndroid = true
				notSeenBefore := true
				if seenBrowsers[browser] != nil {
					notSeenBefore = false
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					var s struct{}
					seenBrowsers[browser] = s
				}
			}
			// Amiran: Поменял regexp.MatchString на strings.Contains, так как мы особо тяжелых регулярок не делаем
			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				notSeenBefore := true
				if seenBrowsers[browser] != nil {
					notSeenBefore = false
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					var s struct{}
					seenBrowsers[browser] = s
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		email := strings.Replace(user.Email, "@", " [at] ", 1)
		fmt.Fprintf(out, "[%d] %s <%s>\n", i, user.Name, email)
	}
	fmt.Fprintf(out, "\n")
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}
