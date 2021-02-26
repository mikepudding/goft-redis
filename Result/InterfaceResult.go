package Result

type InterfaceResult struct {
	Result interface{}
	Error  error
}

func NewInterfaceResut(result interface{}, err error) *InterfaceResult {
	return &InterfaceResult{Result: result, Error: err}
}

func (this *InterfaceResult) Unwrap() interface{} {
	if this.Error != nil {
		panic(this.Error)
	}
	return this.Result
}

func (this *InterfaceResult) Unwrap_Or(str interface{}) interface{} {
	if this.Error != nil {
		return str
	}
	return this.Result
}
