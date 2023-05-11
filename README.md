# InScope

This is a simple command that filters a list of domain names based on a list of allowed scope domains.

```sh
# Install
go install -v github.com/esidate/inscope@latest

# Manual install
# Download
wget https://github.com/esidate/inscope/releases/download/v0.1.0/inscope
# Or build
go build -o inscope
# And install
mv inscope /usr/local/bin

# Usage
cat domains.txt | inscope -scope scope.txt
# Or
inscope -scope scope.txt domains.txt
```
