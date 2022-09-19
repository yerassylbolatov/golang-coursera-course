curl http://localhost:8080/debug/pprof/heap -o mem_out.txt
curl http://localhost:8080/debug/pprof/profile?seconds=5 -o cpu_out.txt

go tool pprof -svg -alloc_objects ./pprof_1 mem_out.txt > mem.svg
go tool pprof -svg ./pprof_1 cpu_out.txt > cpu.svg