package main

import (
	"fmt"
	"math"
)

// type Play map[string]map[string]string
type Play struct {
	Name string
	Kind string
}
type Plays map[string]Play

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

type Statement struct {
	Name   string
	Amount float64
	Seats  int
	Credit float64
}

func CalculateAmount(plays Plays, perf Performance) float64 {
	play := plays[perf.PlayID]
	amount := 0.0

	switch play.Kind {
	case "tragedy":
		amount = 40000
		if perf.Audience > 30 {
			amount += 1000 * (float64(perf.Audience - 30))
		}
	case "comedy":
		amount = 30000
		if perf.Audience > 20 {
			amount += 10000 + 500*(float64(perf.Audience-20))
		}
		amount += 300 * float64(perf.Audience)
	default:
		panic(fmt.Sprintf("unknow type: %s", play.Kind)) //remove panic
	}
	return amount
}

func CalculateCredit(plays Plays, perf Performance) float64 {
	play := plays[perf.PlayID]
	volumeCredits := math.Max(float64(perf.Audience-30), 0)
	// add extra credit for every ten comedy attendees
	if "comedy" == play.Kind {
		volumeCredits += math.Floor(float64(perf.Audience / 5))
	}
	return volumeCredits
}

func CalculateTotalAmount(orders []Statement) float64 {
	total := 0.0
	for _, order := range orders {
		total += order.Amount
	}
	return total
}

func CalculateTotalCredit(orders []Statement) float64 {
	credit := 0.0
	for _, order := range orders {
		credit += order.Credit
	}
	return credit
}

func statement(invoice Invoice, plays Plays) string {
	var statements []Statement

	for _, perf := range invoice.Performances {
		resultStruct := Statement{
			Name:   plays[perf.PlayID].Name,
			Amount: CalculateAmount(plays, perf) / 100,
			Seats:  perf.Audience,
			Credit: CalculateCredit(plays, perf),
		}
		statements = append(statements, resultStruct)
	}

	return RenderHTMLOutput(invoice, statements)
}

func RenderTextOutput(invoice Invoice, orders []Statement) string {
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)
	for _, order := range orders {
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", order.Name, order.Amount, order.Seats)
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", CalculateTotalAmount(orders))
	result += fmt.Sprintf("you earned %.0f credits\n", CalculateTotalCredit(orders))
	return result
}

func RenderHTMLOutput(invoice Invoice, orders []Statement) string {
	result := fmt.Sprintf("<h1>Statement for %s</h1>\n<table>\n<tr><th>play</th><th>seats</th><th>cost</th></tr>\n", invoice.Customer)
	for _, order := range orders {
		result += fmt.Sprintf("<tr><td>%s</td><td>%d</td><td>$%.2f</td></tr>\n", order.Name, order.Seats, order.Amount)
	}
	result += fmt.Sprintf("</table>\n<p>Amount owed is <em>$%.2f</em></p>\n", CalculateTotalAmount(orders))
	result += fmt.Sprintf("<p>you earned <em>%.0f</em> credits</p>\n", CalculateTotalCredit(orders))
	return result
}

func main() {
	inv := Invoice{
		Customer: "Bigco",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 55},
			{PlayID: "as-like", Audience: 35},
			{PlayID: "othello", Audience: 40},
		}}
	plays := Plays{
		"hamlet":  {Name: "Hamlet", Kind: "tragedy"},
		"as-like": {Name: "As You Like It", Kind: "comedy"},
		"othello": {Name: "Othello", Kind: "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
