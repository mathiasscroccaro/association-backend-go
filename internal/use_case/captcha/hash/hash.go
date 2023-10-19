package hash

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func IsHashedKeyEqualToSolution(hashedKey, solution, secretKey string) bool {
	unhashedSolution, err := UnhashSolution(hashedKey, secretKey)

	if err != nil {
		return false
	}

	if unhashedSolution == solution {
		return true
	} else {
		return false
	}
}

func HashSolutionBySecretKey(solution, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"solution": solution,
		"exp":      time.Now().Add(time.Minute * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	hashedSolution, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return hashedSolution, nil
}

func UnhashSolution(hashedSolution, secretKey string) (string, error) {
	token, err := jwt.Parse(hashedSolution, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if solution, ok := claims["solution"].(string); ok {
			return string(solution), nil
		}

	}

	return "", fmt.Errorf("Invalid solution")
}
