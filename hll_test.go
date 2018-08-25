package hll

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/golang/mock/gomock"
  "go-hll/mocks"
  "fmt"
)

func TestNewHLLWithRegisterBits(t *testing.T) {
  numRegisterBits := 6
  numRegisters := 64
  hllInstance := NewHLLWithRegisterBits(numRegisterBits)
  assert.Equal(t, numRegisterBits, hllInstance.numRegisterBits, "Num register bits should be equal")
  expectedRegisters := make([]int, numRegisters)
  assert.Equal(t, expectedRegisters, hllInstance.registers, "Registers should be equal")
}

func TestAdd(t *testing.T) {
  numRegisterBits := 6
  numRegisters := 64
  registers := make([]int, numRegisters)
  mockCtrl := gomock.NewController(t)
  defer mockCtrl.Finish()
  mockMurmur32 := mocks.NewMockHash32(mockCtrl)
  hllInstance := HLL {numRegisterBits, registers, mockMurmur32}
  mockMurmur32.EXPECT().Write([]byte("hello"))
  // 111000001010011100110110
  mockMurmur32.EXPECT().Sum32().Return(uint32(3688933174))
  hllInstance.Add("hello")
  for idx, registerVal := range hllInstance.registers {
    // last 6 bits: 110110 = 54
    if idx == 54 {
      assert.Equal(t, 3, registerVal)
    } else {
      assert.Equal(t, 0, registerVal)
    }
  }
}

func TestCount(t *testing.T) {
  numRegisterBits := 6
  numRegisters := 64
  registers := make([]int, numRegisters)
  mockCtrl := gomock.NewController(t)
  defer mockCtrl.Finish()
  mockMurmur32 := mocks.NewMockHash32(mockCtrl)
  hllInstance := HLL {numRegisterBits, registers, mockMurmur32}
  mockMurmur32.EXPECT().Write([]byte("hello"))
  // 111000001010011100110110
  mockMurmur32.EXPECT().Sum32().Return(uint32(2903491477))
  hllInstance.Add("hello")
  fmt.Printf("%d\n", hllInstance.registers[21])
  fmt.Printf("%f\n", hllInstance.Count())
}