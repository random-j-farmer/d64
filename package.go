/*

Package d64 implements dominictarr's d64 encoding in Go

d64 encodes using a 64 character alphabet (like base64).
Unlike base64, the encoding maintains sort ordering.
It only uses web-safe characters that need no html/uri
escaping and play well with html text fields.

https://github.com/dominictarr/d64

The functions for encoding/decoding byte slices are
transcriptions of dominictarr's implementation.

This package also provides functions for encoding/decoding
integers:  with sufficient 0-padding, the encoding maintains
sort order at considerable space savings.

If you need to combine several d64 fields, say
for a combined index, I suggest using "," (comma).
It sorts before any of the characters in the d64
alphabet.

*/
package d64
