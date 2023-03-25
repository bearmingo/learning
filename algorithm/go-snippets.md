#golang

# Golang Code Snippets

## Max and Min

```go
// 32
const MaxUint32 = int(^uint32(0))       // 4294967295
const MinUint32 = 0                     // 0
const MaxInt32 = int(MaxUint32 >> 1)    // 2147483647
const MinInt32 = -MaxInt32 - 1          // -2147483648

const MaxUint = ^uint(0)   // 18446744073709551615
const MinUint = 0          // 0

const MaxInt = int(MaxUint >> 1) // 9223372036854775807
const MinInt = -MaxInt - 1       // -9223372036854775808
```

## XOR

```go
if boolA != boolB {

}
```