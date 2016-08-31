d64
===

A Go implementation of dominictarr's
[d64](https://github.com/dominictarr/d64).

d64 encodes using a 64 character alphabet (like base64).
Unlike base64, the encoding maintains sort ordering.
It only uses web-safe characters that need no html/url
escaping and play well with html text fields.

The functions for encoding/decoding byte slices are
transcriptions of dominictarr's implementation.

This package also provides functions for encoding/decoding
integers:  with sufficient 0-padding, the encoding maintains
sort order at considerable space savings.

If you need to combine several d64 fields, say
for a combined index, I suggest using "," (comma).
It sorts before any of the characters in the d64
alphabet.

Example
-------

Sortable second-resolution timestamps in 6 digits

    for _, s := range []string{"2000-01-01T00:00:00Z", "2016-06-01T00:00:00Z", "2032-01-01T00:00:00Z"} {
      dt, _ := time.Parse(time.RFC3339, s)
      fmt.Printf("%s\n", EncodeUInt64(uint64(dt.Unix()), 6))
    }
    // Output:
  	// .sQJD.
  	// 0MIXL.
  	// 0obYy.


Copyright & license
-------------------

Copyright (c) YC118, The Elders of Pator Tech School.

See [LICENSE.md](LICENSE.md)
