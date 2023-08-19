# Go Notes



```bash
go env GOPATH

``` 

The go install command behaves almost identically to go build, but instead of leaving the executable in the current directory, or a directory specified by the -o flag, it places the executable into the $GOPATH/bin directory.


## SemVer
### Major
- Public API changes
- Backward incompatible  
- Update or Deleting  
eg. deleting or renaming a function, returns differnt type, no of parameters in the input


### Minor
- Public API changes
- Backward Compatible
- Adding 
- users can use new features only if they need it
eg. adding new feature.

### Patch
- Non Public api changes
- bug or security fixes
- no changes on how the  user interacts with any API




bug         start a bug report
build       compile packages and dependencies
clean       remove object files and cached files
doc         show documentation for package or symbol
env         print Go environment information
fix         update packages to use new APIs
fmt         gofmt (reformat) package sources
generate    generate Go files by processing source
get         add dependencies to current module and install them
install     compile and install packages and dependencies
list        list packages or modules
mod         module maintenance
work        workspace maintenance
run         compile and run Go program
test        test packages
tool        run specified go tool
version     print Go version
vet         report likely mistakes in packages



1. Go source Files can have only one package per directory  
2. A t.Errorf call is not an assertion. The test continues even after an error is logged  
3. When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.  



Adding to check git

After config settings
