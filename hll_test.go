package hll

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-hll/mocks"
	"testing"
)

func TestNewHLLWithRegisterBits(t *testing.T) {
	numRegisterBits := 6
	numRegisters := 64
	hllInstance := NewHLLWithRegisterBits(numRegisterBits)
	assert.Equal(t, numRegisterBits, hllInstance.numRegisterBits, "Num register bits should be equal")
	expectedRegisters := make([]int, numRegisters)
	assert.Equal(t, expectedRegisters, hllInstance.registers, "Registers should be equal")
}

func TestAddString(t *testing.T) {
	numRegisterBits := 6
	numRegisters := 64
	registers := make([]int, numRegisters)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockMurmur32 := mocks.NewMockHash32(mockCtrl)
	hllInstance := HLL{numRegisterBits, registers, mockMurmur32}
	mockMurmur32.EXPECT().Write([]byte("hello"))
	// sample input/output from http://content.research.neustar.biz/blog/hll.html
	// 111000001010011100110110
	mockMurmur32.EXPECT().Sum32().Return(uint32(3688933174))
	hllInstance.AddString("hello")
	for idx, registerVal := range hllInstance.registers {
		// last 6 bits: 110110 = 54
		if idx == 54 {
			assert.Equal(t, 3, registerVal)
		} else {
			assert.Equal(t, 0, registerVal)
		}
	}
}

func TestAddCountSequence(t *testing.T) {
        assert := assert.New(t)
        numRegisterBits := 6
	hllInstance := NewHLLWithRegisterBits(numRegisterBits)
        hllInstance.AddHash(1455387899)
        assert.Equal(1.0, hllInstance.Count())
        hllInstance.AddHash(619839696)
        assert.Equal(2.0, hllInstance.Count())
        hllInstance.AddHash(3568685273)
        assert.Equal(3.0, hllInstance.Count())
        hllInstance.AddHash(1436505107)
        assert.Equal(4.0, hllInstance.Count())
        hllInstance.AddHash(2298164309)
        assert.Equal(5.0, hllInstance.Count())
}

func TestCountSmallRangeCorrection(t *testing.T) {
	numRegisterBits := 6
	numRegisters := 64
	registers := make([]int, numRegisters)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockMurmur32 := mocks.NewMockHash32(mockCtrl)
	hllInstance := HLL{numRegisterBits, registers, mockMurmur32}
	// sample input/output from http://content.research.neustar.biz/blog/hll.html
	// 100101101011110000110000
	hllInstance.AddHash(uint32(1771486256))
	assert.Equal(t, 1.0, hllInstance.Count())
}

func TestCountSmallRangeNoCorrection(t *testing.T) {
	numRegisterBits := 4
	numRegisters := 16
	registers := make([]int, numRegisters)
	// set all registers to 1 to force numZeroRegisters == 0 condiition
	for idx := range registers {
		registers[idx] = 1
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockMurmur32 := mocks.NewMockHash32(mockCtrl)
	hllInstance := HLL{numRegisterBits, registers, mockMurmur32}
	assert.Equal(t, 21.536, hllInstance.Count())
}

func TestCountIntermediateRangeNoCorrection(t *testing.T) {
	numRegisterBits := 6
	numRegisters := 64
	registers := make([]int, numRegisters)
	// set all registers to 20 to force the intermediate range case
	for idx := range registers {
		registers[idx] = 20
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockMurmur32 := mocks.NewMockHash32(mockCtrl)
	hllInstance := HLL{numRegisterBits, registers, mockMurmur32}
	assert.Equal(t, 47580184.576, hllInstance.Count())
}

func TestCountLargeRangeCorrection(t *testing.T) {
	numRegisterBits := 6
	numRegisters := 64
	registers := make([]int, numRegisters)
	// set all registers to 26 to force large range correction
	for idx := range registers {
		registers[idx] = 25
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockMurmur32 := mocks.NewMockHash32(mockCtrl)
	hllInstance := HLL{numRegisterBits, registers, mockMurmur32}
	assert.Equal(t, 2712319089.293557, hllInstance.Count())
}

func TestMergeUnequalRegisterBits(t *testing.T) {
	hllInstance1 := NewHLLWithRegisterBits(6)
	hllInstance2 := NewHLLWithRegisterBits(5)
	assert.Equal(t, errors.New("hll: can't merge HLLs with different number of registers"), hllInstance1.Merge(hllInstance2))
}

func TestMergeValid(t *testing.T) {
	numRegisterBits := 6
	hllInstance := NewHLLWithRegisterBits(numRegisterBits)
	hllInstance.registers[10] = 5
	other := NewHLLWithRegisterBits(numRegisterBits)
	other.registers[5] = 2
	assert.Nil(t, hllInstance.Merge(other))
	for idx, registerVal := range hllInstance.registers {
		if idx == 5 {
			assert.Equal(t, 2, registerVal)
		} else if idx == 10 {
			assert.Equal(t, 5, registerVal)
		} else {
			assert.Equal(t, 0, registerVal)
		}
	}
}

func TestGetAlphaByNumRegisters(t *testing.T) {
	assert.Equal(t, 0.673, getAlphaByNumRegisters(16))
	assert.Equal(t, 0.697, getAlphaByNumRegisters(32))
	assert.Equal(t, 0.709, getAlphaByNumRegisters(64))
	assert.Equal(t, 0.7213/(1+1.079/128.0), getAlphaByNumRegisters(128))
}
