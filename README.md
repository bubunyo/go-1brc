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
