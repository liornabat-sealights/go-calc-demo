package types; import __sealights__ "github.com/liornabat-sealights/go-calc-demo/__sealights__"

type ResultResponse struct {
	ValueA float64 `json:"valueA"`
	ValueB float64 `json:"valueB"`
	Result float64 `json:"result"`
}

func NewResultResponse() *ResultResponse {__sealights__.TraceFunc("0741f9fa27c7646a69");

	return &ResultResponse{}
}
func (r *ResultResponse) SetValues(valueA, valueB float64) *ResultResponse {__sealights__.TraceFunc("472995e6a32cb3de24");

	r.ValueA = valueA
	r.ValueB = valueB
	return r
}

func (r *ResultResponse) SetResult(value float64) *ResultResponse {__sealights__.TraceFunc("157d0677bbb91fad7e");

	r.Result = value
	return r
}
