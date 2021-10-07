# simUduck
## Uses Strategy Pattern:
The strategy pattern defines a family of algorithms, encapsulates each one and makes them interchangeable.
Strategy lets the algorithm vary independently from the clients that use it.

## Here
Here, flyWithWings, flyWithRocket, quackMute, quackNormal are various algorithms. These are divided into two families -> fly and quack. All of these algorithms implement the interface for their respective classes. The duck class composes (Has-a) fly and quack behaviour in it. Now, each duckType (which itself extends duck) can have whatever flyBehaviour or quackBehaviour it wants. At the same time, we can change any particular algorithm in one file and see changes accross the program, i.e. easier to maintain. We can add new behaviours without breaking old ducktypes, i.e. extensibility. As an added advantage we can change the flyBehaviour at runtime as well which provides more room to play with. 
Thus, for code reuse, inheritance is not the only way and not always the best way. This example shows how composition of family of algorithm, i.e. Strategy Pattern works.