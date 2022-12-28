package encapsulation

type Validation struct {
	pattern string
}

func NewValidation() *Validation {
	pattern := "dfasdjfoiweif093ru" // set pattern in constructor
	return &Validation{
		pattern: pattern,
	}
}

func (v *Validation) GetPattern() string {
	return v.pattern
}
