package newmath
 
import "testing"
 
func TestSqrt(t *testing.T) {
    const given = 4
    const expect = 2
    actual := Sqrt(given)
    if actual != expect {
        t.Errorf("Sqrt(%v) = %v, expect %v", given, actual, expect)
    }
}
