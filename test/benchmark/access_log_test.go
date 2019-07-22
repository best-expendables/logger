package benchmark

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"bitbucket.org/snapmartinc/logger"
)

type (
	Request struct {
		Header http.Header
		Method string
		URL    string
		Body   string
	}

	Response struct {
		Headers    http.Header
		StatusCode int
		Body       string
	}
)

var content = logger.Fields{
	"request": Request{
		Header: http.Header{
			"User-Agent":           []string{"Go-http-client/1.1"},
			"X-Forwarded-For":      []string{"127.0.0.1"},
			"X-Lel-Context-Id":     []string{"tms-api-644f77b5c9-tnbpm/V5FWILixRC-9430825"},
			"X-Lel-User-Email":     []string{"package.lel@lazada.com"},
			"X-Lel-User-Id":        []string{"9ab652b0-1d20-4b97-868a-2776760e53d0"},
			"X-Lel-User-Name":      []string{"package.lel"},
			"X-Lel-User-Platforms": []string{"LAZADA_ID", "LAZADA_VN", "LAZADA_PH", "LAZADA_TH", "LAZADA_MY", "ALIBABA", "LAZADA_SG"},
			"X-Lel-User-Roles":     []string{"PlatformTransportation"},
		},
		URL:    "/api/addresses/administrative-levels/R80000448/parents",
		Method: "GET",
	},
	"response": Response{
		Headers: http.Header{
			"Content-Type":     []string{"application/json; charset=utf-8"},
			"X-Lel-Context-Id": []string{"tms-api-644f77b5c9-tnbpm/V5FWILixRC-9430825"},
		},
		StatusCode: http.StatusOK,
		Body:       `{"data":[{"id":"R6941055","name":"Quận Thủ Đức","nameLocal":"","parentId":"R1973756","isActive":true,"geoCoordinates":{"latitude":10.852588,"longitude":106.7558383}},{"id":"R1973756","name":"Hồ Chí Minh","nameLocal":"","parentId":"R49915","isActive":true,"geoCoordinates":{"latitude":10.7758439,"longitude":106.7017555}},{"id":"R49915","name":"Viet Nam","nameLocal":"","parentId":null,"isActive":true,"geoCoordinates":null}]}`,
	},
}

func BenchmarkAccessLog(b *testing.B) {
	out := new(bytes.Buffer)

	factory := logger.NewLoggerFactory(
		logger.DebugLevel,
		logger.SetOut(out),
	)

	e := factory.Logger(context.Background())

	b.Run("Single", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e.WithFields(content).Info()
		}
		out.Reset()
	})

	b.Run("Parallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				e.WithFields(content).Info()
			}
		})
	})
}
