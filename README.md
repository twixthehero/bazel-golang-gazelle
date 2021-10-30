# bazel-golang-gazelle

Example Bazel workspace configured to build a Golang program binary as well as manage BUILD
files with Gazelle.

## Usage

### Running

```bash
$ bazel run //example
```

### Adding new Golang dependencies

1. Add the new `go.mod` dependency
2. Update `go.sum`
3. Update the bazel dependencies
4. Update BUILD files

   ```bash
   $ cd example
   $ go get <repo>                              # 1
   $ go mod tidy                                # 2
   $ bazel run //example:gazelle-update-repos   # 3
   $ bazel run //example:gazelle                # 4
   ```

# License

MIT
