# Infake

Generate fake InfluxDB metrics

## Setup

Requirements:

* Go 1.6+
* [bazel](https://bazel.build/) 0.16.0+

## Usage

### Build

    bazel build //infake

    # Binary will be in bazel-bin/infake/*/infake

### Run

With a config file:

    infake --config ./config.sample.yaml

Or, without log output:

    infake --config ./config.sample.yaml 2>/dev/null

You can also run directly with `bazel run`:

    bazel run //infake -- --config $PWD/config.sample.yaml
