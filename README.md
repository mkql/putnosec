# putnosec

add `// #nosec Gxxx` directives according to `gosec` output json.

## Install

`go install`

## Usage

- `gosec -fmt=json <path> > gosec_output.json`
- `putnosec < gosec_output.json`
