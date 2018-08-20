package hll

import (
  "github.com/spaolacci/murmur3"
  "hash"
  "math"
  "math/bits"
  // "fmt"
)

type HLL struct {
  numRegisterBits int
  registers []int
  murmur32 hash.Hash32
}

func NewHLL() HLL {
  return NewHLLWithRegisterBits(6)
}

func NewHLLWithRegisterBits(numRegisterBits int) HLL {
  numRegisters := int(math.Exp2(float64(numRegisterBits)))
  registers := make([]int, numRegisters)
  murmur32 := murmur3.New32()
  hllInstance := HLL {numRegisterBits, registers, murmur32}
  return hllInstance
}

func (hll HLL) Add(value string) {
  hll.murmur32.Write([]byte(value))
  hashedValue := hll.murmur32.Sum32()
  // bit mask to fetch bits representing register index to update
  maskRegisterBits := ^uint32(0) >> uint32(32 - hll.numRegisterBits)
  registerIndex := uint32(hashedValue & maskRegisterBits)
  remainingBits := hashedValue >> uint32(hll.numRegisterBits)
  trailingZeroes := bits.TrailingZeros32(remainingBits)
  registerValue := 0
  if trailingZeroes != 32 {
    registerValue = trailingZeroes
  }
  hll.registers[registerIndex] = int(math.Max(float64(hll.registers[registerIndex]), float64(registerValue)))
}

func Count() {

}

func Merge() {

}
