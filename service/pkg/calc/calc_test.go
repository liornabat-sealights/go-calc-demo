package calc; import __sealights__ "github.com/liornabat-sealights/go-calc-demo/__sealights__"

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {__sealights__.StartTestFunc("8f45c4ad6e8717d5e0",t);defer func() { __sealights__.EndTestFunc("8f45c4ad6e8717d5e0",t)}();
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
			name: "Add positive numbers",
			args: args{a: 4.5, b: 5.5},
			want: 10,
		},
		{
			name: "Add negative numbers",
			args: args{a: -1.5, b: -2.5},
			want: -4,
		},
		{
			name: "Add positive and negative numbers",
			args: args{a: 5, b: -3},
			want: 2,
		},
		{
			name: "Add zero and number",
			args: args{a: 0, b: 5},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("e00a8ca51ea005f210",t);defer func() { __sealights__.EndTestFunc("e00a8ca51ea005f210",t)}();
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {__sealights__.StartTestFunc("bf8538f0539d96083e",t);defer func() { __sealights__.EndTestFunc("bf8538f0539d96083e",t)}();
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
			name: "Positive and positive",
			args: args{a: 10.0, b: 5.0},
			want: 5.0,
		},
		{
			name: "Positive and negative",
			args: args{a: 10.0, b: -5.0},
			want: 15.0,
		},
		{
			name: "Negative and positive",
			args: args{a: -10.0, b: 5.0},
			want: -15.0,
		},
		{
			name: "Negative and negative",
			args: args{a: -10.0, b: -5.0},
			want: -5.0,
		},
		{
			name: "Zero subtract positive",
			args: args{a: 0.0, b: 5.0},
			want: -5.0,
		},
		{
			name: "Zero subtract negative",
			args: args{a: 0.0, b: -5.0},
			want: 5.0,
		},
		{
			name: "Zero subtract zero",
			args: args{a: 0.0, b: 0.0},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("246099bf07ffff38ee",t);defer func() { __sealights__.EndTestFunc("246099bf07ffff38ee",t)}();
			if got := Subtract(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {__sealights__.StartTestFunc("295b24774435089f36",t);defer func() { __sealights__.EndTestFunc("295b24774435089f36",t)}();
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
			name: "Multiply positive numbers",
			args: args{a: 4.5, b: 2},
			want: 9,
		},
		{
			name: "Multiply negative numbers",
			args: args{a: -3, b: -2},
			want: 6,
		},
		{
			name: "Multiply positive and negative numbers",
			args: args{a: -3, b: 4},
			want: -12,
		},
		{
			name: "Multiply by zero",
			args: args{a: 0, b: 5},
			want: 0,
		},
		{
			name: "Multiply zero by zero",
			args: args{a: 0, b: 0},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("62481d249b1df9712d",t);defer func() { __sealights__.EndTestFunc("62481d249b1df9712d",t)}();
			if got := Multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {__sealights__.StartTestFunc("240a468da383b61b64",t);defer func() { __sealights__.EndTestFunc("240a468da383b61b64",t)}();
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
			name: "Divide positive numbers",
			args: args{a: 10, b: 2},
			want: 5,
		},
		{
			name: "Divide negative numbers",
			args: args{a: -10, b: -2},
			want: 5,
		},
		{
			name: "Divide positive by negative number",
			args: args{a: 10, b: -2},
			want: -5,
		},
		{
			name: "Divide zero by number",
			args: args{a: 0, b: 2},
			want: 0,
		},
		{
			name: "Divide number by zero",
			args: args{a: 10, b: 0},
			want: math.Inf(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {__sealights__.StartTestFunc("618460480ff3cb4959",t);defer func() { __sealights__.EndTestFunc("618460480ff3cb4959",t)}();
			if got := Divide(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
