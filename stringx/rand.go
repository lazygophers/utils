package stringx

import (
	"github.com/lazygophers/log"
	"math/rand"
)

func RandLetters(n int) string {
	log.Debugf("RandLetters: generating %d random letters", n)
	if n <= 0 {
		log.Warn("RandLetters: non-positive length provided")
		return ""
	}
	result := RandStringWithSeed(n, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	log.Debugf("RandLetters: generated string of length %d", len(result))
	return result
}

func RandLowerLetters(n int) string {
	log.Debugf("RandLowerLetters: generating %d random lowercase letters", n)
	if n <= 0 {
		log.Warn("RandLowerLetters: non-positive length provided")
		return ""
	}
	result := RandStringWithSeed(n, []rune("abcdefghijklmnopqrstuvwxyz"))
	log.Debugf("RandLowerLetters: generated string of length %d", len(result))
	return result
}

func RandUpperLetters(n int) string {
	log.Debugf("RandUpperLetters: generating %d random uppercase letters", n)
	if n <= 0 {
		log.Warn("RandUpperLetters: non-positive length provided")
		return ""
	}
	result := RandStringWithSeed(n, []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	log.Debugf("RandUpperLetters: generated string of length %d", len(result))
	return result
}

func RandNumbers(n int) string {
	log.Debugf("RandNumbers: generating %d random numbers", n)
	if n <= 0 {
		log.Warn("RandNumbers: non-positive length provided")
		return ""
	}
	result := RandStringWithSeed(n, []rune("0123456789"))
	log.Debugf("RandNumbers: generated string of length %d", len(result))
	return result
}

func RandLetterNumbers(n int) string {
	log.Debugf("RandLetterNumbers: generating %d random alphanumeric characters", n)
	if n <= 0 {
		log.Warn("RandLetterNumbers: non-positive length provided")
		return ""
	}
	result := RandStringWithSeed(n, []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	log.Debugf("RandLetterNumbers: generated string of length %d", len(result))
	return result
}

func RandLowerLetterNumbers(n int) string {
	log.Debugf("RandLowerLetterNumbers: generating %d random lowercase alphanumeric characters", n)
	if n <= 0 {
		log.Warn("RandLowerLetterNumbers: non-positive length provided")
		return ""
	}
	result := RandStringWithSeed(n, []rune("0123456789abcdefghijklmnopqrstuvwxyz"))
	log.Debugf("RandLowerLetterNumbers: generated string of length %d", len(result))
	return result
}

func RandUpperLetterNumbers(n int) string {
	log.Debugf("RandUpperLetterNumbers: generating %d random uppercase alphanumeric characters", n)
	if n <= 0 {
		log.Warn("RandUpperLetterNumbers: non-positive length provided")
		return ""
	}
	result := RandStringWithSeed(n, []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	log.Debugf("RandUpperLetterNumbers: generated string of length %d", len(result))
	return result
}

func RandStringWithSeed(n int, seed []rune) string {
	log.Debugf("RandStringWithSeed: generating %d characters from seed of length %d", n, len(seed))
	if n <= 0 {
		log.Warn("RandStringWithSeed: non-positive length provided")
		return ""
	}
	if len(seed) == 0 {
		log.Error("RandStringWithSeed: empty seed provided")
		return ""
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = seed[rand.Intn(len(seed))]
	}
	result := string(b)
	log.Debugf("RandStringWithSeed: generated string of length %d", len(result))
	return result
}
