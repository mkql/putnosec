# putnosec

add `// #nosec Gxxx` directives according to `gosec` output json.

## Install

`go install github.com/mkql/putnosec:latest`

## Usage

- `gosec -fmt=json <path> > gosec_output.json`
- `putnosec -w < gosec_output.json` 
  - This command overwrites source files with gosec issues. Without -w option, putnosec prints planned change and exit.

You can see other options with `putnosec -help`.