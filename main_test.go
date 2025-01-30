package main; import __sealights__ "github.com/liornabat-sealights/go-calc-demo/__sealights__"

import (
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/liornabat-sealights/go-calc-demo/lib/types"
	main2 "github.com/liornabat-sealights/go-calc-demo/service/pkg/calc"
	"github.com/liornabat-sealights/go-calc-demo/service/server"
	"os"
	"reflect"
	"testing"
	"time"
)

var restyClient = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

func NewRestRequest() *resty.Request {__sealights__.TraceFunc("8b5c7671ae2dd27123");
	return restyClient.NewRequest()
}

// TestMain is a sample to run an endpoint test
func TestMain(m *testing.M) {__sealights__.StartTestMainFunc("e869415959be0c1886")
	shuttingDown := make(chan struct{})
	go main()
	time.Sleep(2 * time.Second)
	code := __sealights__.RunMainTestFunc("e869415959be0c1886",m)
	close(shuttingDown)
	os.Exit(code)
}

func call(path string, a, b string) (*types.ResultResponse, error) {__sealights__.TraceFunc("68a80157ce08b32856");
	time.Sleep(4 * time.Second)
	result := &types.ResultResponse{}
	resp, err := NewRestRequest().SetQueryParams(map[string]string{
		"a": a,
		"b": b,
	}).SetResult(result).
		Get("http://localhost:10000" + path)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("response: %s", resp.String())
	}
	return result, nil
}
func TestAdd(t *testing.T) {__sealights__.StartTestFunc("05e66b5095ee947f31",t);defer func() { __sealights__.EndTestFunc("05e66b5095ee947f31",t)}();
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "TestAdd",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("dcd09299410d71ab9d",t);defer func() { __sealights__.EndTestFunc("dcd09299410d71ab9d",t)}();
			if got := main2.Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {__sealights__.StartTestFunc("fd7c818533008c1c22",t);defer func() { __sealights__.EndTestFunc("fd7c818533008c1c22",t)}();
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "TestDivide",
			args: args{
				a: 1,
				b: 2,
			},
			want: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("800b285f15d3ecf93e",t);defer func() { __sealights__.EndTestFunc("800b285f15d3ecf93e",t)}();
			if got := main2.Divide(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {__sealights__.StartTestFunc("c098039ece1c1dfab7",t);defer func() { __sealights__.EndTestFunc("c098039ece1c1dfab7",t)}();
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "TestMultiply",
			args: args{
				a: 1,
				b: 2,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("9f6b56eb24fc62aba2",t);defer func() { __sealights__.EndTestFunc("9f6b56eb24fc62aba2",t)}();
			if got := main2.Multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewResultResponse(t *testing.T) {__sealights__.StartTestFunc("05d1b5dbfb4c947611",t);defer func() { __sealights__.EndTestFunc("05d1b5dbfb4c947611",t)}();
	tests := []struct {
		name string
		want *types.ResultResponse
	}{
		{
			name: "TestNewResultResponse",
			want: &types.ResultResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("f20b1695e332ec597d",t);defer func() { __sealights__.EndTestFunc("f20b1695e332ec597d",t)}();
			if got := types.NewResultResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResultResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewServer(t *testing.T) {__sealights__.StartTestFunc("90d11c669d49583bce",t);defer func() { __sealights__.EndTestFunc("90d11c669d49583bce",t)}();
	tests := []struct {
		name string
		want *server.Server
	}{
		{
			name: "TestNewServer",
			want: &server.Server{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("0a4bf4490aa468adb7",t);defer func() { __sealights__.EndTestFunc("0a4bf4490aa468adb7",t)}();
			if got := server.NewServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResultResponse_SetResult(t *testing.T) {__sealights__.StartTestFunc("632c9b80f239b5833c",t);defer func() { __sealights__.EndTestFunc("632c9b80f239b5833c",t)}();
	type fields struct {
		ValueA float64
		ValueB float64
		Result float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *types.ResultResponse
	}{
		{
			name: "TestResultResponse_SetResult",
			fields: fields{
				ValueA: 1,
				ValueB: 2,
				Result: 3,
			},
			args: args{
				value: 4,
			},
			want: &types.ResultResponse{
				ValueA: 1,
				ValueB: 2,
				Result: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("b7803137fbe8a6b3eb",t);defer func() { __sealights__.EndTestFunc("b7803137fbe8a6b3eb",t)}();
			r := &types.ResultResponse{
				ValueA: tt.fields.ValueA,
				ValueB: tt.fields.ValueB,
				Result: tt.fields.Result,
			}
			if got := r.SetResult(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResultResponse_SetValues(t *testing.T) {__sealights__.StartTestFunc("bd3aa86fe2965d4ed6",t);defer func() { __sealights__.EndTestFunc("bd3aa86fe2965d4ed6",t)}();
	type fields struct {
		ValueA float64
		ValueB float64
		Result float64
	}
	type args struct {
		valueA float64
		valueB float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *types.ResultResponse
	}{
		{
			name: "TestResultResponse_SetValues",
			fields: fields{
				ValueA: 1,
				ValueB: 2,
				Result: 3,
			},
			args: args{
				valueA: 4,
			},
			want: &types.ResultResponse{
				ValueA: 4,
				ValueB: 0,
				Result: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("7fc810467cdab9b2ca",t);defer func() { __sealights__.EndTestFunc("7fc810467cdab9b2ca",t)}();
			r := &types.ResultResponse{
				ValueA: tt.fields.ValueA,
				ValueB: tt.fields.ValueB,
				Result: tt.fields.Result,
			}
			if got := r.SetValues(tt.args.valueA, tt.args.valueB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_handel_add(t *testing.T) {__sealights__.StartTestFunc("970a65b27cd916faaf",t);defer func() { __sealights__.EndTestFunc("970a65b27cd916faaf",t)}();
	tests := []struct {
		name       string
		a          string
		b          string
		wantResult *types.ResultResponse
		wantErr    bool
	}{
		{
			name: "TestServer_handelAdd",
			a:    "1",
			b:    "2",
			wantResult: &types.ResultResponse{
				ValueA: 1,
				ValueB: 2,
				Result: 3,
			},
			wantErr: false,
		},
		{
			name: "TestServer_handelAdd- bad request",
			a:    "bad",
			b:    "2",
			wantResult: &types.ResultResponse{
				ValueA: 1,
				ValueB: 2,
				Result: 3,
			},
			wantErr: true,
		},
		{
			name: "TestServer_handelAdd- bad request 2",
			a:    "1",
			b:    "bad",
			wantResult: &types.ResultResponse{
				ValueA: 1,
				ValueB: 2,
				Result: 3,
			},
			wantErr: true,
		},
		{
			name: "TestServer_handelAdd- bad request 3",
			a:    "1",
			b:    "",
			wantResult: &types.ResultResponse{
				ValueA: 1,
				ValueB: 2,
				Result: 3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("788cc8ec7197e700e2",t);defer func() { __sealights__.EndTestFunc("788cc8ec7197e700e2",t)}();
			resp, err := call("/add", tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("handelAdd() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("handelAdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if resp.Result != tt.wantResult.Result {
				t.Errorf("handelAdd() error = %v, wantErr %v", resp.Result, tt.wantResult.Result)
				return
			}
		})
	}
}
func TestServer_handel_sub(t *testing.T) {__sealights__.StartTestFunc("1c0d883902ce5e1e3d",t);defer func() { __sealights__.EndTestFunc("1c0d883902ce5e1e3d",t)}();
	tests := []struct {
		name       string
		a          string
		b          string
		wantResult *types.ResultResponse
		wantErr    bool
	}{
		{
			name: "TestServer_handelsub",
			a:    "1",
			b:    "2",
			wantResult: &types.ResultResponse{
				ValueA: 1,
				ValueB: 2,
				Result: -1,
			},
			wantErr: false,
		},
		{
			name: "TestServer_handelsub- bad request",
			a:    "bad",
			b:    "2",
			wantResult: &types.ResultResponse{
				ValueA: 1,
				ValueB: 2,
				Result: 3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("767d8da52e948ca5db",t);defer func() { __sealights__.EndTestFunc("767d8da52e948ca5db",t)}();
			resp, err := call("/sub", tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("handelAdd() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("handelAdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if resp.Result != tt.wantResult.Result {
				t.Errorf("handelAdd() error = %v, wantErr %v", resp.Result, tt.wantResult.Result)
				return
			}
		})
	}
}
func TestServer_handel_mul(t *testing.T) {__sealights__.StartTestFunc("629400591a996b1bd0",t);defer func() { __sealights__.EndTestFunc("629400591a996b1bd0",t)}();
	tests := []struct {
		name       string
		a          string
		b          string
		wantResult *types.ResultResponse
		wantErr    bool
	}{
		{
			name: "TestServer_handelmul",
			a:    "1",
			b:    "2",
			wantResult: &types.ResultResponse{
				ValueA: 1,
				ValueB: 2,
				Result: 2,
			},
			wantErr: false,
		},
		{
			name: "TestServer_handelmul- bad request",
			a:    "bad",
			b:    "2",
			wantResult: &types.ResultResponse{
				ValueA: 1,
				ValueB: 2,
				Result: 3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("4d34211a152e554eb0",t);defer func() { __sealights__.EndTestFunc("4d34211a152e554eb0",t)}();
			resp, err := call("/mul", tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("handelAdd() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("handelAdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if resp.Result != tt.wantResult.Result {
				t.Errorf("handelAdd() error = %v, wantErr %v", resp.Result, tt.wantResult.Result)
				return
			}
		})
	}
}
func TestServer_handel_div(t *testing.T) {__sealights__.StartTestFunc("f110d8fd40451fd579",t);defer func() { __sealights__.EndTestFunc("f110d8fd40451fd579",t)}();
	tests := []struct {
		name       string
		a          string
		b          string
		wantResult *types.ResultResponse
		wantErr    bool
	}{
		{
			name: "TestServer_handeldiv",
			a:    "1",
			b:    "2",
			wantResult: &types.ResultResponse{
				ValueA: 1,
				ValueB: 2,
				Result: 0.5,
			},
			wantErr: false,
		},
		{
			name: "TestServer_handeldiv- bad request",
			a:    "bad",
			b:    "2",
			wantResult: &types.ResultResponse{
				ValueA: 1,
				ValueB: 2,
				Result: 3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("199dc87ee20e349a6f",t);defer func() { __sealights__.EndTestFunc("199dc87ee20e349a6f",t)}();
			resp, err := call("/div", tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("handelAdd() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("handelAdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if resp.Result != tt.wantResult.Result {
				t.Errorf("handelAdd() error = %v, wantErr %v", resp.Result, tt.wantResult.Result)
				return
			}
		})
	}
}

//
//	func TestServer_handelDiv(t *testing.T) {
//		type fields struct {
//			echoWebServer *echo.Echo
//		}
//		type args struct {
//			c echo.Context
//		}
//		tests := []struct {
//			name    string
//			fields  fields
//			args    args
//			wantErr bool
//		}{
//			// TODO: Add test cases.
//		}
//		for _, tt := range tests {
//			t.Run(tt.name, func(t *testing.T) {
//				s := &Server{
//					echoWebServer: tt.fields.echoWebServer,
//				}
//				if err := s.handelDiv(tt.args.c); (err != nil) != tt.wantErr {
//					t.Errorf("handelDiv() error = %v, wantErr %v", err, tt.wantErr)
//				}
//			})
//		}
//	}
//
//	func TestServer_handelMul(t *testing.T) {
//		type fields struct {
//			echoWebServer *echo.Echo
//		}
//		type args struct {
//			c echo.Context
//		}
//		tests := []struct {
//			name    string
//			fields  fields
//			args    args
//			wantErr bool
//		}{
//			// TODO: Add test cases.
//		}
//		for _, tt := range tests {
//			t.Run(tt.name, func(t *testing.T) {
//				s := &Server{
//					echoWebServer: tt.fields.echoWebServer,
//				}
//				if err := s.handelMul(tt.args.c); (err != nil) != tt.wantErr {
//					t.Errorf("handelMul() error = %v, wantErr %v", err, tt.wantErr)
//				}
//			})
//		}
//	}
//
//	func TestServer_handelSub(t *testing.T) {
//		type fields struct {
//			echoWebServer *echo.Echo
//		}
//		type args struct {
//			c echo.Context
//		}
//		tests := []struct {
//			name    string
//			fields  fields
//			args    args
//			wantErr bool
//		}{
//			// TODO: Add test cases.
//		}
//		for _, tt := range tests {
//			t.Run(tt.name, func(t *testing.T) {
//				s := &Server{
//					echoWebServer: tt.fields.echoWebServer,
//				}
//				if err := s.handelSub(tt.args.c); (err != nil) != tt.wantErr {
//					t.Errorf("handelSub() error = %v, wantErr %v", err, tt.wantErr)
//				}
//			})
//		}
//	}

func TestSubtract(t *testing.T) {__sealights__.StartTestFunc("e89793a2496229b1ad",t);defer func() { __sealights__.EndTestFunc("e89793a2496229b1ad",t)}();
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "TestSubtract",
			args: args{
				a: 10,
				b: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("087b72f1f091a55de4",t);defer func() { __sealights__.EndTestFunc("087b72f1f091a55de4",t)}();
			if got := main2.Subtract(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}
