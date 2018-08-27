# go-hll
An implementation of HyperLogLog written in Go. HyperLogLog is a cardinality estimation algorithm explained well [here](https://research.neustar.biz/2012/10/25/sketch-of-the-day-hyperloglog-cornerstone-of-a-big-data-infrastructure/). 

## Usage

### Create an HLL instance
`hllInstance := NewHLLWithRegisterBits(numRegisterBits)`  
**Note**: defaults to 6 register bits if not explicitly specified

### Add a value to the HLL instance
```
hllInstance.AddString("hello")
hllInstance.AddHash(123456)
```
Under the hood, `AddString` uses `MurmurHash32` to hash the provided string into a 32 bit integer.

### Get the count/cardinality of the HLL instance
`count := hllInstance.Count()`

## Development

### Running Tests

`go test`

### Formatting

All code is formatted with `go fmt`
