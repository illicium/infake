# Infake

Generate fake InfluxDB metrics

## Setup

Requirements:
    
* Go 1.6+
* glide (OS X: `brew install glide`)

## Usage

    glide install

    go install ./infake

Then run:

    infake --config ./config.sample.yaml

Or, without log output:

    infake --config ./config.sample.yaml 2>/dev/null
