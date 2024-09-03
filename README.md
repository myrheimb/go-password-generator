# go-password-generator
A simple password generator written in Go.

## Prerequisites
If you have Go installed globally, simply run it as shown in [Usage](#usage).

If you have Nix, devenv, and direnv installed, `cd` into the project folder and run `direnv allow` .
- [Nix](https://nixos.org/download#download-nix)
- [devenv](https://devenv.sh/getting-started)
- [direnv](https://direnv.net/docs/installation.html)

## Usage
Run the program in one of the following two ways:

1. `go run main.go`.
2. First run `go build` then run the binary `./go-password-generator`.

**Example 1**: Password with lower and upper case letters, numbers, and special characters
```
> go run main.go -length 16
Generated Password: hv$7EwaAitJN!UgL
```

**Example 2**: Password without special characters
```
> go run main.go -length 20 -types lun
Generated Password: fIMFiaQt1qmRcPZ2OIM7
```

**Example 3**: PIN code
```
> go run main.go -length 4 -types n
Generated Password: 7115
```

**Help message**
```
> go run main.go -help
Usage of go-password-generator:
  -length int
    	Length of the password (default "12")
  -types string
    	Character types to include (l=lower, u=upper, n=numbers, s=special) (default "luns")
```

## Contribution
Feel free to open an issue or create a Pull Request if you have suggestions for improving the password generator!