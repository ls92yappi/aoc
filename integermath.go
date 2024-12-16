package aoc

///////////////////////////////////////////////////////////
//                                                       //
//   Many of these functions below operate on or return  //
//   int slices. None have any import dependencies.      //
//                                                       //
//                                                       //
//                                                       //
///////////////////////////////////////////////////////////
//                                                       //
// const MaxInt int = 9_223_372_036_854_775_807          //
// const Ten18  int = 1_000_000_000_000_000_000          //
//                                                       //
// var InversePowersOf2 map[int]int      =62 max         //
// var Factorials []int                 [20] max         //
// var LargestPrimeLessThan2Nth []int   [37-1] max       //
//                                                       //
// func NumDigits(n int) int                             //
// func TenToTheN(n int) int                             //
// func Abs(x int) int                                   //
// func Sign(x int) int                                  //
// func CeilPow2(x int) int                              //
// func FloorPow2(x int) int                             //
// func IsPow2(v int) bool                               //
// func SqrtFloor(x int) (int,bool)                      //
// func Factor(num int) []int                            //
// func GCD(a, b int) int                                //
// func LCMv(a, b int, integers ...int) int              //
// func LCM(a []int) int                                 //
// func Sum(a []int) int                                 //
// func Prod(a []int) int                                //
// func Min2(x, y int) int                               //
// func Max2(x, y int) int                               //
// func Min(a []int) int                                 //
// func Max(a []int) int                                 //
// func Factorial(n int) (int, error)                    //
// func Binomial(n, k int) (int,error)                   //
// func Permutations(n, k int) (int,error)               //
// func ModSolver(d1, r1, d2, r2 int) int                //
//                                                       //
///////////////////////////////////////////////////////////

type TooBigFor64Bits struct{}

func (e *TooBigFor64Bits) Error() string {
	return "Too big to fit in 64-bits"
}

const MaxInt int = 9_223_372_036_854_775_807
const Ten18  int = 1_000_000_000_000_000_000

// Largest Prime < 2^n, from n=1 to n=37 from https://oeis.org/A014234/list
var LargestPrimeLessThan2Nth []int = []int{
	2,3,7,13,31,61,127,251,509,1021,2039,4093,8191,
	16381,32749,65521,131071,262139,524287,1048573,
	2097143,4194301,8388593,16777213,33554393,
	67108859,134217689,268435399,536870909,1073741789,
	2147483647,4294967291,8589934583,17179869143,
	34359738337,68719476731,137438953447}

/*
// Per https://go.dev/ref/spec#Constants
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
//
int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
*/

var InversePowersOf2 map[int]int // defined in init()
// fibonacci grows slowly enough there's no real need to Memoize it OEIS A000045
var Factorials []int // defined in init() https://oeis.org/A000142/list

func init() {
	InversePowersOf2 = map[int]int{
	    1:0,2:1,4:2,8:3,
	    16:4,32:5,64:6,128:7,
	    256:8,512:9,1024:10,2048:11,
	    4092:12,8192:13,16384:14,32768:15,
	    65536:16,131072:17,262144:18,524288:19,
	    1048576:20,2097152:21,4194304:22,8388608:23,
	    16777216:24,33554432:25,67108864:26,134217728:27,
	    268435456:28,536870912:29,1073741824:30,2147483648:31,
	    4294967296:32,8589934592:33,17179869184:34,34359738368:35,
	    68719476736:36,137438953472:37,274877906944:38,549755813888:39,
	    1_099_511_627_776:40,2_199_023_255_552:41,
	    4_398_046_511_104:42,8_796_093_022_208:43,
	    17_592_186_044_416:44,35_184_372_088_832:45,
	    70_368_744_177_664:46,140_737_488_355_328:47,
	    281_474_976_710_656:48,562_949_953_421_312:49,
	    1_125_899_906_842_624:50,2_251_799_813_685_248:51,
	    4_503_599_627_370_496:52,9_007_199_254_740_992:53,
	    18_014_398_509_481_984:54,36_028_797_018_963_968:55,
	    72_057_594_037_927_936:56,144_115_188_075_855_872:57,
	    288_230_376_151_711_744:58,576_460_752_303_423_488:59,
	    1_152_921_504_606_846_976:60,2_305_843_009_213_693_952:61,
	    4_611_686_018_427_387_904:62,
	    //9_223_372_036_854_775_808:63, 2^63 causes signed overflow
	}
	Factorials = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800,
		39916800, 479001600, 6227020800, 87178291200, 1307674368000,
		20922789888000, 355687428096000, 6402373705728000, 121645100408832000, 2432902008176640000,
	}   //51_090_942_171_709_440_000, 1124000727777607680000 are too big
}

