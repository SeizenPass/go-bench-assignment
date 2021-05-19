gen:
	go test -bench . -benchmem -cpuprofile=cpu.out -memprofile=mem.out -memprofilerate=1
mem:
	go tool pprof hw3_bench.test.exe mem.out
cpu:
	go tool pprof hw3_bench.test.exe cpu.out
json:
	easyjson -all user/user.go
