package main

import "fmt"

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

/*
func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	token, err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{Status: InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}

	data, err := getdata(token, file)
	if err != nil {
		return nil, StatusErr{Status: NotFound,
			Message: fmt.Sprintf("file %s not found", file)}
	}

	return data, nil
}
*/

// error is an interface, when return an interface of type genErr, the interface is never nil
// (interface only nil if type and value of the interface is nil)
func GenerateErrorBroken(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{Status: NotFound}
	}
	return genErr
}

func GenerateErrorOK(flag bool) error {
	var genErr error
	if flag {
		genErr = StatusErr{Status: NotFound}
	}
	return genErr
}

func main() {
	err := GenerateErrorBroken(true)
	fmt.Println("GenerateErrorBroken(true) returns non-nil error:", err != nil)
	err = GenerateErrorBroken(false)
	fmt.Println("GenerateErrorBroken(false) returns non-nil error:", err != nil)

	err = GenerateErrorOK(true)
	fmt.Println("GenerateErrorOK(true) returns non-nil error:", err != nil)
	err = GenerateErrorOK(false)
	fmt.Println("GenerateErrorOK(false) returns non-nil error:", err != nil)
}
