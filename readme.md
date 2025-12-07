# Password generator

This is a password generator that creates passwords based on this criteria:
1. Passwords must not contain repeating characters.
2. Each password created must be unique (not repeated across multiple re-runs).
3. If the user has selected multiple character sets, the password must contain at least one character from each selected set.

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

## Flags:
- **--length** <integer> (required)
    Length of the password to generate. Must be a positive integer.
- **--numbers**
    Include numbers (0-9).
- **--lower**
    Include lowercase letters (a-z).
- **--upper**
    Include uppercase Latin letters (A-Z) in the password.

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

```

# Todo:

- [x] Create a simple password generation function.
- [x] Make symbols unique
- [x] Add input for length
- [x] Add option for sets (password should have one of each)
- [x] Make passwords unique
- [x] Dockerize
- [ ] Create a TUI