// Number of Digits in an integer. Max is 18 for 64-bit numbers.
func NumDigits(n int) int {
	if (n < 0) {
		return 0
	}
	cmp := Ten18
	dig := 0
	for cmp = Ten18; cmp>=1; cmp/=10 {
		if n >= cmp {
			dig++
		}
	}
	return dig
}

// 10^n
func TenToTheN(n int) int {
	base := 1
	for range(n) {
		base *= 10
	}
	return base
}

// Maximum Power of 2 <= x
func FloorPow2(x int) int {
	x = x | (x >> 1)
	x = x | (x >> 2)
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x | (x >> 16)
	x = x | (x >> 32)
	return x - (x >> 1)
}

// Minimum Power of 2 >= x
func CeilPow2(x int) int {
	x = x - 1
	x = x | (x >> 1)
	x = x | (x >> 2)
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x | (x >> 16)
	x = x | (x >> 32)
	return x + 1
}

func IsPow2(v int) bool {
	return v!=0 && (v & (v - 1))!=0
}

// Compute Integer SqrtFloor, returns the floor, and if it is a perfect square
// Negatives return the negative of its positive's square root and are Never perfect
// SqrtFloor is computed via binary search
func SqrtFloor(x int) (int,bool) {
    var sign = 1
    if x < 0 {
        x = -x
        sign = -1
    }
    // Base Cases
    if x < 2 {
        return x*sign, sign>0
    }
 
    // Do Binary Search for floor(sqrt(x))
    var start,end,ans int = 1,x/2,0

    for start <= end {
        mid := (start + end) / 2;
 
        // If x is a perfect square
        if (mid * mid == x) {
            return mid*sign, sign>0
        }
 
        // Since we need floor, we update answer when
        // mid*mid is smaller than x, and move closer to
        // sqrt(x)
        if (mid * mid < x) {
            start = mid + 1
            ans = mid
        } else {
            // If mid*mid is greater than x
            end = mid - 1
        }
    }
    return ans*sign,false
}

// After testing, one cannot parallelize GCD() computations of a slice
// without using language level parallelism. Doing so would require
// the same number of steps for each path to resolve.

// Find the Greatest Common Divisor via Euclidean algorithm
// Assumes a,b >= 0
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Find the Least Common Multiple (LCM) via GCD with a variadic function
// Assumes a,b,integers >= 0
func LCMv(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCMv(result, integers[i])
	}
	return result
}

// Without recursion or variadic function. Uses less memory
// Assumes a[] values each >= 0
func LCM(a []int) int {
	if len(a) == 0 {
		return 1
	}
	if len(a) == 1 {
		return a[0]
	}
	result := a[0] * a[1] / GCD(a[0], a[1])
	for i := 2; i < len(a); i++ {
		result = a[i] * result / GCD(a[i], result)
	}
	return result
}

func Sum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	r := a[0]
	for i := 1; i < len(a); i++ {
		r += a[i]
	}
	return r
}

func Prod(a []int) int {
	if len(a) == 0 {
		return 1
	}
	r := a[0]
	for i := 1; i < len(a); i++ {
		r *= a[i]
	}
	return r
}

func Abs(x int) int {
	r := x * (1 - 2*( (x*3)/(x*3+1) ))
	return r
}

func Sign(x int) int {
	if x >= 0 { return 1 }
	return -1
}

func Min2(x, y int) int {
	d := x - y
	abs := d * (1 - ((d*3)/(d*3+1))*2)
	r := (x + y - abs) / 2
	return r
}


func Max2(x, y int) int {
	d := x - y
	abs := d * (1 - ((d*3)/(d*3+1))*2)
	r := (x + y + abs) / 2
	return r
}

// no conditionals inside the for loop for speed
func Min(a []int) int {
	if len(a) == 0 {
		return 0
	}
	r := a[0]
	for i := 1; i < len(a); i++ {
		d := r - a[i]
		abs := d * (1 - ((d*3)/(d*3+1))*2)
		r = (r + a[i] - abs) / 2
	}
	return r
}

// no conditionals inside the for loop for speed
func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	r := a[0]
	for i := 1; i < len(a); i++ {
		d := r - a[i]
		abs := d * (1 - ((d*3)/(d*3+1))*2)
		r = (r + a[i] + abs) / 2
	}
	return r
}

