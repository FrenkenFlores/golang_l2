# Design Pattern: Facade
## Usage
The Facade design pattern is used when there are multiple subsystems that need to be interacted with. This pattern involves creating a unified interface that encapsulates the interactions with the subsystem interfaces, providing a simplified way to access their functionality.

## Pros
- Improves readability: By providing a single, high-level interface, the Facade pattern makes the code easier to understand and use.

- Promotes separation of concerns: It decouples client code from the complexities of the subsystems, making the codebase more modular and maintainable.

- Simplifies testing: With a Facade, testing becomes easier as the interactions with subsystems are centralized, allowing for more focused and isolated tests.

## Cons
- Limits customization: The Facade pattern can make it harder to implement custom behavior, as it abstracts away the details of the subsystems.

- Risk of becoming a "god object": If not designed carefully, the Facade can grow too large and take on too many responsibilities, violating the Single Responsibility Principle.