package main

import (
	"testing"
)

type Profile struct {
	Name    string
	Surname string
	City    string
}

func BenchmarkAssign(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p1 := Profile{
			Name:    "Ivan",
			Surname: "Sukach",
			City:    "Wroclaw",
		}
		p2 := Profile{
			Name:    "Denis",
			Surname: "Sukach",
			City:    "Warsaw",
		}
		var changed bool
		if p1.Name != p2.Name || p1.Surname != p2.Surname || p1.City != p2.City {
			p1.Name = p2.Name
			p1.Surname = p2.Surname
			p1.City = p2.City
			changed = true
		}
		changed = !changed
	}
}

func BenchmarkCompare(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p1 := Profile{
			Name:    "Ivan",
			Surname: "Sukach",
			City:    "Wroclaw",
		}
		p2 := Profile{
			Name:    "Denis",
			Surname: "Sukach",
			City:    "Warsaw",
		}
		var changed bool
		if p1.Name != p2.Name {
			p1.Name = p2.Name
			changed = true
		}
		if p1.Surname != p2.Surname {
			p1.Surname = p2.Surname
			changed = true
		}
		if p1.City != p2.City {
			p1.City = p2.City
			changed = true
		}
		changed = !changed
	}
}

func BenchmarkAssignTwoIdentical(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p1 := Profile{
			Name:    "Ivan",
			Surname: "Sukach",
			City:    "Wroclaw",
		}
		p2 := Profile{
			Name:    "Denis",
			Surname: "Sukach",
			City:    "Wroclaw",
		}
		var changed bool
		if p1.Name != p2.Name || p1.Surname != p2.Surname || p1.City != p2.City {
			p1.Name = p2.Name
			p1.Surname = p2.Surname
			p1.City = p2.City
			changed = true
		}
		changed = !changed
	}
}

func BenchmarkCompareTwoIdentical(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p1 := Profile{
			Name:    "Ivan",
			Surname: "Sukach",
			City:    "Wroclaw",
		}
		p2 := Profile{
			Name:    "Denis",
			Surname: "Sukach",
			City:    "Wroclaw",
		}
		var changed bool
		if p1.Name != p2.Name {
			p1.Name = p2.Name
			changed = true
		}
		if p1.Surname != p2.Surname {
			p1.Surname = p2.Surname
			changed = true
		}
		if p1.City != p2.City {
			p1.City = p2.City
			changed = true
		}
		changed = !changed
	}
}

func BenchmarkAssignTwoIdenticalWithoutBooleanFlag(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p1 := Profile{
			Name:    "Ivan",
			Surname: "Sukach",
			City:    "Wroclaw",
		}
		p2 := Profile{
			Name:    "Denis",
			Surname: "Sukach",
			City:    "Wroclaw",
		}
		if p1.Name != p2.Name || p1.Surname != p2.Surname || p1.City != p2.City {
			p1.Name = p2.Name
			p1.Surname = p2.Surname
			p1.City = p2.City
		}
	}
}

func BenchmarkCompareTwoIdenticalWithoutBooleanFlag(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p1 := Profile{
			Name:    "Ivan",
			Surname: "Sukach",
			City:    "Wroclaw",
		}
		p2 := Profile{
			Name:    "Denis",
			Surname: "Sukach",
			City:    "Wroclaw",
		}
		if p1.Name != p2.Name {
			p1.Name = p2.Name
		}
		if p1.Surname != p2.Surname {
			p1.Surname = p2.Surname
		}
		if p1.City != p2.City {
			p1.City = p2.City
		}
	}
}

func BenchmarkAssignLong(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p1 := Profile{
			Name:    "IvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvan",
			Surname: "SukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukach",
			City:    "WroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclaw",
		}
		p2 := Profile{
			Name:    "DenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenis",
			Surname: "SukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukach",
			City:    "WroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWarsawWroclawWroclawWroclawWroclaw",
		}
		var changed bool
		if p1.Name != p2.Name || p1.Surname != p2.Surname || p1.City != p2.City {
			p1.Name = p2.Name
			p1.Surname = p2.Surname
			p1.City = p2.City
			changed = true
		}
		changed = !changed
	}
}

func BenchmarkCompareLong(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p1 := Profile{
			Name:    "IvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvan",
			Surname: "SukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukach",
			City:    "WroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclaw",
		}
		p2 := Profile{
			Name:    "DenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenis",
			Surname: "SukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukach",
			City:    "WroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWarsawWroclawWroclawWroclawWroclaw",
		}
		var changed bool
		if p1.Name != p2.Name {
			p1.Name = p2.Name
			changed = true
		}
		if p1.Surname != p2.Surname {
			p1.Surname = p2.Surname
			changed = true
		}
		if p1.City != p2.City {
			p1.City = p2.City
			changed = true
		}
		changed = !changed
	}
}

func BenchmarkAssignLongTwoIdentical(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p1 := Profile{
			Name:    "IvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvan",
			Surname: "SukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukach",
			City:    "WroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclaw",
		}
		p2 := Profile{
			Name:    "DenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenis",
			Surname: "SukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukach",
			City:    "WroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclaw",
		}
		var changed bool
		if p1.Name != p2.Name || p1.Surname != p2.Surname || p1.City != p2.City {
			p1.Name = p2.Name
			p1.Surname = p2.Surname
			p1.City = p2.City
			changed = true
		}
		changed = !changed
	}
}

func BenchmarkCompareLongTwoIdentical(b *testing.B) {

	for i := 0; i < b.N; i++ {
		p1 := Profile{
			Name:    "IvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvan",
			Surname: "SukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukach",
			City:    "WroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclaw",
		}
		p2 := Profile{
			Name:    "DenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenisDenis",
			Surname: "SukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukachSukach",
			City:    "WroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclawWroclaw",
		}
		var changed bool
		if p1.Name != p2.Name {
			p1.Name = p2.Name
			changed = true
		}
		if p1.Surname != p2.Surname {
			p1.Surname = p2.Surname
			changed = true
		}
		if p1.City != p2.City {
			p1.City = p2.City
			changed = true
		}
		changed = !changed
	}
}
