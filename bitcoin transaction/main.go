package main

import "unicode"


type Cipher interface{

	Encryption(string) string
	Decryption(string) string 

}

type Cipher [] int 

func (c cipher) cipherAlgorithm(letters string, shift func(int, int)int) string{
	shiftedText := ""
	for _, letter := range letters {
		if !unicode.IsLetter(letters){
			continue
		}
		shiftDist := c [len(shiftedText)%len(c)]
		s := shift(int(unicode.ToLower(letter)), shiftDist)

		switch {
		case s < 'a':
			s+= 'z' - 'a' + 1
		case 'z' < s:
			s -= 'z' - 'a' + 1
		}
		shiftedText += string(s)
	}
	return shiftedText
}

func (c *cipher) Encryption(plainText string) string {
	return c.cipherAlgorithm(plainText, func(a, b int) int { return a + b })
}

func (c *cipher) Decryption(cipherText string) string {
	return c.cipherAlgorithm(cipherText, func(a, b int) int { return a - b })
}

func NewCase(key int) Cipher {
	return NewShift(key)
}

func NewShift(key int) Cipher {
	if shift < -25 || 25 < shift || shift == 0 {
		return nil
	}
	c:= cipher([]int{shift})
	return &c
}

func main(){
	c := NewCase(1)
	fmt.Println("encrypt key(01) abcd => , " c.Encryption("abcd"))	
	fmt.Println("decrypt key(01) pstr => , " c.Decryption("pstr"))
	fmt.Println()
}