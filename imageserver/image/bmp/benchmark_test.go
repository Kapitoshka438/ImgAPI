package bmp

import (
	"testing"

	"github.com/kapitoshka438/imgapi/imgapi-service"
	imageserver_image_test "github.com/kapitoshka438/imgapi/imgapi-service/image/_test"
	_ "github.com/kapitoshka438/imgapi/imgapi-service/image/jpeg"
	"github.com/kapitoshka438/imgapi/files"
)

func Benchmark(b *testing.B) {
	enc := &Encoder{}
	params := imageserver.Params{}
	for _, tc := range []struct {
		name string
		im   *imageserver.Image
	}{
		{"Small", testdata.Small},
		{"Medium", testdata.Medium},
		{"Large", testdata.Large},
		{"Huge", testdata.Huge},
	} {
		b.Run(tc.name, func(b *testing.B) {
			imageserver_image_test.BenchmarkEncoder(b, enc, tc.im, params)
		})
	}
}
