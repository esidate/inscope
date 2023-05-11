# InScope

This is a simple command that filters a list of domain names based on a list of allowed scope domains.

```sh
# Clone the project and build it
go build -o inscope
mv inscope /usr/local/bin

# Usage
cat domains.txt | inscope -scope scope.txt
# Or
inscope -scope scope.txt domains.txt
```