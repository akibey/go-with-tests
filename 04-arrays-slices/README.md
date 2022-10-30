# Arrays and Slices

## Objectives

1. Write test to check for sum of an array and verify that it fails.
2. Write minimal code to get the test to run.
3. Write enough code to get the test to pass.
4. Refactor the code.
5. Write a test for sum of a slice.
6. Write minimal code to get the test to run.
7. Change the function to take slice as input.
8. Refactor the test to remove redundancy and check the test coverage.
10. Write test for SumAll function to take arbitrary number of slices as input and return the sum of each slice. 
11. Write minimal code to get the test to run.
12. As go doesn't let us compare two slices we use the DeepEqual function form reflect package.
13. Refactor the code to use append function to avoid runtime error if we excede slice capacity.
14. Write test for sumtails function to get sum of all elements of a slice except the first element.
15. Get the test to pass.
16. Make modifications for an empty slice.
17. Refactor the tests to remove repeated code.


## Test Cover

## reflect.DeepEqual
### **Note: DeepEqual is not type safe.** 
