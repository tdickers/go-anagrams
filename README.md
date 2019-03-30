# go-anagrams

One-off puzzle solver (for a specific puzzle) and an excuse to try Go.

The original puzzle was: `BASHLACK MUSTRANT NOSYSHIP`.

# anagrams
Bruce-force solver using permuations.

```
go get modernc.org/mathutil
go build anagrams.go
./anagrams <anagram 1> <anagram 2>
```

# better-anagrams
Simpler solution without an external dependency and likely better performance for longer words.

```
go build better-anagrams.go
./better-anagrams <anagram 1> <anagram 2>
```

# Performance

Measured on a low-powered Chromebook using `time`. Measuring more points in between
 and getting a graph would be a nice future test. (Currently the Chromebook doesn't
 even have the memory to keep the editor, terminal, and browser open, so switching
 between is painful.)

## 8-character anagrams

```
./anagrams BASHLACK MUSTRANT NOSYSHIP  0.65s user 0.19s system 100% cpu 0.839 total
./better-anagrams BASHLACK MUSTRANT NOSYSHIP  4.88s user 0.45s system 124% cpu 4.290 total
```

Original `anagrams` is 5.1x faster.

## 15-character anagrams

```
./anagrams salternativenes canthropomorphi sparticularitie # > 1h. Did not finish (this is a Chromebook, after all)
./better-anagrams salternativenes canthropomorphi sparticularitie  5.64s user 0.98s system 114% cpu 5.764 total
```
