package main

import (
	"bytes"
	"strings"
	"testing"

	pb "github.com/dreamjz/golang-notes/protocol-buffers/address-book/protobuffers"
)

func TestWritePersonWritesPerson(t *testing.T) {
	buf := new(bytes.Buffer)
	// [START populate_proto]
	p := pb.Person{
		Name:  "kesa",
		Id:    1234,
		Email: "kesa@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_MOBILE},
		},
	}
	// [END populate_proto]
	writePerson(buf, &p)
	got := buf.String()
	want := `Person ID: 1234
 Name: kesa
 E-mail address: kesa@example.com
 Mobile phone #: 555-4321
`
	if got != want {
		t.Errorf("writePerson(%s) =>\n\t%q, want %q", p.String(), got, want)
	}
}

func TestListPeopleWritesList(t *testing.T) {
	buf := new(bytes.Buffer)
	in := pb.AddressBook{
		People: []*pb.Person{
			{
				Name:  "A1",
				Id:    101,
				Email: "E1",
			},
			{
				Name:  "B2",
				Id:    102,
				Email: "E2",
			},
			{
				Name:  "C3",
				Id:    103,
				Email: "E3",
			},
		},
	}
	listPeople(buf, &in)
	want := strings.Split(`Person ID: 101
 Name: A1
 E-mail address: E1
Person ID: 102
 Name: B2
 E-mail address: E2
Person ID: 103
 Name: C3
 E-mail address: E3
`, "\n")
	got := strings.Split(buf.String(), "\n")
	if len(got) != len(want) {
		t.Errorf(
			"listpeople(%s) => \n\t%q has %d lines, want %d",
			in.String(),
			buf.String(),
			len(got),
			len(want),
		)
	}
	lines := len(got)
	if lines > len(want) {
		lines = len(want)
	}
	for i := 0; i < lines; i++ {
		if got[i] != want[i] {
			t.Errorf(
				"listPeople(%s) =>\n\tline %d %q, want %q\n\t%v\n\t%v",
				in.String(),
				i,
				got[i],
				want[i],
				[]byte(got[i]),
				[]byte(want[i]))
		}
	}
}
