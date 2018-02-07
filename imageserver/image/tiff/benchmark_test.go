package tiff

import (
	"testing"

	"github.com/kapitoshka438/imgapi/imgapi-service"
	imageserver_image_test "github.com/kapitoshka438/imgapi/imgapi-service/image/_test"
	_ "github.com/kapitoshka438/imgapi/imgapi-service/image/jpeg"
	"github.com/kapitoshka438/imgapi/files"
)

func Benchmark(b *testing.B) {
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
			imageserver_image_test.BenchmarkEncoder(b, &Encoder{}, tc.im, imageserver.Params{})
		})
	}
}
