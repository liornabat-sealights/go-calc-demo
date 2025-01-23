package main

import (
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/liornabat-sealights/go-calc-demo/lib/types"
	main2 "github.com/liornabat-sealights/go-calc-demo/service/pkg/calc"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var restyClient = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

func NewRestRequest() *resty.Request {

	return restyClient.NewRequest()
}

func call(path string, a, b string) (*types.ResultResponse, error) {

	time.Sleep(4 * time.Second)
	result := &types.ResultResponse{}
	resp, err := NewRestRequest().SetQueryParams(map[string]string{
		"a": a,
		"b": b,
	}).SetResult(result).
		Get("http://localhost:8080" + path)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("response: %s", resp.String())
	}
	return result, nil
}

var _ = BeforeSuite(func() {
	// Start the server (assuming main() starts your server)
	go main()
	time.Sleep(2 * time.Second) // Wait for server to start
})

var _ = Describe("Calculator Service", func() {
	var _ = BeforeEach(
		func() {
			fmt.Println("Before each test in the describe")

		})
	Describe("Basic Calculator Operations", func() {
		Context("Add Operation", func() {
			It("should correctly add two numbers", func() {
				result := main2.Add(1, 2)
				Expect(result).To(Equal(float64(3)))
			})
		})

		Context("Subtract Operation", func() {
			It("should correctly subtract two numbers", func() {
				result := main2.Subtract(10, 5)
				Expect(result).To(Equal(float64(5)))
			})
		})

		Context("Multiply Operation", func() {
			It("should correctly multiply two numbers", func() {
				result := main2.Multiply(1, 2)
				Expect(result).To(Equal(float64(2)))
			})
		})

		Context("Divide Operation", func() {
			It("should correctly divide two numbers", func() {
				result := main2.Divide(1, 2)
				Expect(result).To(Equal(float64(0.5)))
			})
		})
	})

	Describe("ResultResponse Type", func() {
		Context("NewResultResponse", func() {
			It("should create a new empty ResultResponse", func() {
				response := types.NewResultResponse()
				Expect(response).To(Equal(&types.ResultResponse{}))
			})
		})

		Context("SetResult", func() {
			It("should set the result value correctly", func() {
				r := &types.ResultResponse{
					ValueA: 1,
					ValueB: 2,
					Result: 3,
				}
				result := r.SetResult(4)
				Expect(result).To(Equal(&types.ResultResponse{
					ValueA: 1,
					ValueB: 2,
					Result: 4,
				}))
			})
		})

		Context("SetValues", func() {
			It("should set the input values correctly", func() {
				r := &types.ResultResponse{
					ValueA: 1,
					ValueB: 2,
					Result: 3,
				}
				result := r.SetValues(4, 0)
				Expect(result).To(Equal(&types.ResultResponse{
					ValueA: 4,
					ValueB: 0,
					Result: 3,
				}))
			})
		})
	})

	Describe("Server API Endpoints", func() {
		Context("Addition Endpoint", func() {
			It("should correctly handle valid addition requests", func() {
				resp, err := call("/add", "1", "2")
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Result).To(Equal(float64(3)))
			})

			It("should handle invalid input gracefully", func() {
				_, err := call("/add", "bad", "2")
				Expect(err).To(HaveOccurred())

				_, err = call("/add", "1", "bad")
				Expect(err).To(HaveOccurred())

				_, err = call("/add", "1", "")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("Subtraction Endpoint", func() {
			It("should correctly handle valid subtraction requests", func() {
				resp, err := call("/sub", "1", "2")
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Result).To(Equal(float64(-1)))
			})

			It("should handle invalid input gracefully", func() {
				_, err := call("/sub", "bad", "2")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("Multiplication Endpoint", func() {
			It("should correctly handle valid multiplication requests", func() {
				resp, err := call("/mul", "1", "2")
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Result).To(Equal(float64(2)))
			})

			It("should handle invalid input gracefully", func() {
				_, err := call("/mul", "bad", "2")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("Division Endpoint", func() {
			It("should correctly handle valid division requests", func() {
				resp, err := call("/div", "1", "2")
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Result).To(Equal(float64(0.5)))
			})

			It("should handle invalid input gracefully", func() {
				_, err := call("/div", "bad", "2")
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
