# Concept Basic Architecture

```text
          ┌─────────────────┐
          │      main       │
          │   Composition   │
          │      Root       │
          └────────┬────────┘
                   │
                   │ inject
                   ▼
          ┌─────────────────┐
          │    Calculator   │
          │    Service      │
          └────────┬────────┘
                   │
                   │ depends on
                   ▼
          ┌─────────────────┐
          │    Operation    │
          │    Interface    │
          └────────┬────────┘
                   │
      ┌────────────┼────────────┐
      ▼            ▼            ▼
Addition     Subtraction   Multiplication
```

# Flow Dependency Injection

```text
main()
  |
  | creates
  v
Addition{}
  |
  | inject
  v
NewCalculator(addition)
  |
  v
Calculator
  |
  | Calculate(10, 5)
  v
Operation.Apply()
  |
  v
Addition.Apply()
  |
  v
15
```