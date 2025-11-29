# Acknowledgements

This project makes use of and was inspired by code from:

- [Hypersequent/uuid7](https://github.com/Hypersequent/uuid7) (MIT License)  
- [google/uuid](https://github.com/google/uuid) (BSD License)

# UUIDv7

This package implement [rfc9562](https://www.rfc-editor.org/rfc/rfc9562.html) [Section 5.7. UUID Version 7](https://www.rfc-editor.org/rfc/rfc9562.html#name-uuid-version-7).

The origin structure of UUID Version 7 are following the structure below:

```
 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                           unix_ts_ms                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|          unix_ts_ms           |  ver  |       rand_a          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|var|                        rand_b                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                            rand_b                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```

- unix_ts_ms:
    48-bit big-endian unsigned number of the Unix Epoch timestamp in milliseconds as per Section 6.1. Occupies bits 0 through 47 (octets 0-5).
- ver:
    The 4-bit version field as defined by [Section 4.2](https://www.rfc-editor.org/rfc/rfc9562.html#version_field), set to 0b0111 (7). Occupies bits 48 through 51 of octet 6.
- rand_a:
    12 bits of pseudorandom data to provide uniqueness as per [Section 6.9](https://www.rfc-editor.org/rfc/rfc9562.html#unguessability) and/or optional constructs to guarantee additional monotonicity as per [Section 6.2](https://www.rfc-editor.org/rfc/rfc9562.html#monotonicity_counters). Occupies bits 52 through 63 (octets 6-7).
- var:
    The 2-bit variant field as defined by [Section 4.1](https://www.rfc-editor.org/rfc/rfc9562.html#variant_field), set to 0b10. Occupies bits 64 and 65 of octet 8.
- rand_b:
    The final 62 bits of pseudorandom data to provide uniqueness as per [Section 6.9](https://www.rfc-editor.org/rfc/rfc9562.html#unguessability) and/or an optional counter to guarantee additional monotonicity as per [Section 6.2](https://www.rfc-editor.org/rfc/rfc9562.html#monotonicity_counters). Occupies bits 66 through 127 (octets 8-15).

