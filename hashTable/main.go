package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

type Entry struct {
	Key       string
	Value     string
	isDeleted bool
}

type HashTable struct {
	table []*Entry
	size  int
}

var (
	ErrNotFound        = errors.New("not found")
	ErrEmptyValue      = errors.New("key or value is empty")
	ErrDuplicationHash = errors.New("duplicate key")
)

func preparationForTheSearch(k string) (string, error) {
	if k == "" {
		return "", ErrEmptyValue
	}

	H := sha256.New()

	H.Write([]byte(k))

	desiredHash := string(H.Sum(nil))

	return desiredHash, nil
}

func (ht *HashTable) Add(k, v string) error {
	if k == "" || v == "" {
		return ErrEmptyValue
	}

	H := sha256.New()

	H.Write([]byte(k))

	entry := &Entry{
		Key:   string(H.Sum(nil)),
		Value: v,
	}

	if ht.size == 0 {
		ht.table = append(ht.table, entry)
		ht.size++
		return nil
	}

	var hashMatch bool
	for _, val := range ht.table {
		if val.Key == entry.Key && !val.isDeleted {
			hashMatch = true
		}
	}

	if hashMatch {
		return ErrDuplicationHash
	}

	ht.table = append(ht.table, entry)
	ht.size++

	return nil
}

func (ht *HashTable) Get(k string) (string, error) {

	requestedHash, err := preparationForTheSearch(k)
	if err != nil {
		return "", err
	}

	for _, val := range ht.table {
		if val.Key == requestedHash && !val.isDeleted {
			return val.Value, nil
		}
	}

	return "", ErrNotFound
}

func (ht *HashTable) Delete(k string) error {

	preparedForRemoval, err := preparationForTheSearch(k)
	if err != nil {
		return err
	}

	for _, val := range ht.table {
		if val.Key == preparedForRemoval && !val.isDeleted {
			val.isDeleted = true
			ht.size--
			return nil
		}
	}

	return ErrNotFound
}

func main() {

	hashT := new(HashTable)

	fmt.Println(hashT.Add("Dima", "123"))
	fmt.Println(hashT.size)
	fmt.Println(hashT.Add("Rodion", "12345"))
	fmt.Println(hashT.Add("Alexander", "12300"))
	fmt.Println(hashT.Add("Alexey", "100345"))
	fmt.Println(hashT.size)

	fmt.Println(hashT.Add("Dima", "123"))
	fmt.Println(hashT.Add("", "123"))
	fmt.Println(hashT.Add("Dima", ""))

	fmt.Println(hashT.Get("Dima"))
	fmt.Println(hashT.Get("Valera"))
	fmt.Println(hashT.Get(""))

	fmt.Println(hashT.Delete("Valera"))
	fmt.Println(hashT.Delete("Dima"))
	fmt.Println(hashT.size)
}
