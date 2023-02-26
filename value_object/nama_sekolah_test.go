package value_object

import (
	"latihan-unit-test-golang/value_object"
	"testing"

	"github.com/stretchr/testify/assert"
)

var namaSekolahDasar string = "SDN 5 Merpati"

// func TestName(t *testing.T) {
// 	_, err := value_object.NewNamaSekolahDasar(namaSekolahDasar)
// 	assert.Nil(t, err)
// }

func TestName(t *testing.T) {
	if _, err := value_object.NewNamaSekolahDasar(namaSekolahDasar); err != nil {
		return nil, err
	}

	assert.Nil(t, err)
}
