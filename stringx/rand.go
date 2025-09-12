package stringx

import "math/rand"

func RandLetters(n int) string {
	if n <= 0 {
		return ""
	}
	return RandStringWithSeed(n, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
}

func RandLowerLetters(n int) string {
	if n <= 0 {
		return ""
	}
	return RandStringWithSeed(n, []rune("abcdefghijklmnopqrstuvwxyz"))
}

func RandUpperLetters(n int) string {
	if n <= 0 {
		return ""
	}
	return RandStringWithSeed(n, []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
}

func RandNumbers(n int) string {
	if n <= 0 {
		return ""
	}
	return RandStringWithSeed(n, []rune("0123456789"))
}

func RandLetterNumbers(n int) string {
	if n <= 0 {
		return ""
	}
	return RandStringWithSeed(n, []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
}

func RandLowerLetterNumbers(n int) string {
	if n <= 0 {
		return ""
	}
	return RandStringWithSeed(n, []rune("0123456789abcdefghijklmnopqrstuvwxyz"))
}

func RandUpperLetterNumbers(n int) string {
	if n <= 0 {
		return ""
	}
	return RandStringWithSeed(n, []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
}

func RandStringWithSeed(n int, seed []rune) string {
	if n <= 0 {
		return ""
	}
	if len(seed) == 0 {
		return ""
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = seed[rand.Intn(len(seed))]
	}
	return string(b)
}
