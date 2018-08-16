package hll

type hll struct {
  registers []int
}

func New(numRegisters int) hll {
  registers := make([]int, numRegisters)
  hllInstance := hll {registers}
  return hllInstance
}

func add() {

}

func count() {

}

func merge() {

}
