package rightCounter

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"reflect"
)

type Testfunc func([...]int64) [...]int64

func RangeRand(min, max int64) int64 {
	// return a random number in [min, max)
	if min > max {
		panic("the min is greater than max!")
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}

func IsEquals(testTime, maxSize, maxValue int64, test1, test2 Testfunc) {
	success := true
	for i := int64(0); i < testTime; i++ {
		var arr [...]int64
		for j := int64(0); j < maxSize; j++ {
			arr[j] = RangeRand(0, maxValue)
		}
		arr1 := test1(arr)
		arr2 := test2(arr)
		if reflect.DeepEqual(arr1, arr2) {
			success = false
			break
		}
		fmt.Printf("%+v\n: %v", success, i)
	}
}
