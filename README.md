<img src="https://procyon-projects.github.io/img/logo.png" width="128">

# Procyon Core
![alt text](https://goreportcard.com/badge/github.com/procyon-projects/procyon-core)
[![codecov](https://codecov.io/gh/procyon-projects/procyon-core/branch/master/graph/badge.svg?token=F9WA517EG9)](https://codecov.io/gh/procyon-projects/procyon-core)
[![Build Status](https://travis-ci.com/procyon-projects/procyon-core.svg?branch=master)](https://travis-ci.com/procyon-projects/procyon-core)


This gives you a basic understanding of Procyon Core Module.

## Components
All instances managed by the framework such as Controller, Initializers and Processors are considered as a component. 

### Register Component
It's used to register the components. However, you can pass only a construction-function into it.
Construction function means a function returning only one type but it can have multiple parameters.
```go
func Register(components ...Component)
```
The example is given below.

```go
type MyComponent struct {

}

func NewMyComponent() MyComponent {
    return MyComponent{}
}

func init() {
    core.Register(NewMyComponent) 
}
```

### Component Processor
This allows you to process the components registered before application starts.
```go
type ComponentProcessor interface {
	SupportsComponent(typ goo.Type) bool
	ProcessComponent(typ goo.Type) error
}
```
* **SupportsComponent** if it returns true, it means that the component will be processed by this 
component processor.
* **ProcessComponent** If the method **SupportComponent** returns true, it will be invoked. 

The example of a custom component processor is given below.

```go
type CustomComponentProcessor struct {
}

func NewCustomComponentProcessor() CustomComponentProcessor {
	return CustomComponentProcessor{}
}

func (processor CustomComponentProcessor) SupportsComponent(typ goo.Type) bool {
	// do whatever you want
    return false
}

func (processor CustomComponentProcessor) ProcessComponent(typ goo.Type) error {
	// do whatever you want
    return nil
}
```

Note that you need to register the component processors by using the function **core.Register**.

## License
Procyon Framework is released under version 2.0 of the Apache License
