package server; import __sealights__ "github.com/liornabat-sealights/go-calc-demo/__sealights__"

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/liornabat-sealights/go-calc-demo/lib/types"
	"github.com/liornabat-sealights/go-calc-demo/service/pkg/calc"
	"net/http"
	"time"
)

type Server struct {
	echoWebServer *echo.Echo
}

func NewServer() *Server {__sealights__.TraceFunc("d2179752ed768d71af");
	return &Server{}
}

func (s *Server) Start(ctx context.Context) error {__sealights__.TraceFunc("762b3d2e2c9be81646");
	s.echoWebServer = echo.New()
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.GET("/add", func(c echo.Context) error {__sealights__.TraceFunc("77f6c958377182c232");
		return s.handelAdd(c)
	})
	e.GET("/sub", func(c echo.Context) error {__sealights__.TraceFunc("60ad0667f68ca895de");
		return s.handelSub(c)
	})
	e.GET("/mul", func(c echo.Context) error {__sealights__.TraceFunc("8de5c564af72f08169");
		return s.handelMul(c)
	})
	e.GET("/div", func(c echo.Context) error {__sealights__.TraceFunc("17a96042ec3b21b642");
		return s.handelDiv(c)
	})
	s.echoWebServer = e
	errCh := make(chan error, 1)
	go func() {__sealights__.TraceFunc("3e5e1e33ba03f2a55e");
		errCh <- s.echoWebServer.Start("0.0.0.0:10000")
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
		e.Logger.Infof("api server started at port 10000")
		return nil
	case <-time.After(1 * time.Second):
		return nil
	case <-ctx.Done():
		return fmt.Errorf("error strarting api server, %w", ctx.Err())
	}

}

func (s *Server) handelAdd(c echo.Context) error {__sealights__.TraceFunc("43b6d31c6b71643eeb");
	a, b, err := s.parseParams(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	result := types.NewResultResponse().SetValues(a, b).SetResult(calc.Add(a, b))
	return c.JSON(http.StatusOK, result)
}

func (s *Server) handelSub(c echo.Context) error {__sealights__.TraceFunc("c6dc5fb09551d005bc");
	fmt.Println("handelSub")
	a, b, err := s.parseParams(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	result := types.NewResultResponse().SetValues(a, b).SetResult(calc.Subtract(a, b))
	return c.JSON(http.StatusOK, result)
}

func (s *Server) handelMul(c echo.Context) error {__sealights__.TraceFunc("d3a5e38b91abb5992a");
	fmt.Println("handelMul")
	a, b, err := s.parseParams(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	result := types.NewResultResponse().SetValues(a, b).SetResult(calc.Multiply(a, b))
	return c.JSON(http.StatusOK, result)
}

func (s *Server) handelDiv(c echo.Context) error {__sealights__.TraceFunc("fc128a99c7f5ee96b7");
	fmt.Println("handelDiv")
	a, b, err := s.parseParams(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	result := types.NewResultResponse().SetValues(a, b).SetResult(calc.Divide(a, b))
	return c.JSON(http.StatusOK, result)
}

func (s *Server) parseParams(c echo.Context) (float64, float64, error) {__sealights__.TraceFunc("51ac78a305ac3330db");
	var a, b float64
	valueA := c.QueryParam("a")
	valueB := c.QueryParam("b")
	if valueA == "" || valueB == "" {
		return 0, 0, fmt.Errorf("missing params")
	}
	if _, err := fmt.Sscanf(valueA, "%f", &a); err != nil {
		return 0, 0, fmt.Errorf("invalid param a")
	}
	if _, err := fmt.Sscanf(valueB, "%f", &b); err != nil {
		return 0, 0, fmt.Errorf("invalid param b")
	}
	return a, b, nil
}
