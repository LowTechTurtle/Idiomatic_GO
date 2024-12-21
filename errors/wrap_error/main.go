package main

import (
	"errors"
	"fmt"
	"os"
)

func FileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in file checker: %w", err)
	}

	f.Close()
	return nil
}

type Status int

type StatusErr struct {
	Status  Status
	Message string
	Err     error
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Unwrap() error {
	return se.Err
}

const (
	InvalidLogin Status = iota + 1
	NotFound
)

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	token, err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
			Err:     err,
		}
	}
	data, err := getData(token, file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
			Err:     err,
		}
	}
	return data, nil
}

func main() {
	err := FileChecker("this_file_doesnt_exists_at_all.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}

}
