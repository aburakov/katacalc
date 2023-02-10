package ServiceLogic

import (
	"bufio"
	"fmt"
	"github.com/imbue11235/roman"
	"log"
	"os"
	"strings"
)

type CalcReader struct {
	ArabDigits map[string]int32
	RomeDigits map[string]int32
	LocalCalc  Calc
}

func (p *CalcReader) Read() {
	availableOperations := []string{"+", "-", "*", "/"}
	(*p).ArabDigits = map[string]int32{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10}
	(*p).RomeDigits = map[string]int32{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

	reader := bufio.NewReader(os.Stdin)

	for {
		checked := false
		fmt.Printf("\nВведите значение: ")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("%s", err.Error())
		}

		result := strings.Split(text, " ")

		if len(result) < 3 {
			fmt.Printf("%s", "Вывод ошибки, так как строка не является математической операцией.\n")
			continue
		} else if len(result) > 3 {
			fmt.Printf("%s", "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).\n")
			continue
		}

		for _, val := range availableOperations {
			if val == result[1] {
				checked = true
			}
		}
		if !checked {
			panic("Такой операции не существует")
		}

		result[len(result)-1] = strings.TrimSpace(result[len(result)-1])

		if p.CheckDigits(result[0], result[2]) {
			check, aParsed, bParsed := p.CheckNotaion(result[0], result[2])
			if check == "error" {
				fmt.Printf("Вывод ошибки, так как используются одновременно разные системы счисления.\n")
				break
			}
			if check == "rome" {
				result := p.MakeOperation(aParsed, bParsed, result[1])
				if result < 1 {
					fmt.Printf("Вывод ошибки, так как в римской системе нет отрицательных чисел.\n")
					continue
				}
				value := roman.FromArabic(int(result))
				fmt.Printf("%s", value)
			} else if check == "arab" {
				fmt.Printf("%d", p.MakeOperation(aParsed, bParsed, result[1]))
			}
		} else {
			panic("Не подходящие числа, завершаем работу")
		}
	}
}

func (p *CalcReader) CheckNotaion(a, b string) (string, int32, int32) {
	aArab, aIsIncludedArab := (*p).ArabDigits[a]
	bArab, bIsIncludedArab := (*p).ArabDigits[b]

	aRome, aIsIncludedRome := (*p).RomeDigits[a]
	bRome, bIsIncludedRome := (*p).RomeDigits[b]

	if aIsIncludedArab && bIsIncludedArab {
		return "arab", aArab, bArab
	} else if aIsIncludedRome && bIsIncludedRome {
		return "rome", aRome, bRome
	} else {
		return "error", 0, 0
	}
}

func (p *CalcReader) CheckDigits(a, b string) bool {
	_, aIsIncludedArab := (*p).ArabDigits[a]
	_, bIsIncludedArab := (*p).ArabDigits[b]

	_, aIsIncludedRome := (*p).RomeDigits[a]
	_, bIsIncludedRome := (*p).RomeDigits[b]

	if aIsIncludedArab && bIsIncludedArab {
		return true
	} else if aIsIncludedRome && bIsIncludedRome {
		return true
	} else {
		return false
	}
}

func (p *CalcReader) MakeOperation(a, b int32, operation string) (result int32) {
	p.LocalCalc = Calc{}
	result = 0
	if operation == "+" {
		result = p.LocalCalc.Add(a, b)
	} else if operation == "-" {
		result = p.LocalCalc.Minus(a, b)
	} else if operation == "*" {
		result = p.LocalCalc.Multiply(a, b)
	} else if operation == "/" {
		result = p.LocalCalc.Division(a, b)
	}
	return
}
