# hexen

hex viewer (and eventually editor) with support for custom decoding text tables. These tools only seem to exist for Windows, so this is an attempt to bring the feature to linux.

**until this is removed, assume this is totally broken**

## Use

```
# Dump hex blob with charTable applied
#
# blob: binary data
# charTable: translation table alternate to ASCII
hexen [blob] [charTable]

```

## TODO
- [x] dump hex from blobs
- [x] wrap encoding/hex
- [x] data structure for charlist
- [x] intercept hex.Dump with charlist
- [ ] support multi-char replacement
- [ ] all the editor things

