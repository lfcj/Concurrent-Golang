# Concurrent-Golang
Golang and concurrent Programming

[Smokers Problem Solution]

The smokers problem describes a situation in which 3 smokers sit at a table. Every one of them has tobacco, smoking paper, and matches, respectively, but do not share with each other. 
The waiter comes to the table and puts 2 smoking utensils on it. The smoker who has the complementary utensil can smoke now. He does it, stops, and then all 3 wait again for the waiter.

[Smokers Problem Solution]: https://github.com/lfcj/Concurrent-Golang/blob/master/smokersProblem.go

[Concurrently multiplication of matrixes]

The multiplication of 2 matrixes succeds concurrently. The submultiplication of every row by a columb is done with a go routine. When done, it waits for the others at a Barrier. 

[Concurrently multiplication of matrixes]: https://github.com/lfcj/Concurrent-Golang/blob/master/matrixMultiplicationConcurrent.go