// See https://t5k.org/lists/small/millions/ for examples of primes
//n := 198491317*179424673 //11 millionth prime * 10 millionth prime // my Factor() takes about 13 seconds to compute
// Prime factorize an integer
func Factor(num int) []int {
	var fl []int = make([]int,0)
	if num<0 {
		fl = append(fl, -1)
		num = -num
	}
	if num==0 || num==1 {
		fl = append(fl, num)
		return fl
	}

	i := 0
	// strip off powers of 2 first
	for {
		if num%2 == 0 {
			fl = append(fl, 2)
			num /= 2
		} else {
			// no longer divisible by powers of 2, move on
			break
		}
	} // factors of 2
	// then strip off odd factors
	sq_root,perf := SqrtFloor(num)
	// if we found a perfect square, return early
	if perf {
		fl = append(fl,Factor(sq_root)...)
		fl = append(fl,Factor(sq_root)...)
		return fl
	}
	//sq_root := int(math.Sqrt(float64(num)))
	for i = 3; i <= sq_root; i+=2 {
		for {
			if num%i == 0 {
				fl = append(fl, i)
				num /= i
			} else {
				// no longer divisible by powers of i, move on
				break
			}
		} // factors of odd i
		// revise upper bound downward as we go through the possible factors
		sq_root,perf = SqrtFloor(num)
		//sq_root = int(math.Sqrt(float64(num)))
		// if we found a perfect square, return early
		if perf {
			fl = append(fl,sq_root)
			fl = append(fl,sq_root)
			return fl
		}
	} // all possible odd i factors
	// list final factor if not ending in a square
	if i > sq_root {
		if num != 1 {
			fl = append(fl, num)
		}
	}
	return fl
}

// Return n! where n<=20 (fits in 64-bit int) else returns 0,error
func Factorial(num int) (int, error) {
	// Factorials 0..20 from https://oeis.org/A000142
	if num > len(Factorials)-1 {
		return 0, &TooBigFor64Bits{}
	}

	return Factorials[num],nil
	//if num < 3 {
	//	return num
	//}
	//return num*Factorial(num-1)
}

// 86_C_15 is ok, even though max 20!
// Also known as n Choose k, or n_C_k or Combinations(n,k)
func Binomial(n, k int) (int,error) {
	// Assumes 0 <= k <= n, else 0
	if k > n  || k < 0 || n < 0 {
		//fmt.Println("Bad Range")
		return 0,nil
	}
	// max factorial that fits within 64-bits
	if n < len(Factorials) {
		nf,_ := Factorial(n)
		kf,_ := Factorial(k)
		nkf,_ := Factorial(n-k)
		return (nf / kf / nkf),nil
	}
	// if difference between the numbers is too great to fit in 64-bits
	r := n-k
	l := Min2(r, k)
	if l > len(Factorials)-1 {
		//fmt.Println("Too big a difference")
		return 0,&TooBigFor64Bits{}
	}
	// attempt to pre-factor out Max2(k!,(n-k)!), then
	// alternately multiply followed by divide
	product := 1
	last := 1
	for i:=n-l+1; i<=n; i++ {
		last = product
		product *= i
		// check for overflow on multiplication
		if product < last  || (MaxInt/product) < (i-n+l) {
			//fmt.Println("Overflow on multiplication")
			return 0,&TooBigFor64Bits{}
		}
		// alternate mults with divisions
		product = product / (i-n+l)
	}
	//lf,_ := Factorial(l)
	//return (product/lf),nil
	return (product),nil
}

// 30_P_12 is ok, even though max 20!
// Permutations, also known as Arrangements, or Orders, or n_P_k
func Permutations(n, k int) (int,error) {
	// Assumes 0 <= k <= n, else 0
	if k > n  || k < 1 || n < 1 {
		//fmt.Println("Bad Range")
		return 0,nil
	}
	// max factorial that fits within 64-bits
	if n < len(Factorials) {
		nf,_ := Factorial(n)
		kf,_ := Factorial(k)
		return (nf / kf),nil
	}
	// Otherwise, try to compute it, pre-factoring out (n-k)!
	product := 1
	last := 1
	for i:=n-k+1; i<=n; i++ {
		last = product
		product *= i
		// check for overflow on multiplication
		if product < last  || (MaxInt/product) < i {
			//fmt.Println("Overflow on multiplication")
			return 0,&TooBigFor64Bits{}
		}
	}
	return product, nil
}

// find n such that n%d1=r1 && n%d2=r2
// solves r1 base repeats every d1 passes and r2 base repeats every d2 passes
// effectively, this is the Chinese Remainder Theorem for 2 divisors
func ModSolver(d1, r1, d2, r2 int) int {
	// verify we have valid modular arithmetic numbers
	if d1 <= 0 || d2 <= 0 || r1 < 0 || r2 < 0 || r1 >= d1 || r2 >= d2 {
		return -1
	}

	// limited trial and error
	guess := r1
	for range(d2) {
		guess += d1 // maintains guess%d1=r1
		if guess%d2 == r2 {
			return guess
		}
	}

	// This should never occur given the top filter above
	return -1
}
