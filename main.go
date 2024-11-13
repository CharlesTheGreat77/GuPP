package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Person struct {
	FirstName       string
	LastName        string
	NickName        string
	PartnerName     string
	PartnerNickname string
	ChildName       string
	ChildNickname   string
	PetName         string
	CompanyName     string
	KeyWords        []string
}

func main() {
	cowsay()
	info := collectInfo() // get user input

	fmt.Printf("\n[*] Generating wordlist... please wait...\n")
	startTime := time.Now()
	wordlist := generateWordlist(info) // generate wordlist

	file, err := os.Create("wordlist.txt") // default output file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("[*] Wordlist Length: \033[31m%d\033[0m\n", len(wordlist))

	// save to file
	fmt.Println("[*] Saving wordlist to file...\n")
	for _, word := range wordlist {
		file.WriteString(word + "\n")
		//fmt.Printf("\r[wordlist.txt] \033[31m%s\033[0m\n", word)
	}

	elapsedTime := time.Since(startTime)

	fmt.Println("\n[!] Wordlist Generated and saved to wordlist.txt")
	fmt.Printf(" -> Time Lapse: %s", elapsedTime)
}

// replace chars with common replacements, change/add as necessary
func replaceChars(word string) []string {
	replacements := map[string]string{
		"a": "@",
		"o": "0",
		"l": "1",
		"s": "$",
		"i": "1",
		"S": "5",
		"e": "3",
	}

	var variations []string
	variations = append(variations, word)
	replacedWord := word
	for original, replacement := range replacements {
		replacedWord = strings.ReplaceAll(replacedWord, original, replacement)
		replacedWord = strings.ReplaceAll(replacedWord, strings.ToUpper(original), replacement)
	}

	if replacedWord != word {
		variations = append(variations, replacedWord)
	}

	return variations

}

func generateWordlist(target Person) []string {
	uniqueWords := make(map[string]bool)

	var wordlist []string

	var numbers []string
	for i := 0; i < 10000; i++ {
		numbers = append(numbers, fmt.Sprintf("%d", i), fmt.Sprintf("%04d", i))
	}

	names := []string{
		target.FirstName, target.LastName, target.NickName, target.PartnerName, target.PartnerNickname,
		target.ChildName, target.ChildNickname, target.PetName, target.CompanyName,
	}

	wordChannel := make(chan string)

	var wg sync.WaitGroup

	// go routine for shit.. turbo baby
	processNameCombination := func(name string) {
		defer wg.Done()
		if name == "" {
			return
		}

		firstInitial := strings.ToUpper(string(name[0]))
		lastInitial := strings.ToUpper(string(name[len(name)-1]))

		firstLetterUpper := strings.ToUpper(string(name[0]))
		firstLetterLower := strings.ToLower(string(name[0]))
		nameLower := strings.ToLower(name)
		nameUpper := strings.ToUpper(name[:1]) + name[1:]

		originalCombinations := []string{
			name + nameUpper[:1],
			nameLower + nameUpper[:1],
			nameUpper + name[1:],
			nameLower + name[1:],
			firstLetterUpper + name[1:],
			firstLetterLower + name[1:],
		}

		for _, nameCombo := range originalCombinations {
			for _, num := range numbers {
				combos := []string{
					nameCombo + num,
					num + nameCombo,
				}

				for _, combo := range combos {
					for _, variation := range replaceChars(combo) {
						wordChannel <- variation
					}
				}
			}
		}

		// more lil combination bs aside from looping through
		nameCombinations := []string{
			target.FirstName + target.LastName,
			target.LastName + target.FirstName,
			target.FirstName + lastInitial,
			lastInitial + target.FirstName,
			firstInitial + target.LastName,
			target.LastName + firstInitial,
		}

		for _, nameCombo := range nameCombinations {
			for _, num := range numbers {
				combos := []string{
					nameCombo + num,
					num + nameCombo,
				}

				for _, combo := range combos {
					for _, variation := range replaceChars(combo) {
						wordChannel <- variation
					}
				}
			}
		}
	}

	for _, name := range names {
		wg.Add(1)
		go processNameCombination(name)
	}

	for _, keyWord := range target.KeyWords {
		if keyWord == "" {
			continue
		}

		wg.Add(1)
		go func(keyWord string) {
			defer wg.Done()
			for _, variation := range replaceChars(keyWord) {
				wordChannel <- variation
			}
		}(keyWord)
	}

	go func() {
		wg.Wait()
		close(wordChannel)
	}()

	for word := range wordChannel {
		// try to prevent duplicates
		if _, exists := uniqueWords[word]; !exists {
			uniqueWords[word] = true
			wordlist = append(wordlist, word)
		}
	}

	return wordlist
}

// get info from user
func collectInfo() Person {
	fmt.Println("\n\n[*] Insert information about the target to make a dictionary\n -> Enter for unknown information is accepted...\n")

	var info Person
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("[?] First Name: ")
	info.FirstName = readAndTrim(reader)

	fmt.Printf("[?] Last Name: ")
	info.LastName = readAndTrim(reader)

	fmt.Printf("[?] Nick Name: ")
	info.NickName = readAndTrim(reader)

	fmt.Printf("[?] Partner's First Name: ")
	info.PartnerName = readAndTrim(reader)

	fmt.Printf("[?] Partner's Nick Name: ")
	info.PartnerNickname = readAndTrim(reader)

	fmt.Printf("[?] Child's Name: ")
	info.ChildName = readAndTrim(reader)

	fmt.Printf("[?] Child's Nick Name: ")
	info.ChildNickname = readAndTrim(reader)

	fmt.Printf("[?] Pet's Name: ")
	info.PetName = readAndTrim(reader)

	fmt.Printf("[?] Company Name: ")
	info.CompanyName = readAndTrim(reader)

	fmt.Printf("[?] Do you want to add keywords to the wordlist (y/n): ")
	option := readAndTrim(reader)

	if strings.ToLower(option) == "y" {
		fmt.Printf("[*] Key words separated by spaces (e.g., hacker juice black): ")
		keyWordsInput := readAndTrim(reader)
		info.KeyWords = strings.Fields(keyWordsInput)
	}

	return info
}

// get user input
func readAndTrim(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func cowsay() {
	fmt.Printf(` ______
< GuPP >
 ------
        \   ^__^
         \  (oo)\_______      // Go User
            (__)\       )\/\  // Password
                ||----w |     // Profiler
                ||     ||     [Credits: Mebus | https://github.com/Mebus/]
                            --[Go Written: DoobTheGoober]
  `)
}
