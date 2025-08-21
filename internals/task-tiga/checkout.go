package tasktiga

import (
	"errors"
	"fmt"
)

type Checkout interface {
	Payment(subtotal []int16) (string, error)
}

type Bank struct {
	Name string
}

type Online struct {
	Name string
}

type Fiktif struct {
	ListItem []int16
}

func (b *Bank) Payment(subtotal []int16) (string, error) {
	var total int16 = 0
	for _, v := range subtotal {
		total += v
	}

	return fmt.Sprintf("Pembarayan sebesar %d via bank berhasil diproses", total), nil
}

func (o *Online) Payment(subtotal []int16) (string, error) {
	var total int16 = 0
	for _, v := range subtotal {
		total += v
	}

	return fmt.Sprintf("Pembayaran sebesar %d via online berhasil diproses", total), nil
}

func (f *Fiktif) Payment(subtotal []int16) ([]int16, error) {
	var result []int16

	for _, v := range subtotal {
		if v <= 0 {
			return result, errors.New("tidak dapat melakukan pembayaran <= 0")
		}

		result = append(result, v)
	}

	return result, nil
}
