package main

type Car interface {
	StartEngine()
	Move()
	Construct()
	GetHorsePower() int
	GetPrice() int
	GetModelName() string
}
