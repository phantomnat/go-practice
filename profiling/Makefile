
.PHONY: concat-string
concat-string:
	go test -bench="BenchmarkConcat.*" -benchmem

.PHONY: strconv
strconv:
	go test -bench="BenchmarkStrconv.*" -benchmem


.PHONY: regexp
regexp:
	go test -bench="BenchmarkRegexp.*" -benchmem

.PHONY: sort
sort:
	go test -bench="BenchmarkSort.*" -benchmem

.PHONY: rand
rand:
	go test -bench="BenchmarkRand.*" -benchmem

.PHONY: http
http:
	go test -bench="BenchmarkHTTP.*" -benchmem
