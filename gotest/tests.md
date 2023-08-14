`go test .` or `got test ./folder-name/` to run all tests  
`go test -run TestArea/Rectangle ./structs/` to run only subtest  


### coverage
[The Cover Story](https://go.dev/blog/cover)  
`go test -cover .` finds test coverage  
   
  
 ask go test to write a “coverage profile” for us, a file that holds the collected statistics  

`go test -coverprofile=coverage.out`  
from the `coverage.out` file we can call two commands.  
1. `go tool cover -func=coverage.out`  Split the Coverage percentage by function wise 
2. `go tool cover -html=coverage.out`  Generate the html to view the coverage 







### Lessons learned in TDD
https://agilewarrior.wordpress.com/2016/02/16/lessons-learned-in-tdd/

- Not writing tests against behaviors.
- Coupling our implementation details to our tests.


The trigger for writing a new test is to have some piece of behavior that 
I want to implement. The test needs to capture that behavior.  



