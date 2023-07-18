package lib

import (
	morse "github.com/emanuelvsz/decode/lib/morse"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Converts a text into binary code.
func TextToBinary(text string) []string {
	binaries := make([]string, len(text))
	for i, char := range text {
		ascii := int(char)
		binary := strconv.FormatInt(int64(ascii), 2)
		binaries[i] = binary
	}
	return binaries
}

// Converts binary code into text.
func BinaryToText(binaries []string) string {
	text := ""
	for _, binary := range binaries {
		decimal, err := strconv.ParseInt(binary, 2, 64)
		if err != nil {
			fmt.Println("Error converting binary to decimal:", err)
			continue
		}
		character := string(decimal)
		text += character
	}
	return text
}

// Converts a text to hexadecimal.
func TextToHex(text string) string {
	hexadecimal := hex.EncodeToString([]byte(text))
	return hexadecimal
}

// Converts a hexadecimal to text.
func HexToText(hexadecimal string) (string, error) {
	bytes, err := hex.DecodeString(hexadecimal)
	if err != nil {
		return "", err
	}
	text := string(bytes)
	return text, nil
}

// Converts a character to its binary representation.
func CharToBinary(char byte) string {
	return fmt.Sprintf("%08b", char)
}

// Converts binary representation to its corresponding character.
func BinaryToChar(binary string) (byte, error) {
	decimal, err := strconv.ParseInt(binary, 2, 8)
	if err != nil {
		return 0, err
	}
	return byte(decimal), nil
}

// Converts a text to Morse code, adding slashes between words.
func TextToMorse(text string) string {
	morseCodeMap := morse.MorseData()

	text = strings.ToUpper(text)
	words := strings.Fields(text)

	var morseCode []string
	for _, word := range words {
		var wordCode string
		for _, char := range word {
			if code, ok := morseCodeMap[char]; ok {
				wordCode += code + " "
			}
		}
		if len(wordCode) > 0 {
			wordCode = wordCode[:len(wordCode)-1] // Remove the trailing whitespace.
			morseCode = append(morseCode, wordCode)
		}
	}

	return strings.Join(morseCode, " / ")
}

// Converts Morse code to text, removing slashes between words.
func MorseToText(morseCode string) string {
	morseCode = strings.TrimSpace(morseCode)

	words := strings.Split(morseCode, " / ")
	var text []string
	for _, wordCode := range words {
		var word string
		codes := strings.Split(wordCode, " ")
		for _, code := range codes {
			char, err := MorseToChar(code)
			if err != nil {
				fmt.Println("Error converting Morse code to character:", err)
				continue
			}
			word += char
		}
		if len(word) > 0 {
			text = append(text, word)
		}
	}

	return strings.Join(text, " ")
}

// Converts a character to its Morse code representation.
func CharToMorse(char byte) (string, error) {
	morseCodeMap := morse.MorseData()

	char = byte(unicode.ToUpper(rune(char)))
	if code, ok := morseCodeMap[rune(char)]; ok {
		return code, nil
	}

	return "", fmt.Errorf("no Morse code found for character '%c'", char)
}

// Converts Morse code representation to the corresponding character.
func MorseToChar(m string) (string, error) {
	morseCodeMap := morse.MorseData()

	for char, code := range morseCodeMap {
		if code == m {
			return string(char), nil
		}
	}

	return "", fmt.Errorf("no character found for Morse code '%s'", m)
}

// Converts Morse code representation to an array of arrays of strings,
// where each inner array represents a word and contains multiple binary strings representing each character.
func MorseToBinary(m string) ([][]string, error) {
	morseCodeMap := morse.MorseData()

	words := strings.Fields(m)
	result := make([][]string, 0)
	for _, word := range words {
		var binaryStrings []string
		letters := strings.Split(word, " ")
		for _, letter := range letters {
			if letter == "/" {
				binaryStrings = append(binaryStrings, "/")
			} else {
				found := false
				for char, code := range morseCodeMap {
					if code == letter {
						binaryStrings = append(binaryStrings, fmt.Sprintf("%08b", char))
						found = true
						break
					}
				}
				if !found {
					return nil, fmt.Errorf("no binary representation found for Morse code '%s'", letter)
				}
			}
		}
		if len(binaryStrings) > 0 {
			result = append(result, binaryStrings)
		}
	}

	return result, nil
}

// Converts the binary representation back to Morse code.
func BinaryToMorse(binary [][]string) (string, error) {
	morseCodeMap := morse.MorseData()

	var morseCode []string
	for _, word := range binary {
		var wordCode string
		for _, binaryString := range word {
			if binaryString == "/" {
				wordCode += "/"
			} else {
				code, err := strconv.ParseInt(binaryString, 2, 64)
				if err != nil {
					return "", fmt.Errorf("invalid binary representation: %s", binaryString)
				}
				char := byte(code)
				if morse, ok := morseCodeMap[rune(char)]; ok {
					wordCode += morse
				} else {
					return "", fmt.Errorf("no Morse code found for binary representation: %s", binaryString)
				}
			}
		}
		morseCode = append(morseCode, wordCode)
	}

	return strings.Join(morseCode, " "), nil
}

func main() {
	// Example usage
	text := "Hello World"
	binary := TextToBinary(text)
	fmt.Println("Text to Binary:", binary)

	decodedText := BinaryToText(binary)
	fmt.Println("Binary to Text:", decodedText)

	hexadecimal := TextToHex(text)
	fmt.Println("Text to Hex:", hexadecimal)

	decodedHex, err := HexToText(hexadecimal)
	if err != nil {
		fmt.Println("Error decoding Hex to Text:", err)
	} else {
		fmt.Println("Hex to Text:", decodedHex)
	}

	morseCode := TextToMorse(text)
	fmt.Println("Text to Morse Code:", morseCode)

	decodedMorse := MorseToText(morseCode)
	fmt.Println("Morse Code to Text:", decodedMorse)

	char := 'A'
	charBinary := CharToBinary(byte(char))
	fmt.Printf("Character '%c' to Binary: %s\n", char, charBinary)

	decodedChar, err := BinaryToChar(charBinary)
	if err != nil {
		fmt.Println("Error decoding Binary to Character:", err)
	} else {
		fmt.Printf("Binary to Character: '%c'\n", decodedChar)
	}

	morseBinary, err := MorseToBinary(morseCode)
	if err != nil {
		fmt.Println("Error converting Morse Code to Binary:", err)
	} else {
		fmt.Println("Morse Code to Binary:", morseBinary)
	}

	morseFromBinary, err := BinaryToMorse(morseBinary)
	if err != nil {
		fmt.Println("Error converting Binary to Morse Code:", err)
	} else {
		fmt.Println("Binary to Morse Code:", morseFromBinary)
	}
}
