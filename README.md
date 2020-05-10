# Glean
Glean is a web server sample for an app server with Go.

This code structure aims for a structure that:
1. Easy to maintain perpetually.
2. Testable, with specific func responsibility.
3. Works when the code is outlive the engineer life span on the team.
4. Support highly mutable business requirements. Since we know that business is never converge and it is only diverge linearly as the business grows.
5. Work with monolith project or macro service.


Note: This is not a dogmatic approach and please use your own judgement for your own project. 

# Folder Structure

## internal
This folder used to reduce our source code being imported since the code you build may only works within this project context.

## server
This package will handle your server initialization. For this structure, http routes also placed in here.
In case of using GRPC, the server initialization can be placed here.

## entity
This folder will contain all the type that going to support each domain in repository, usecase and controller

## controller
This layer will act as a presenter for deliver your API response.
In this package we can define API support for HTTP REST, GRPC and/or HTML.
This layer is going to accept input from client based on what kind of protocol being used by the client.
The validation of client request happen in this layer such as required parameters, formatting, etc.
Controller layer will call usecase layer for processing the client request and response to client.

## usecase
Use case Layer
This layer will act as a logic processor for controller layer and repository layer.
In this package we will do any kind of business logic, from calling func that query to database, do the caching and
communicate with different domain.
Communicate with different domain also happen in this layer by injecting other domain usecase to
the designated domain usecase.

## repository
This layer act as data layer for the project. It handle database queries, redis cache and HTTP call
to the other microservice. It should not do any business logic since all the value already handled by use case
layer. Simply CRUD to DB or http call, so the dependencies to this layer would be implementation of DB connection,
redis connection, or http client instance.

## pkg
This folder is containing any helper package that you made. The helper/wrapper in this `pkg` is not placed inside `internal` 
with the purpose to made all the packages here are reusable by the other project. So the helper package that you made suppose to be standalone here.

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
// With this, the Something func don't have to know how abstract being implemented.
func Something(a abstract) *SomeType {
    return &SomeType{
         dependencyOne: a,
    }
}

// Do is method from the concrete type that call a method from abstraction.
func (s *SomeType) Do() error {
    return s.dependencyOne.AbstractMethod(1, "a")
}
```

# Personal Opinion
Based on my experience, the knowledge sharing for the new commer to the team is relaively simple, since layering structure is a common practice
in another programming language. What they really need to have before learning the structure works is how `interface{}` works in golang.

For maintaining and development at the same time with few peoples contribute to the same project is not a difficult stuff since they have the same
understanding of how the structure works and where to add or remove stuff.

Creating unit test was easy, and the unit test can be really specific to each layer. This help my team to understand how a particular unit works as well.

The issue that we are facing during after the project getting so big are:
1. Nested folder. It is started to make us scrolling a lot in the folder explorer. This issue can be easily handled by your favorite editor / IDE in my opinion. But still, those scrolling is taking time.
2. A lot of `interface{}` type. This is happen since the structure aiming for abstraction on each layer for easy to mock. Most of the time the abstraction can be ignored
once it was added.
3. If your microservice responsibility is very small until you don't have to separate any domain inside of it. This structure is just overkill (IMHO).

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
