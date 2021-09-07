benchstat: benchstat-masterminds benchstat-hashicorp

benchstat-hashicorp: ours.txt hashicorp.txt
	go run golang.org/x/perf/cmd/benchstat -geomean hashicorp.txt ours.txt

benchstat-masterminds: ours.txt masterminds.txt
	go run golang.org/x/perf/cmd/benchstat -geomean masterminds.txt ours.txt

BENCH=go test -bench . . | grep "^Bench"

ours.txt:
	touch $@
	cd ../ && $(BENCH) >> bench/$@
	cd ../ && $(BENCH) >> bench/$@
	cd ../ && $(BENCH) >> bench/$@
	cd ../ && $(BENCH) >> bench/$@
	cd ../ && $(BENCH) >> bench/$@
	cd ../ && $(BENCH) >> bench/$@
	cd ../ && $(BENCH) >> bench/$@
	cd ../ && $(BENCH) >> bench/$@
	cd ../ && $(BENCH) >> bench/$@
	cd ../ && $(BENCH) >> bench/$@

masterminds.txt:
	touch $@
	cd masterminds && $(BENCH) >> ../$@
	cd masterminds && $(BENCH) >> ../$@
	cd masterminds && $(BENCH) >> ../$@
	cd masterminds && $(BENCH) >> ../$@
	cd masterminds && $(BENCH) >> ../$@
	cd masterminds && $(BENCH) >> ../$@
	cd masterminds && $(BENCH) >> ../$@
	cd masterminds && $(BENCH) >> ../$@
	cd masterminds && $(BENCH) >> ../$@
	cd masterminds && $(BENCH) >> ../$@

hashicorp.txt:
	touch $@
	cd hashicorp && $(BENCH) >> ../$@
	cd hashicorp && $(BENCH) >> ../$@
	cd hashicorp && $(BENCH) >> ../$@
	cd hashicorp && $(BENCH) >> ../$@
	cd hashicorp && $(BENCH) >> ../$@
	cd hashicorp && $(BENCH) >> ../$@
	cd hashicorp && $(BENCH) >> ../$@
	cd hashicorp && $(BENCH) >> ../$@
	cd hashicorp && $(BENCH) >> ../$@
	cd hashicorp && $(BENCH) >> ../$@
