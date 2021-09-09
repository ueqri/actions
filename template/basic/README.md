# Basic template for configuration of simple benchmarks

## Configuration

- `runner.yaml` contains 3 main sections, that is **template**, **override**, **runner**.
  * **template** is an array in YAML scheme, used to clarify which benchmarks we need to use from template list, e.g. `actions.json` in this folder.
  * **override** is used for experimentation of different types of parameters of a certain benchmark defined in template list, in order not to rewrite or copy too much when switching to a new investigation.
  * **global** is a field specified for global runner config, for example, the number of ranks we would use in total, is defined as `runner` in this field.
- `actions.json` is derived from the scheme in [Automated Benchmark Runner for akkalat](https://github.com/sarchlab/akkalat/wiki/Automated-Benchmark-Runner#benchmark-configuration). The reason why we use two types of configs is that, the config varies from experimentation to experimentation, but we could have one typical template list, i.e., `actions.json`, which covers almost all benchmarks with default or recommended configs.

## Build & Run

```bash
cd /path/to/actions/template/basic
go build
./basic -config=runner.yaml -template=actions.json -rank=0
```
