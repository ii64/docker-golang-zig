# golang-zig

Golang with Zig

```bash
make gen TAG=1.18-alpine3.15
make build ZIGVER=0.10.0-dev.2851+f639cb33a \
    IMAGENAME=golang-zig:go1.18-alpine3.15-zig

# publish
make publish IMAGENAME=golang-zig:go1.18-alpine3.15-zig

# test
make test-run IMAGENAME=golang-zig:go1.18-alpine3.15-zig
```

Thanks to `zig cc` and `zig c++`