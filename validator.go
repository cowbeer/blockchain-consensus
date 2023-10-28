package consensus

var Validators = []Validator{{
	Address: "Validator 1",
	Pledge:  100,
}, {
	Address: "Validator 2",
	Pledge:  200,
}, {
	Address: "Validator 3",
	Pledge:  300,
}, {
	Address: "Validator 4",
	Pledge:  400,
}, {
	Address: "Validator 5",
	Pledge:  500,
}}

type Validator struct {
	Address string
	Pledge  int64
}
