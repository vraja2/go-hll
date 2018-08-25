package hll

import (
  "github.com/spaolacci/murmur3"
  "hash"
  "math"
  "math/bits"
//  "fmt"
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
  hll.registers[registerIndex] = int(math.Max(float64(hll.registers[registerIndex]), float64(registerValue + 1)))
}

func (hll HLL) Count() float64 {
  harmonicMean := 0.0
  numZeroes := 0.0
  for _, registerVal := range hll.registers {
    harmonicMean += 1.0 / math.Pow(2.0, float64(registerVal))
    if registerVal == 0 {
      numZeroes += 1.0
    }
  }
  harmonicMean = 1.0 / harmonicMean
  // TODO: figure out what alpha param means
  estimate := getAlphaByNumRegisters(len(hll.registers)) * math.Pow(float64(len(hll.registers)), float64(2)) * float64(harmonicMean)
  count := 0.0
  // small range correction
  if estimate <= (5.0 / 2.0) * float64(len(hll.registers)) {
    if numZeroes == 0 {
      count = estimate
    } else {
      count = math.Round(float64(len(hll.registers)) * math.Log2(float64(len(hll.registers)) / numZeroes))
    }

    return count
  }

  if estimate <= 1.0 / 30.0 * math.Exp2(32.0) {
    // intermediate range, no correction
    count = estimate
  } else {
    // large range correction
    count = math.Pow(-2.0, 32.0) * math.Log2(1 - estimate / math.Pow(2, 32))
  }

  return count
}

func Merge() {

}

func getAlphaByNumRegisters(numRegisters int) float64 {
  var alpha float64
  if numRegisters == 16 {
    alpha = 0.673
  } else if numRegisters == 32 {
    alpha = 0.697
  } else if numRegisters == 64 {
    alpha = 0.709
  } else {
    alpha = 0.7213 / (1 + 1.079 / float64(numRegisters))
  }

  return alpha
}