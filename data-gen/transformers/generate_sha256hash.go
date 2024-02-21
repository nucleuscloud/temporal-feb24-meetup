package transformers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"

	"github.com/benthosdev/benthos/v4/public/bloblang"
)

func init() {
	spec := bloblang.NewPluginSpec()

	err := bloblang.RegisterFunctionV2("generate_sha256hash", spec, func(args *bloblang.ParsedParams) (bloblang.Function, error) {
		return func() (any, error) {
			val, err := GenerateRandomSHA256Hash()

			if err != nil {
				return false, fmt.Errorf("unable to generate sha256 hash")
			}
			return val, nil
		}, nil
	})
	if err != nil {
		panic(err)
	}
}

/* Generates a random SHA256 hashed value */
func GenerateRandomSHA256Hash() (string, error) {
	min := int64(1)
	max := int64(9)

	str, err := generateRandomStringWithInclusiveBounds(min, max)
	if err != nil {
		return "", err
	}

	// hash the value
	bites := []byte(str)
	hasher := sha256.New()
	_, err = hasher.Write(bites)
	if err != nil {
		return "", err
	}

	// compute sha256 checksum and encode it into a hex string
	hashed := hasher.Sum(nil)
	var buf bytes.Buffer
	e := hex.NewEncoder(&buf)
	_, err = e.Write(hashed)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

var alphanumeric = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"

func generateRandomStringWithInclusiveBounds(min, max int64) (string, error) {
	if min < 0 || max < 0 || min > max {
		return "", fmt.Errorf("the min and max can't be less than 0 and the min can't be greater than the max")
	}

	var length int64

	if min == max {
		length = min
	} else {
		randlength, err := generateRandomInt64InValueRange(min, max)
		if err != nil {
			return "", fmt.Errorf("unable to generate a random length for the string")
		}

		length = randlength
	}

	result := make([]byte, length)

	for i := int64(0); i < length; i++ {
		// Generate a random index in the range [0, len(alphabet))
		//nolint:all
		index := rand.Intn(len(alphanumeric))

		// Get the character at the generated index and append it to the result
		result[i] = alphanumeric[index]
	}

	return strings.ToLower(string(result)), nil
}

func generateRandomInt64(randomizeSign bool, min, max int64) (int64, error) {
	var returnValue int64

	if randomizeSign {
		res, err := generateRandomInt64InValueRange(absInt64(min), absInt64(max))
		if err != nil {
			return 0, err
		}

		returnValue = res
		randInt := rand.Intn(2)
		if randInt == 1 {
			// return the positive value
			return returnValue, nil
		} else {
			// return the negative value
			return returnValue * -1, nil
		}
	} else {
		res, err := generateRandomInt64InValueRange(min, max)
		if err != nil {
			return 0, err
		}

		returnValue = res
	}

	return returnValue, nil
}

func generateRandomInt64InValueRange(min, max int64) (int64, error) {
	if min > max {
		min, max = max, min
	}

	if min == max {
		return min, nil
	}

	rangeVal := max - min + 1
	return min + rand.Int63n(rangeVal), nil
}

func absInt64(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}
