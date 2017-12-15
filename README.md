# splitsms
[![Build Status](https://travis-ci.org/foril/splitsms.svg?branch=master)](https://travis-ci.org/foril/splitsms) [![Go Report Card](https://goreportcard.com/badge/github.com/foril/splitsms)](https://goreportcard.com/report/github.com/foril/splitsms)

Go library for split SMS.

This library support SMS in GSM 7, basic extended GSM 7 table and Unicode charset.

The size of UDH can be defined for concatened SMS.

[More info on SMS](https://en.wikipedia.org/wiki/GSM_03.38)

```
go get github.com/foril/splitsms
```

## Example

```Go
message := "\n\f\r !\\#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_abcdefghijklmnopqrstuvwxyz{|}~¡£¤¥§¿ÄÅÆÇÉÑÖØÜßàäåæèéìñòöøùüΓΔΘΛΞΠΣΦΨΩ€\f[\\]^{|}~€"

msgInfo := splitsms.Message{FullContent: message}
split, err := msgInfo.Split()

if err != nil {
	fmt.Println(err)
}

fmt.Println(split)
```
### Output

```Go
&{
	GSM // Charset detected
	154 // Total length of message
	180 // Total bytes of message
	2 // Number of SMS in message
	[ // SMS parts (content of sms, SMS bytes, SMS length )
		{\n\f\r !\\#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_abcdefghijklmnopqrstuvwxyz{|}~¡£¤¥§¿ÄÅÆÇÉÑÖØÜßàäåæèéìñòöøùüΓΔΘΛΞΠΣΦ 153 139} 
		{ΨΩ€\f[\\]^{|}~€ 27 15}
	] 
	126 //Characters remaining on last SMS
}

```

## Example - Force UDH length for long message
The default length of UDH is (6), SMS parts limited to 153 characters in GSM 7 and 67 characters in Unicode in this case.

For force to UDH length at (7) bytes

```Go

// Length of SMS is limited to 152 chars in GSM 7 and 66 characters in Unicode
msgInfo := splitsms.Message{FullContent: "My long message .....", UDH: 7}

```

## Example - Force Charset

```Go

// Characters Unicode forced in GSM 7
msgInfo := splitsms.Message{FullContent: "🐿🐿🐿🐿🐿🐿🐿🐿🐿🐿🐿🐿🐿🐿🐿🐿🐿🐿🐿", Charset: "GSM"}

```
