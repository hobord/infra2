## Stress test

https://github.com/tsenart/vegeta

```
echo "GET http://127.0.0.1:8080/" | vegeta attack -duration=60s -workers=250 | tee results.bin | vegeta report
  vegeta report -type=json results.bin > metrics.json
  cat results.bin | vegeta plot > plot.html
  cat results.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"
```

```
echo "GET http://127.0.0.1:8080/test" | vegeta attack -duration=60s -workers=250 | tee results.bin | vegeta report
  vegeta report -type=json results.bin > metrics.json
  cat results.bin | vegeta plot > plot.html
  cat results.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"
```

```
echo 'GET http://localhost:8080' | \
    vegeta attack -rate 5000 -duration 10m | vegeta encode | \
    jaggr @count=rps \
          hist\[100,200,300,400,500\]:code \
          p25,p50,p95:latency \
          sum:bytes_in \
          sum:bytes_out | \
    jplot rps+code.hist.100+code.hist.200+code.hist.300+code.hist.400+code.hist.500 \
          latency.p95+latency.p50+latency.p25 \
          bytes_in.sum+bytes_out.sum
```

```
cd reports
echo "GET http://10.20.35.111:30004/" | vegeta attack -duration=60s -rate=500/s -workers=1000 | tee results.bin | vegeta report
  vegeta report -type=json results.bin > metrics.json
  cat results.bin | vegeta plot > plot.html
  cat results.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"


Requests      [total, rate, throughput]  30000, 500.02, 500.00
Duration      [total, attack, wait]      1m0.000084823s, 59.998191509s, 1.893314ms
Latencies     [mean, 50, 95, 99, max]    2.778777ms, 2.221414ms, 5.718059ms, 11.029322ms, 32.701488ms
Bytes In      [total, mean]              216840000, 7228.00
Bytes Out     [total, mean]              0, 0.00
Success       [ratio]                    100.00%
Status Codes  [code:count]               200:30000  
Error Set:
/reports #   vegeta report -type=json results.bin > metrics.json
/reports #   cat results.bin | vegeta plot > plot.html
/reports #   cat results.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"
Bucket           #      %        Histogram
[0s,     100ms]  30000  100.00%  ###########################################################################
[100ms,  200ms]  0      0.00%    
[200ms,  300ms]  0      0.00%    
[300ms,  +Inf]   0      0.00%  

```