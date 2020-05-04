# Glean
Glean is a web server sample that I had used for more than 3 years writing app server with Go.
This code structure aims for a structure that is easy to maintain, testable, 
works when the code is outlive the engineer life span on the team.
I personally do not find any critical issue in maintaining the code base perpetually.

Note: This is not a dogmatic approach and please use your own judgement for your own project. 

# Folder Structure

# Common Pattern
Controller layer and usecase layer has a common pattern in this project structure. It has a struct (concrete type) and 
interface (abstract type). The struct is used to containing the dependencies of the layer and defining the receiver 
methods. 

The interfaces are used to abstract away the dependencies.  
Purpose of doing this are:
1. Decoupled implementation from certain package
2. Easily switch between implementations or provide multiple one
3. Easily mocking the implementation so unit test can be more specific on the func, not the detail implementation of its
dependencies.

Sample:
```go
package something

// abstract is abstraction type that will help the package: 
// 1. generate mock with the help of gomock
// 2. Defining the abstraction at the point of use
type abstract interface {
    // AbstractMethod being defined at the point of use.
    // This method can easily mocked with gomock.
    AbstractMethod(p1 int, p2 string) error
}

// SomeType is a concrete type that containing abstraction type
type SomeType struct {
    dependencyOne abstract
}

// Something creates SomeType struct and injecting the dependencies to it. 
func Something(a abstract) *SomeType {
    return &SomeType{
         dependencyOne: a,
    }
}
// Do is method from the concrete type
func (s *SomeType) Do() error {
    return s.dependencyOne.AbstractMethod(1, "a")
}
```

# Reference
[Industrial Programming](https://peter.bourgon.org/go-for-industrial-programming/)
```
"Said another way, interfaces are consumer contracts, not producer (implementor) contracts—so, as a rule, 
we should be defining them at call sites in consuming code, rather than in the packages that provide implementations." 
- Peter Bourgon
```

[Understanding Interface](https://youtu.be/F4wUrj6pmSI)
```
"Return concrete types, receive interfaces as parameters" - Francesc Campoy
```

