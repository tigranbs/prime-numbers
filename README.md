# Database for first 50 Mln Prime Numbers
Generating Prime numbers could be very challenging and time consuming. And it turns out that there is 
publicly available generated first 50 Mln prime numbers here [https://primes.utm.edu/lists/small/millions/](https://primes.utm.edu/lists/small/millions/).

This repository is aims to make already generated prime numbers available as a programming interface 
for different programming languages.

# Getting numbers
`prime_grabber.go` file is for downloading text files from public website, parsing them, and encoding numbers to
`4 bytes BigEndian`, which is giving very easy representation for reading numbers from different programming languages.
Also it is very efficient to keep all numbers as a 4 bytes data representation. Later it would be easy to 
navigate to specific prime number.
```bash
~# go run prime_grabber.go
# This will take some time (~2-5 min), based on internet connectivity
```
After running this command you will have BigEndian encoded prime numbers in `primes.data` (~200MB) file in current directory.

# Contribution
`primes.data` file format is very simple, it is just a sequence of 4 byte BigEndian numbers. So it is very easy to add
`N'th prime number` functionality to any programming language.

So you need just to generate that files using `prime_grabber.go` script and start using it. Send me `Pull Requests` if you 
did some implementation for any other programming language!