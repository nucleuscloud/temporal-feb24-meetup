package transformers

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/benthosdev/benthos/v4/public/bloblang"
)

var defaultSSNLength = int64(10)

func init() {
	spec := bloblang.NewPluginSpec()

	err := bloblang.RegisterFunctionV2("generate_ssn", spec, func(args *bloblang.ParsedParams) (bloblang.Function, error) {
		return func() (any, error) {
			val, err := generateRandomSSN()

			if err != nil {
				return false, fmt.Errorf("unable to generate random ssn")
			}
			return val, nil
		}, nil
	})
	if err != nil {
		panic(err)
	}
}

/* Generates a random social security number in the format XXX-XX-XXXX */
func generateRandomSSN() (string, error) {
	val, err := generateRandomInt64InLengthRange(defaultSSNLength, defaultSSNLength)
	if err != nil {
		return "", err
	}

	returnVal := fmt.Sprintf("%03d-%02d-%04d", val/10000000, (val/10000)%100, val%10000)

	return returnVal, nil
}

func generateRandomInt64InLengthRange(min, max int64) (int64, error) {
	if min > max {
		min, max = max, min
	}

	// Ensure the length doesn't exceed the limit for int64
	if min > 19 || max > 19 {
		return 0, fmt.Errorf("length is too large")
	}

	val, err := generateRandomInt64InValueRange(min, max)
	if err != nil {
		return 0, fmt.Errorf("unable to generate a value in the range provided")
	}

	res, err := generateRandomInt64FixedLength(val)
	if err != nil {
		return 0, fmt.Errorf("unable to generate a value in the range provided")
	}

	return res, nil
}
func generateRandomInt64FixedLength(l int64) (int64, error) {
	if l <= 0 {
		return 0, fmt.Errorf("the length has to be greater than zero")
	}

	// Ensure the length doesn't exceed the limit for int64
	if l > 19 {
		return 0, fmt.Errorf("length is too large")
	}

	min := int64(math.Pow10(int(l - 1)))
	max := int64(math.Pow10(int(l))) - 1

	// Generate a random number in the range
	//nolint:all
	return min + rand.Int63n(max-min+1), nil
}
