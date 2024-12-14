package main

import (
	"errors"
	"fmt"
	"net/http"
)

// just normal func implement based on concrete type
func LogOutput(message string) {
	fmt.Println(message)
}

// struct for storing key, value
type SimpleDataStore struct {
	userData map[string]string
}

// use the ok idiom to get name from id
func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// factory func for SDS
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

// interface for data storing
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// interface for logger
type Logger interface {
	Log(message string)
}

// LA and Log meet the business logic, but have no idea that it does
type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

type Logic interface {
	SayHello(userID string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}
func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}
