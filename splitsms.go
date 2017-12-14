package splitsms

/**
* Go library for split SMS.
* This library support SMS in GSM 7, basic extended GSM 7 table and Unicode charset.
* The size of UDH can be defined for concatened SMS.
* https://en.wikipedia.org/wiki/GSM_03.38#GSM_7-bit_default_alphabet_and_extension_table_of_3GPP_TS_23.038_.2F_GSM_03.38
 */

import (
	"errors"
	_ "log"
	"unicode/utf16"
)

const (
	smsBytes      int = 140
	singleGsm7    int = 160
	singleUnicode int = 70
)

var (
	multiGSM7    int
	multiUnicode int
)

type Message struct {
	FullContent string
	Charset     string
	UDH         int
}

type Split struct {
	Charset        string // GSM or Unicode
	Length         int    // Total length of message
	Bytes          int    // Total bytes of message
	CountParts     int    // Number of SMS in message
	Parts          []Sms  // SMS parts
	RemainingChars int    // Remaining char in current SMS
}

type Sms struct {
	Content string
	Bytes   int
	Length  int
}

func IsGSM7(msg string) bool {
	for _, c := range msg {
		_, gsm7 := gsm7Chars[c]
		if !gsm7 {
			return false
		}
	}
	return true
}

func isGSM7Ext(c rune) bool {
	_, gsm7Ext := gsm7ExtChars[c]

	return gsm7Ext
}

func IsUnicode(msg string) bool {
	return !IsGSM7(msg)
}

func (m *Message) Split() (*Split, error) {

	var isGsm bool

	msg := m.FullContent

	if m.Charset != "" && m.Charset != "GSM" && m.Charset != "Unicode" {
		return nil, errors.New("Charset not supported - GSM or Unicode")
	}

	if m.UDH != 0 && (m.UDH != 6 && m.UDH != 7) {
		return nil, errors.New("UDH Length is 6 bytes or 7 bytes")
	}

	// Default isGsm is false at declaration (Unicode)
	if m.Charset != "" && m.Charset == "GSM" {
		isGsm = true
	} else if m.Charset != "Unicode" {
		isGsm = IsGSM7(msg)
	}

	if m.UDH == 7 {
		multiGSM7 = ((smsBytes * 8) - (m.UDH * 8)) / 7
		multiUnicode = ((smsBytes * 8) - (m.UDH * 8)) / 16
	} else {
		// Default UDH Length 6 bytes
		multiGSM7 = ((smsBytes * 8) - (6 * 8)) / 7
		multiUnicode = ((smsBytes * 8) - (6 * 8)) / 16
	}

	split := &Split{Charset: "GSM"}

	bytes := 0
	length := 0
	curSMS := ""

	for _, char := range msg {

		if isGsm {

			if isGSM7Ext(char) {

				if bytes == multiGSM7-1 {
					split.appendSms(curSMS, bytes, length)
					bytes = 0
					length = 0
					curSMS = ""
				}
				// Add escape code
				bytes++
			}

			bytes++
			length++

		} else {

			split.Charset = "Unicode"

			if isHighSurrogate(char) {

				if bytes == (multiUnicode-1)*2 {
					split.appendSms(curSMS, bytes, length)
					bytes = 0
					length = 0
					curSMS = ""
				}

				bytes += 2
			}

			bytes += 2
			length++

		}

		curSMS += string(char)

		if (isGsm && bytes == multiGSM7) || (!isGsm && bytes == (multiUnicode*2)) {
			split.appendSms(curSMS, bytes, length)
			bytes = 0
			length = 0
			curSMS = ""
		}
	}

	split.appendSms(curSMS, bytes, length)
	bytes = 0
	length = 0
	curSMS = ""

	if (isGsm && len(split.Parts) > 1 && split.Bytes <= singleGsm7) || (!isGsm && len(split.Parts) > 1 && split.Bytes <= (singleUnicode*2)) {
		split.Parts[0].Content += split.Parts[1].Content
		split.Parts[0].Bytes += split.Parts[1].Bytes
		split.Parts[0].Length += split.Parts[1].Length
		split.Parts = split.Parts[:len(split.Parts)-1]
		split.CountParts = 1

		if isGsm {
			split.RemainingChars = singleGsm7 - split.Bytes
		} else {
			split.RemainingChars = (singleUnicode * 2) - split.Bytes
		}
	}

	return split, nil
}

func (m *Split) appendSms(sms string, bytes int, length int) {

	if bytes > 0 {
		m.Parts = append(m.Parts, Sms{sms, bytes, length})
		m.Length += length
		m.Bytes += bytes
		m.CountParts = len(m.Parts)

		if m.Charset == "GSM" {
			m.RemainingChars = singleGsm7 - m.Bytes
			if len(m.Parts) > 1 {
				m.RemainingChars = multiGSM7 - m.Parts[len(m.Parts)-1].Bytes
			}
		} else {
			m.RemainingChars = singleUnicode - (m.Bytes / 2)
			if len(m.Parts) > 1 {
				m.RemainingChars = multiUnicode - (m.Parts[len(m.Parts)-1].Bytes / 2)
			}
		}
	}

}

func isHighSurrogate(r rune) bool {
	r1, _ := utf16.EncodeRune(r)
	return r1 >= 0xD800 && r1 <= 0xDBFF
}
