package kernel

type Major struct {
	Name  string
	Value float64
}

func NewMajor(name string, initial_value float64) Major {
	major := Major{
		name,
		initial_value,
	}
	return major
}
