# go-1brc

The 1 billion row challenge done in Go.

## Generate billion row file 

```sh
python3 create_measurements.py 1_000_000_000
```

## Run Tests 
```sh
time go test . -run 'TestMain$' -v
```

## Run with 1 billion rows
```sh
time go test . -run 'TestMain_1br$' -v
```

## Runs

### Run 1
```
--- PASS: TestMain_1br (397.62s)
PASS
ok  	github.com/bubunyo/go-1brc	398.460s
go test . -run 'TestMain_1br$' -v  199.66s user 31.35s system 57% cpu 6:38.70 total
```

### Run 2
```
--- PASS: TestMain_1br (335.76s)
PASS
ok  	github.com/bubunyo/go-1brc	336.844s
go test . -run 'TestMain_1br$' -v  197.98s user 26.06s system 66% cpu 5:37.45 total
```

### Run 3 

Channel length at 1000
```
--- PASS: TestMain_1br (188.02s)
PASS
ok  	github.com/bubunyo/go-1brc	188.125s
go test . -run 'TestMain_1br$' -v  342.19s user 15.30s system 189% cpu 3:08.33 total
```

### Run 4 

Channel length at 1000000
```
--- PASS: TestMain_1br (169.15s)
PASS
ok  	github.com/bubunyo/go-1brc	169.361s
go test . -run 'TestMain_1br$' -v  325.78s user 11.39s system 198% cpu 2:50.08 total
```

### Run 4 

Channel length at 100000
```
--- PASS: TestMain_1br (176.82s)
PASS
ok  	github.com/bubunyo/go-1brc	177.045s
go test . -run 'TestMain_1br$' -v  341.63s user 10.24s system 197% cpu 2:57.84 total
```

### Run 5 

```
--- PASS: TestMain_1br (294.38s)
PASS
ok  	github.com/bubunyo/go-1brc	294.694s
go test . -run 'TestMain_1br$' -v -cpuprofile profiles/cpu-1br-5.out  871.12s user 119.69s system 335% cpu 4:54.98 total
```

bucket size - 1000
```
--- PASS: TestMain_1br (214.48s)
PASS
ok  	github.com/bubunyo/go-1brc	215.031s
go test . -run 'TestMain_1br$' -v -cpuprofile profiles/cpu-1br-5.out  611.22s user 92.13s system 325% cpu 3:35.93 total
```

bucket size - 1
```
--- PASS: TestMain_1br (298.30s)
PASS
ok  	github.com/bubunyo/go-1brc	298.644s
go test . -run 'TestMain_1br$' -v -cpuprofile profiles/cpu-1br-5-1.out  506.32s user 108.02s system 205% cpu 4:59.49 total
```

bucket size - 3000
```
--- PASS: TestMain_1br (203.60s)
PASS
ok  	github.com/bubunyo/go-1brc	203.975s
go test . -run 'TestMain_1br$' -v -cpuprofile profiles/cpu-1br-5-3000.out  615.53s user 83.22s system 341% cpu 3:24.87 total
```


