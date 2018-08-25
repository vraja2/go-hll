package hll

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/golang/mock/gomock"
  "go-hll/mocks"
//  "fmt"
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

func TestCountSmallRangeCorrection(t *testing.T) {
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
  assert.Equal(t, 1.0, hllInstance.Count())
}

func TestCountSmallRangeNoCorrection(t *testing.T) {

}

func TestCountIntermediateRangeNoCorrection(t *testing.T) {

}

func TestCountLargeRangeCorrection(t *testing.T) {

}

func TestMerge(t *testing.T) {

}

func TestGetAlphaByNumRegisters(t *testing.T) {
  assert.Equal(t, 0.673,  getAlphaByNumRegisters(16))
  assert.Equal(t, 0.697, getAlphaByNumRegisters(32))
  assert.Equal(t, 0.709, getAlphaByNumRegisters(64))
  assert.Equal(t, 0.7213 / (1 + 1.079 / 128.0), getAlphaByNumRegisters(128))
}
