package main

import(
	"fmt"
	"os"
	"bufio"
	"myStack"
	"strconv"
	"strings"
	"unicode"
)

func Transfer(exp string) string {		//infix transfer to suffix
    stack := myStack.New()
    suffix := ""
    expLen := len(exp)
    // traversal
    for i := 0; i < expLen; i++ {

        char := string(exp[i])

        switch char {
        case " ":
            continue
        case "(":
            stack.Push("(")
        case ")":
            for !stack.IsEmpty() {
                preChar := stack.Top()
                if preChar == "(" {
                    stack.Pop() 
                    break
                }
                suffix += preChar.(string)
                stack.Pop()
            }

        case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
            j := i
            digit := ""
            for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {		//check if it is a number
                digit += string(exp[j])
            }
            suffix += digit
            i = j - 1 

        default:
            for !stack.IsEmpty() {
				t := stack.Top()
				tt := t.(byte)
                if tt == 40 || isLower(tt, char) {	//pop if meet a lower operator
                    break
                }
                suffix += string(tt)
                stack.Pop()
            }
            stack.Push(char)
        }
    }
    for !stack.IsEmpty() {
        suffix += stack.Top().(string)
        stack.Pop()
    }

    return suffix
}

// compare top operator
func isLower(top interface{}, newTop interface{}) bool {
    switch top {
    case "+", "-":
        if newTop == "*" || newTop == "/" {
            return true
        }
    case "(":
        return true
    }
    return false
}

func calculate(suffix string) int {
    stack := myStack.New()
    sufLen := len(suffix)
    for i := 0; i < sufLen; i++ {
        nextChar := string(suffix[i])
        if unicode.IsDigit(rune(suffix[i])) {
            stack.Push(nextChar)
        } else {
            num1, _ := strconv.Atoi(stack.Pop().(string))
            num2, _ := strconv.Atoi(stack.Pop().(string))
            switch nextChar {
			case "+":
				stack.Push(strconv.Itoa(num1 + num2))
			case "-":
				stack.Push(strconv.Itoa(num1 - num2))
			case "*":
				stack.Push(strconv.Itoa(num1 * num2))
			case "/":
				stack.Push(strconv.Itoa(num1 / num2))
            }
        }
    }
    result, _ := strconv.Atoi(stack.Top().(string))
    return result
}


func main(){
	print("please input a math statement: ")
	var stat string
	reader := bufio.NewReader(os.Stdin)
	stat, _ = reader.ReadString('\n')
	stat = strings.TrimSpace(stat)
	suffix := Transfer(stat)
	fmt.Printf("The answer is: %.2f\n", calculate(suffix))
}
