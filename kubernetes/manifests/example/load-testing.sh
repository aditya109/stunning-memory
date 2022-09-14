echo "GET http://192.168.49.2:31504/ping" | vegeta attack -duration=4m -rate=200 | tee results.bin | vegeta report
vegeta report -type=json results.bin > metrics.json
cat results.bin | vegeta plot > plot.html
cat results.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"
