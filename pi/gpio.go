package pi

import (
	"os"
	"reflect"
	"syscall"
	"unsafe"
)

// Mode should have a comment
type Mode uint8

const (
	gpioBase = 0x200000

	blockSize = 4 * 1024
)

// Pin mode
const (
	Input Mode = iota
	Output
)

var (
	moduleSetup = false

	gpio8 []byte
	gpio  []uint32
)

// SetupGpio should have a comment
func SetupGpio() (err error) {
	var file *os.File

	if moduleSetup {
		return nil
	}

	moduleSetup = true

	file, err = os.OpenFile("/dev/gpiomem", os.O_RDWR|os.O_SYNC, 0)
	if err != nil {
		return
	}

	gpio8, err = syscall.Mmap(int(file.Fd()), gpioBase, blockSize, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return
	}
	gpio = convertToArray32(gpio8)

	return nil
}

// OpenPin should have a comment
func OpenPin(num int, mode Mode) {
	setupCheck()

}

func setupCheck() {
	if !moduleSetup {
		panic("You have not called one of the SetupGpio functions")
	}
}

func convertToArray32(arr8 []byte) []uint32 {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&arr8))
	header.Len /= (32 / 8) // (32 bit = 4 bytes)
	header.Cap /= (32 / 8)
	return *(*[]uint32)(unsafe.Pointer(&header))
}
