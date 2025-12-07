# Password generator

## Criteria
This is a password generator that creates passwords based on this criteria:
1. Passwords must not contain repeating characters.
2. Each password created must be unique (not repeated across multiple re-runs).
3. If the user has selected multiple character sets, the password must contain at least one character from each selected set.

The 1. criteria makes 62 characters a maximum possible length of the password if user uses all the character sets.
Additionally according to 3. criteria if user selects all possible sets, the minimum length would be 3. One character for each set.

## Project structure
This project is of such a size that no folders are needed imho.Therefore all files are in the root of the project.

- `main.go` - The star of the show and start of the program. This mile contains an entry function along with: Input validation, file handling, error handling.
- `pass.go` - This file contains the algorithm. Main function being `GenPwd()`.
- `pass_test.go` - Contains tests for `GenPwd()`.

## Project development
This project has two feature branches. `gen-by-rune` and `shuffle`, which explored two different ways to randomize the password. After benchmarking the time, the `gen-by-rune` was selected into `main`.

`gen-by-rune` - This branch generates password step by step, rune by rune, until the necessary length is reached.

`shuffle` - Experiment branch that shuffles the whole character set and selects a slice of the whole string. Although the code is easier to read in this case, shuffling the whole character string while only selecting a part of the work done, turns out not to be such a great idea:)

# Instructions
## Start
1. Build the image
```
$ docker compose up
```

2. Run the program with all the necessary flags:
```
$ docker compose run go run . --length 30 --numbers --upper
```

3. Run tests
```
$ docker compose run go test
```

## Flags:
- **--length** <integer> (required)
    Length of the password to generate. Must be a positive integer.
- **--numbers**
    Include numbers (0-9).
- **--lower**
    Include lowercase letters (a-z).
- **--upper**
    Include uppercase Latin letters (A-Z) in the password.
- **--timeout** <integer>
    Time in seconds before timeout when retrying generation (default 30s). Must be positive.
- **--cleanup**
    Ignores all other flags and cleans the history of passwords.

## Example:
```
# 30 characters consisting of numbers and upper case letters
$ docker compose run go run . --length 30 --numbers --upper

# 30 characters consisting of numbers and lower case letters
$ docker compose run go run . --length 30 --numbers --lower

# 1 random number
$ docker compose run go run . --length 1 --numbers

# maximum length of the password
$ docker compose run go run . --length 62 --numbers --upper --lower

# delete the history
$ docker compose run go run . --cleanup

# increase the time to find the password
$ docker compose run go run . --length 20 --upper --timeout 60
```

# Todo:

- [x] Create a simple password generation function.
- [x] Make symbols unique
- [x] Add input for length
- [x] Add option for sets (password should have one of each)
- [x] Make passwords unique
- [x] Dockerize
- [ ] Create a TUI  <--  maybe some day :)
