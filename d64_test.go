package d64

import (
	"fmt"
	"testing"
	"time"
)

func Test_EncodeDecodeUInt64(t *testing.T) {
	testNumber(t, 0, 0)
	testNumber(t, 1, 4)
	testNumber(t, 17, 4)
	testNumber(t, 365, 4)
	testNumber(t, 666, 4)
	testNumber(t, 65000, 4)
	testNumber(t, 0xFFFFFF, 4)
	testNumber(t, 0xFFFFFFFFFFFF, 4)
	testNumber(t, 0xFFFFFFFFFFFFFFFF, 4)

	_, err := DecodeUInt64(" ")
	if err == nil {
		t.Errorf("decode(space) should be an error!")
	}
}

func Test_EncodeDecodeBytes(t *testing.T) {
	testBytes(t, "123")
	testBytes(t, "123456")
	testBytes(t, "Hello, man! How is it hanging?")
	testBytes(t, "abcd")
	testBytes(t, "abcde")

	_, err := DecodeBytes([]byte(" "))
	if err == nil {
		t.Errorf("decode(space) should be an error!")
	}

}

func testNumber(t *testing.T, n uint64, w int) {
	enc := EncodeUInt64(n, w)
	dec, err := DecodeUInt64(enc)
	if err != nil {
		t.Errorf("testNumber: %v", err)
	}

	if n != dec {
		t.Errorf("n<>decoded: %d<>%d", n, dec)
	}

	t.Logf("%s(%d) \t == %d(%d) %#x(%d)\n",
		enc, len(enc), n, len(fmt.Sprintf("%d", n)), n, len(fmt.Sprintf("%x", n)))
}

func testBytes(t *testing.T, s string) {
	b := []byte(s)
	enc := EncodeBytes(b)
	dec, err := DecodeBytes(enc)
	if err != nil {
		t.Errorf("testBytes: %v", err)
	}

	if string(b) != string(dec) {
		t.Errorf("src<>decoded: >%s<>%s<  %d<>%d", b, dec, len(b), len(dec))
	}

	t.Logf("%s(%d) ==> %s(%d)\n", b, len(b), enc, len(enc))
}

func ExampleEncodeUInt64() {
	// sortable second-resolution timestamps in 6 digits
	for _, s := range []string{"2000-01-01T00:00:00Z", "2016-06-01T00:00:00Z", "2032-01-01T00:00:00Z"} {
		dt, _ := time.Parse(time.RFC3339, s)
		fmt.Printf("%s\n", EncodeUInt64(uint64(dt.Unix()), 6))
	}
	// Output:
	// .sQJD.
	// 0MIXL.
	// 0obYy.
}

/*
func main() {
	now := time.Now()
	fmt.Printf("Current time:\t %s %d\n", now, now.Unix())
	myEpoch, _ := time.Parse(time.RFC3339, "2016-01-01T00:00:00Z")
	fmt.Printf("Epoch 1.1.2016:\t %s %d\n", myEpoch, myEpoch.Unix())
	dt := now.Unix() - myEpoch.Unix()
	fmt.Printf("Since Epoch 1.1.2016:\t %d\n", dt)

	printit(uint64(now.Unix()), 0)
	printit(uint64(myEpoch.Unix()), 0)
	printit(uint64(dt), 0)

	charID := 50000000
	pk := 60000000
	exPlain := fmt.Sprintf("%d|%s|%d", charID, now.Format("20060102150405"), pk)
	exD64 := fmt.Sprintf("%s|%s|%s", EncodeUInt64(50000000, 0), EncodeUInt64(uint64(now.Unix()), 0), EncodeUInt64(60000000, 0))
	fmt.Printf("\n\n\nExample index by charid, time, pk\n")
	fmt.Printf("%s ===> %d\n", exPlain, len(exPlain))
	fmt.Printf("%s ===> %d\n", exD64, len(exD64))

	fmt.Printf("\n\n\nExample REFs:\n")
	ref := "00000:00000000"
	encRef := fmt.Sprintf("%s|%s", EncodeUInt64(0, 0), EncodeUInt64(0, 0))
	fmt.Printf("%s => %s\n", ref, encRef)
	ref = "00000:02a00000"
	encRef = fmt.Sprintf("%s|%s", EncodeUInt64(0, 0), EncodeUInt64(0x2a00000, 0))
	fmt.Printf("%s => %s\n", ref, encRef)
	ref = "63000:3FFFFFFF"
	encRef = fmt.Sprintf("%s|%s", EncodeUInt64(63000, 0), EncodeUInt64(0x3FFFFFFF, 0))
	fmt.Printf("%s => %s\n", ref, encRef)

	fmt.Printf("\n\n\n\nEncodeBytes:\n")
}
*/
