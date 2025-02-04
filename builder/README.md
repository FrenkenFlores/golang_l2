# Design Pattern: Builder
## Usage
The Builder design pattern is used when the created object requires multiple optional parameters and allows various configurations. This pattern enables object creation step by step using chained function calls to configure its parameters.

## Pros
- Isolates complex logic
- Improves code readability
- Allows different versions of the same object
## Cons
- If the created object has only a few parameters, using the Builder pattern may make the code harder to read.