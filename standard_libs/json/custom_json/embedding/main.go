package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID          string    `json:"id"`
	Items       []Item    `json:"items"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string    `json:"customer_id"`
}

func (o Order) MarshalJSON() ([]byte, error) {
	type Dup Order

// There are 2 reasons in using Dup.
// The first is in golang, inheritance doesnt exist, so
// when we use Dup, it will strip the methods of Order
// => when we call Marshal on Dup, it will not call Marshal
// again => no infinite loop
//
// The second reason is we embed a struct, if the containing
// struct have a field that is named identically to a field
// in an embedded struct, Marshal and Unmarshal json will
// ignore the field in the embedded struct

	tmp := struct {
		DateOrdered string `json:"date_ordered"`
		Dup
	}{
		// type casting, (Dup) is of type Dup, and the Dup
		// on the left is an embedded struct, and we are
		// assigning o to it
		Dup: (Dup)(o),
	}
	// since o.DateOrdered is not the format we desired and
	// its ignored, but it has the needed data, we assign 
	// the desired format into tmp.DateOrdered
	tmp.DateOrdered = o.DateOrdered.Format(time.RFC822Z)
	b, err := json.Marshal(tmp)
	return b, err
}

func (o *Order) UnmarshalJSON(b []byte) error {
	type Dup Order

	tmp := struct {
		DateOrdered string `json:"date_ordered"`
		*Dup
	}{
		Dup: (*Dup)(o),
	}

// get the data from the json encoded data
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

// since the o.DateOrdered is what we need outside of the func
// and its ignored, so we parse it and manually set it back
	o.DateOrdered, err = time.Parse(time.RFC822Z, tmp.DateOrdered)
	if err != nil {
		return err
	}
	return nil
}

// after these 2 function, we realized that we only need to
// modify the handling function of json instead of modifying
// the data itself. Its considered idiomatic because
// data for business logic does not depend on the json, but
// rather the handling function depend on the json
//
// when we change the format of the json, we only need to
// change the handling function, not the entire data and 
// related business logic

func main() {
	data := `
	{
		"id": "12345",
		"items": [
			{
				"id": "xyz123",
				"name": "Thing 1"
			},
			{
				"id": "abc789",
				"name": "Thing 2"
			}
		],
		"date_ordered": "01 May 20 13:01 +0000",
		"customer_id": "3"
	}`

	var o Order
	err := json.Unmarshal([]byte(data), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
	fmt.Println(o.DateOrdered.Month())
	out, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}

