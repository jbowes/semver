benchstat: ours.txt masterminds.txt
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
	$(BENCH) >> $@
	$(BENCH) >> $@
	$(BENCH) >> $@
	$(BENCH) >> $@
	$(BENCH) >> $@
	$(BENCH) >> $@
	$(BENCH) >> $@
	$(BENCH) >> $@
	$(BENCH) >> $@
	$(BENCH) >> $@