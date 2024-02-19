# ML Data Gen Demo

## Setup

1. Ensure `go` is installed.
2. Ensure python3 is installed. This demo was tested with `3.11.x`
3. To run the Go worker, simply run: `go run worker/main.go`
4. To run the Python worker:

```
python3 -m venv env
source env/bin/activate
pip3 install -r requirements.txt
python3 worker.py
```

4. To invoke the Train workflow, run: `go run train-starter/main.go`
5. To invoke the Sample workflow, run: `go run sample-starter/main.go`

Once completed, there should now be records in the `test-stage-db` database.
