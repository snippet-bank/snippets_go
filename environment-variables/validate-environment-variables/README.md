Retrieving the values of environment variables and validating them is a common problem.

`github.com/kelseyhightower/envconfig` provides an easy and structured way of doing both of that.
 
It allows you to provide a configuration struct in which you can set:
- expected names of the variables
- expected types of values (e.g. `int` or `[]string`)
- whether a variable is required

More information available at [godoc](https://godoc.org/github.com/kelseyhightower/envconfig).
