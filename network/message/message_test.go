package message

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	dict := map[string]uint16{
		"test.test.test":  100,
		"test.test.test1": 101,
		"test.test.test2": 102,
		"test.test.test3": 103,
	}
	SetDict(dict)
	m1 := &Message{
		Type:       Request,
		ID:         100,
		Route:      "test.test.test",
		Data:       []byte(`hello world`),
		compressed: true,
	}
	em1, err := m1.Encode()
	if err != nil {
		t.Fail()
	}
	dm1, err := Decode(em1)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(m1, dm1) {
		t.Fail()
	}

	m2 := &Message{
		Type:  Request,
		ID:    100,
		Route: "test.test.test4",
		Data:  []byte(`hello world`),
	}
	em2, err := m2.Encode()
	if err != nil {
		t.Fail()
	}
	dm2, err := Decode(em2)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(m2, dm2) {
		t.Fail()
	}

	m3 := &Message{
		Type:  Response,
		ID:    100,
		Route: "test.test.test0",
		Data:  []byte(`hello world`),
	}
	em3, err := m3.Encode()
	if err != nil {
		t.Fail()
	}
	dm3, err := Decode(em3)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(m3, dm3) {
		t.Fail()
	}

	m4 := &Message{
		Type:       Response,
		ID:         100,
		Route:      "test.test.test1",
		Data:       []byte(`hello world`),
		compressed: true,
	}
	em4, err := m4.Encode()
	if err != nil {
		t.Fail()
	}
	dm4, err := Decode(em4)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(m4, dm4) {
		t.Fail()
	}

	m5 := &Message{
		Type:       Notify,
		Route:      "test.test.test",
		Data:       []byte(`hello world`),
		compressed: true,
	}
	em5, err := m5.Encode()
	if err != nil {
		t.Fail()
	}
	dm5, err := Decode(em5)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(m5, dm5) {
		t.Fail()
	}

	m6 := &Message{
		Type:  Notify,
		Route: "test.test.test20",
		Data:  []byte(`hello world`),
	}
	em6, err := m6.Encode()
	if err != nil {
		t.Fail()
	}
	dm6, err := Decode(em6)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(m6, dm6) {
		t.Fail()
	}

	m7 := &Message{
		Type:  Push,
		Route: "test.test.test9",
		Data:  []byte(`hello world`),
	}
	em7, err := m7.Encode()
	if err != nil {
		t.Fail()
	}
	dm7, err := Decode(em7)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(m7, dm7) {
		t.Fail()
	}

	m8 := &Message{
		Type:       Push,
		Route:      "test.test.test3",
		Data:       []byte(`hello world`),
		compressed: true,
	}
	em8, err := m8.Encode()
	if err != nil {
		t.Fail()
	}
	dm8, err := Decode(em8)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(m8, dm8) {
		t.Fail()
	}
}
