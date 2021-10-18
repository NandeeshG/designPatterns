# Design Patterns

## Principles
1. Encapsulate what varies
2. Favor composition over inheritance
3. Program to interface not implementations
4. Strive for loosely coupled design between objects that interact
5. Classes should be open for extension but closed for modification

## Patterns
1. Strategy Pattern (simuDuck)
2. Observer Pattern (weatherStation)
3. Decorator Pattern (starbuzz)

## Principles Explained
5. Classes should be open for extension but closed for modification
- This means we can extend classes and add new behaviours but we are not allowed to touch pre-written code, it is closed for modification.
- Using inheritance determines behaviours statically at compile time. Using composition allows dynamic behaviour switch like we see in FlyingBehaviour of ducks. Using compositions we can perform decorator pattern-esque implementations and add new responsibilities to objects which were not present in original superclass. This also avoids altering existing code to add new functionalities.

### That's all!
![Feature Creep](https://static1.smartbear.co/smartbear/media/blog/wp/frd%20feature%20creep.png)
