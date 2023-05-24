# spam

Ever wanted a simple way to repeat a particular command at a particular cadence? `spam` runs any command you'd like at a set rate. 

This is mainly intended for scenarios where you'd like to generate synthetic traffic via `curl`, but can also work with any other program of your choosing.

## Usage

`spam [-r RATE_PER_SECOND] -- PROGRAM_ARGS`

## Examples

Send 10 req/s to www.google.com

```bash
spam -r 10 -- curl www.google.com
```

Send 1 req/s to localhost:8080

```bash
spam -- curl localhost:8080
```

Add a 10 "hello" lines to a particular file per second

```bash
spam -r 10 -- bash -c 'echo "hello" >> ./text.txt'
```

## Installation

### Mac OS

You can use `brew tap` to install:

```bash
brew tap Steven-Ireland/homebrew-tap
brew install spam
```

### Linux

Head over to [Releases](https://github.com/Steven-Ireland/spam/releases/) and download the latest release binary. You'll need to place it on your PATH manually.

## See Also

This program outputs durations in milliseconds for each run -- you can pipe the output to `asciigraph` for a nice visual moving representation:

```bash
spam -r 10 -- curl www.google.com | asciigraph -r -h 10 -w 40
```
![image](https://github.com/Steven-Ireland/spam/assets/6981727/75737ef7-f807-4ff2-891e-e8aaba37a135)

## Questions / Contributions

Feel free to open an issue or submit a PR if anything comes up. 
It's a pretty tiny program at the moment, so larger features may be left open to community contributions.
