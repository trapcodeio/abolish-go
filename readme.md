# Abolish GoLang Version

## What is this?
An attempt to make a GoLang version of the Abolish project.

## Why?

### Concept
The current validators in GoLang have some downsides that i cannot change.
The concept of abolish will provide a better way to validate data.

## Downsides
Most validators do reflect at runtime, which is not a good thing.
abolish will not do this, which will make it faster at runtime.

No way to define errors at runtime. 😩

## Solution
Compile functions will be provided to compile rules and generate a single function to use at runtime.

## How?
The abolish package will provide a way to validate data in a more flexible way.

```go
// rules via string
var rule = abolish.StringToRules("required|string|min:3|max:10")

// rules via maps
var ruleMap = map[string]any{
    "required": true,
    "string":   true,
    "min":      3,
    "max":      10,
}

// rules via structs
type Rule struct {
	// Note Tag `abolish` is customisable
    Email string `abolish:"required|email"`
}
```

## When?
I am not sure when this will be finished, but i will try to make it as soon as possible.
For now it will only progress as i use it.