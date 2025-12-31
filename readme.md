# Password generator

## About
This is a password generator that creates passwords based on this criteria:
1. Passwords must not contain repeating characters.
2. Each password created must be unique (not repeated across multiple re-runs).
3. If the user has selected multiple character sets, the password must contain at least one character from each selected set.

The 1. criteria makes 62 characters a maximum possible length of the password if user uses all the character sets.
Additionally according to 3. criteria if user selects all possible sets, the minimum length would be 3. One character for each set.

## Project structure
This project is of such a size that no folders are needed imho.Therefore all files are in the root of the project.

- cmd/password_gen/main.go - The star of the program. Handles CLI flags, validates input, and calls internal packages.

- internal/pass/pass.go - Core password generation algorithm (GenPwd). Implements character set selection, uniqueness constraints, and shuffling.

- internal/pass/pass_test.go - Unit tests for the password generation logic.

- internal/history/file.go - Manages the password history file: reading existing hashes and saving new ones.

- internal/history/genUnique.go - Contains TryUniquePassword, which generates passwords and checks them against history to ensure uniqueness with timeout handling.

## Project development
This project has two feature branches. `gen-by-rune` and `shuffle`, which explored two different ways to randomize the password. After benchmarking the time, the `gen-by-rune` was selected into `main`.

`gen-by-rune` - This branch generates password step by step, rune by rune, until the necessary length is reached.

`shuffle` - Experiment branch that shuffles the whole character set and selects a slice of the whole string. Although the code is easier to read in this case, shuffling the whole character string while only selecting a part of the work done, turns out not to be such a great idea:)

# Instructions
## Quick Start
1. Build the image
```bash
$ make build
```

2. Run the program with all the necessary flags:
```bash
$ make run FLAGS="--length 10 --numbers --lower"
OR
$ make example
```

3. Run tests
```bash
$ make test
```

### Other commands from the make file
``` bash
Usage: 
help:    Show help
build:   Build the container
run:     Run the password generator (usage: make run FLAGS="--length 30 --numbers")
example: Run with example flags
cleanup: Delete history
test:    Run tests
clean:   Clean up resources
```

## Flags:
- **--length** <integer> (required)
    Length of the password to generate. Must be a positive integer.
- **--numbers**
    Include numbers (0-9).
- **--lower**
    Include lowercase letters (a-z).
- **--upper**
    Include uppercase Latin letters (A-Z).
- **--timeout** <integer>
    Time in seconds before timeout when retrying generation (default 30s). Must be positive.
- **--cleanup**
    Ignores all other flags and cleans the history of passwords.

## Examples

```bash
# 30 chars, numbers + uppercase
make run FLAGS="--length 30 --numbers --upper"

# 1 random number
make run FLAGS="--length 1 --numbers"

# Max length (62 chars)
make run FLAGS="--length 62 --numbers --upper --lower"

# Clean history
make cleanup
```

# Todo:

- [x] Create a simple password generation function.
- [x] Make symbols unique
- [x] Add input for length
- [x] Add option for sets (password should have one of each)
- [x] Make passwords unique
- [x] Dockerize
- [ ] Create a TUI  <--  maybe some day :)
