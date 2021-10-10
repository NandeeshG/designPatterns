# Starbuzz

## Decorator Design Pattern
Attach additional responsibilities to an object dynamically. Decoration provides a flexible alternative to subclassing for functionality extension.
**Principle**: Classes should be open for extension but closed for modification


## Here
Beverages are concrete classes, on which we can decorate with various condiments. These condiments are similar typed (beverageInterface) and also compose a beverage in them. We can add as many condiments as we want. All of these wrap a beverageInterface in them, allowing for a recursive getCost() and getDescription() resolution. Adding new condiments is as easy as extending Condiments to a new class. Adding new beverages is also easy. 

### Class Diagram
![Class Diagram](/starbuzz/Diagram.png)