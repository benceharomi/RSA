package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func mod(a, b int64) int64 {
	res := a % b
	if res < 0 {
		res += b
	}
	return res
}

func euclidean(a int64, b int64) ([]int64, int64) {
	ai := a
	bi := b

	var quotients []int64
	for i := 1; bi != 0; i++ {
		if i != 1 {
			quotients = append(quotients, ai/bi)
		}
		r := mod(ai, bi)
		ai = bi
		bi = r
	}
	return quotients, ai
}

func extended_euclidean(a, b int64, quotients []int64) int64 {
	p0 := int64(0)
	p1 := int64(1)
	for i := 0; i < len(quotients)-1; i++ {
		p2 := mod(p0-p1*quotients[i], b)
		p0 = p1
		p1 = p2
	}
	return p1
}

func exponentiation(a int64, decimal_k int64, n int64) int64 {
	b := int64(1)
	if decimal_k == 0 {
		return b
	}
	binary_k := strconv.FormatInt(decimal_k, 2)
	binary_k_array := strings.Split(binary_k, "")

	A := a
	length := len(binary_k_array) - 1
	for i := length; i >= 0; i-- {
		if i == length {
			if binary_k_array[i] == "1" {
				b = a
			}
		} else {
			A = mod(A*A, n)
			if binary_k_array[i] == "1" {
				b = mod(A*b, n)
			}
		}
	}
	return b
}

func calculate_decryption_key(e, phi int64) (int64, error) {
	if e <= 1 || phi <= e {
		return 0, fmt.Errorf("e should be greater than 1")
	}
	if phi <= e {
		return 0, errors.New("e should smaller than n")
	}
	quotients, ai := euclidean(e, phi)
	if ai != 1 {
		return 0, errors.New("Wrong e! gcd(e, phi) should be 1")
	}
	d := int64(extended_euclidean(e, phi, quotients))
	return d, nil
}

func rsa() {
	p := 2003
	q := 4999
	n := int64(p * q)
	phi := int64((p - 1) * (q - 1))
	e := int64(1098325)
	d, err := calculate_decryption_key(e, phi)
	m := int64(2002782)

	if m <= 1 {
		err = errors.New("m should be greater than 1 ")
	}
	if n <= m {
		err = errors.New("m should be smaller than n")
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Message:\t", m)

	ciphertext := exponentiation(m, e, n)
	fmt.Println("Ciphertext:\t", ciphertext)

	decrypted_m := exponentiation(ciphertext, d, n)
	fmt.Println("Decrypted m:\t", decrypted_m)

	if decrypted_m == m {
		fmt.Println("RSA working correctly")
	} else {
		fmt.Println("The decrypted m and the m differ")
	}
}

func main() {
	rsa()
}
