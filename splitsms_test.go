package splitsms

import (
	"sync"
	"testing"
)

// Unicode charset detected
func TestIsGSM7(t *testing.T) {

	sms := "\n\f\r !\\#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_abcdefghijklmnopqrstuvwxyz{|}~¡£¤¥§¿ÄÅÆÇÉÑÖØÜßàäåæèéìñòöøùüΓΔΘΛΞΠΣΦΨΩ€\f[\\]^{|}~€ê"
	if IsGSM7(sms) {
		t.Error("Charset SMS is Unicode")
	}
}

// GSM charset detected
func TestIsUnicode(t *testing.T) {

	sms := "\n\f\r !\\#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_abcdefghijklmnopqrstuvwxyz{|}~¡£¤¥§¿ÄÅÆÇÉÑÖØÜßàäåæèéìñòöøùüΓΔΘΛΞΠΣΦΨΩ€\f[\\]^{|}~€"
	if IsUnicode(sms) {
		t.Error("Charset SMS is GSM 7")
	}
}

// Split for GSM 7
func TestSplitGSM(t *testing.T) {

	// 160 characters , 0 remaining
	msg := Message{FullContent: "----------------------------------------------------------------------------------------------------------------------------------------------------------------"}
	split, _ := msg.Split()

	if split.Length != 160 || split.RemainingChars != 0 {
		t.Errorf("This message contain %d characters and %d characters remaining", split.Length, split.RemainingChars)
	}

	// 161 characters and 145 characters remaining on second SMS - UDH 6 bytes
	msg = Message{FullContent: "-----------------------------------------------------------------------------------------------------------------------------------------------------------------"}
	split, _ = msg.Split()

	if split.Length != 161 || split.RemainingChars != 145 {
		t.Errorf("This message contain %d characters and %d characters remaining on second SMS", split.Length, split.RemainingChars)
	}

	// 306 chars on 2 SMS  - UDH 6 bytes
	msg = Message{FullContent: "------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------"}
	split, _ = msg.Split()

	if split.Length != 306 || split.CountParts != 2 {
		t.Errorf("This message contain %d characters on %d SMS", split.Length, split.CountParts)
	}

	// 159 characters with one character one 2 bytes (160 bytes)  - UDH 6 bytes
	msg = Message{FullContent: "--------------------------------------------------------------------------------------------------------------------------------------------------------€------"}
	split, _ = msg.Split()

	if split.Length != 159 || split.Bytes != 160 {
		t.Errorf("This message contain %d characters on %d bytes", split.Length, split.Bytes)
	}

	// 170 characters with one character one 2 bytes - position 153  - UDH 6 bytes
	msg = Message{FullContent: "--------------------------------------------------------------------------------------------------------------------------------------------------------€-----------------"}
	split, _ = msg.Split()

	if split.Parts[0].Length != 152 || split.Length != 170 {
		t.Errorf("This message contain %d characters on %d bytes", split.Parts[0].Length, split.Length)
	}

}

// Split for Unicode
func TestSplitUnicode(t *testing.T) {

	// 70 characters , 0 remaining
	msg := Message{FullContent: "°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°"}
	split, _ := msg.Split()

	if split.Length != 70 || split.RemainingChars != 0 {
		t.Errorf("This message contain %d characters and %d characters remaining", split.Length, split.RemainingChars)
	}

	// 71 characters and 63 characters remaining on second SMS - UDH 6 bytes
	msg = Message{FullContent: "°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°"}
	split, _ = msg.Split()

	if split.Length != 71 || split.RemainingChars != 63 {
		t.Errorf("This message contain %d characters and %d characters remaining on second SMS", split.Length, split.RemainingChars)
	}

	// 60 characters splited on 2 SMS  - UDH 6 bytes
	msg = Message{FullContent: "°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°"}
	split, _ = msg.Split()

	if split.Length != 134 || split.CountParts != 2 {
		t.Errorf("This message contain %d characters on %d parts", split.Length, split.CountParts)
	}

	// 69 characters with one character one 4 bytes (140 bytes)  - UDH 6 bytes
	msg = Message{FullContent: "°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°🐿°°"}
	split, _ = msg.Split()

	if split.Length != 69 || split.Bytes != 140 {
		t.Errorf("This message contain %d characters on %d bytes", split.Length, split.Bytes)
	}

	// 70 characters with one character one 4 bytes - position 67  - UDH 6 bytes
	msg = Message{FullContent: "°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°🐿°°°"}
	split, _ = msg.Split()

	if split.Parts[0].Length != 66 || split.Length != 70 {
		t.Errorf("This message contain %d characters on %d bytes", split.Parts[0].Length, split.Length)
	}

}

