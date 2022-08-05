

all:
	easyjson -all fast.go
	go test -bench . -benchmem -cpuprofile=cpu.out -memprofile=mem.out -memprofilerate=1 main_test.go common.go fast.go fast_easyjson.go

mem:
	go tool pprof mem.out

cpu:
	go tool pprof cpu.out

test:
	go test -v