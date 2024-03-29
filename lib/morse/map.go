package morse

func MorseData() map[rune]string {
	morseCodeMap := map[rune]string{
		'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".", 'F': "..-.", 'G': "--.",
		'H': "....", 'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.",
		'O': "---", 'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-", 'U': "..-",
		'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--", 'Z': "--..",
		'0': "-----", '1': ".----", '2': "..---", '3': "...--", '4': "....-", '5': ".....",
		'6': "-....", '7': "--...", '8': "---..", '9': "----.",
	}

	return morseCodeMap
}