func TestForcesCharset(t *testing.T) {

	// 70 GSM characters forced in Unicode , 0 remaining
	msg := Message{FullContent: "---------------------------------------------------------------------€", Charset: "Unicode"}
	split, _ := msg.Split()

	if split.Length != 70 || split.RemainingChars != 0 || split.Charset != "Unicode" {
		t.Errorf("This message contain %d characters and %d characters remaining, charset : %s", split.Length, split.RemainingChars, split.Charset)
	}

	// 160 Unicode characters forced in GSM 7 , 0 remaining
	msg = Message{FullContent: "°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°", Charset: "GSM"}
	split, _ = msg.Split()

	if split.Length != 160 || split.RemainingChars != 0 || split.Charset != "GSM" {
		t.Errorf("This message contain %d characters and %d characters remaining, charset : %s", split.Length, split.RemainingChars, split.Charset)
	}

	// Charset not supported
	msg = Message{FullContent: "°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°", Charset: "XXXX"}
	_, err := msg.Split()

	if err == nil {
		t.Error("Charset not supported : accepted GSM or Unicode")
	}
}

func TestUDH7(t *testing.T) {

	// 306 chars on 3 SMS  - UDH 7 bytes
	msg := Message{FullContent: "------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------", UDH: 7}
	split, _ := msg.Split()

	if split.Length != 306 || split.CountParts != 3 {
		t.Errorf("This message contain %d characters splited on %d SMS", split.Length, split.CountParts)
	}

	// 134 characters in Unicode splited on 3 SMS  - UDH 7 bytes
	msg = Message{FullContent: "°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°", UDH: 7}
	split, _ = msg.Split()

	if split.Length != 134 || split.CountParts != 3 {
		t.Errorf("This message contain %d characters on %d parts", split.Length, split.CountParts)
	}
}

// TestConcurrentSplitDifferentUDH tests thread safety with different UDH values
func TestConcurrentSplitDifferentUDH(t *testing.T) {
	const goroutines = 100

	var wg sync.WaitGroup
	wg.Add(goroutines)

	errors := make(chan error, goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			defer wg.Done()

			var msg Message
			var expectedParts int

			if idx%2 == 0 {
				// UDH 6 bytes: 306 chars = 2 SMS
				msg = Message{FullContent: "------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------", UDH: 6}
				expectedParts = 2
			} else {
				// UDH 7 bytes: 306 chars = 3 SMS
				msg = Message{FullContent: "------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------", UDH: 7}
				expectedParts = 3
			}

			split, err := msg.Split()
			if err != nil {
				errors <- err
				return
			}

			if split.CountParts != expectedParts {
				t.Errorf("UDH=%d: expected %d parts, got %d", msg.UDH, expectedParts, split.CountParts)
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	for err := range errors {
		t.Errorf("Concurrent split error: %v", err)
	}
}

// TestConcurrentSplitMixedCharsets tests thread safety with GSM and Unicode in parallel
func TestConcurrentSplitMixedCharsets(t *testing.T) {
	const goroutines = 100

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			defer wg.Done()

			var msg Message
			var expectedCharset string

			if idx%2 == 0 {
				// GSM message
				msg = Message{FullContent: "----------------------------------------------------------------------------------------------------------------------------------------------------------------"}
				expectedCharset = "GSM"
			} else {
				// Unicode message
				msg = Message{FullContent: "°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°"}
				expectedCharset = "Unicode"
			}

			split, err := msg.Split()
			if err != nil {
				t.Errorf("Split error: %v", err)
				return
			}

			if split.Charset != expectedCharset {
				t.Errorf("Expected charset %s, got %s", expectedCharset, split.Charset)
			}
		}(i)
	}

	wg.Wait()
}

// TestConcurrentSplitStress stress test with many concurrent calls
func TestConcurrentSplitStress(t *testing.T) {
	const goroutines = 500

	var wg sync.WaitGroup
	wg.Add(goroutines)

	messages := []struct {
		msg            Message
		expectedLength int
		expectedParts  int
	}{
		{Message{FullContent: "Hello World"}, 11, 1},
		{Message{FullContent: "----------------------------------------------------------------------------------------------------------------------------------------------------------------"}, 160, 1},
		{Message{FullContent: "-----------------------------------------------------------------------------------------------------------------------------------------------------------------"}, 161, 2},
		{Message{FullContent: "°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°"}, 70, 1},
		{Message{FullContent: "°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°"}, 71, 2},
	}

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			defer wg.Done()

			tc := messages[idx%len(messages)]
			split, err := tc.msg.Split()
			if err != nil {
				t.Errorf("Split error: %v", err)
				return
			}

			if split.Length != tc.expectedLength {
				t.Errorf("Expected length %d, got %d", tc.expectedLength, split.Length)
			}

			if split.CountParts != tc.expectedParts {
				t.Errorf("Expected %d parts, got %d", tc.expectedParts, split.CountParts)
			}
		}(i)
	}

	wg.Wait()
}
