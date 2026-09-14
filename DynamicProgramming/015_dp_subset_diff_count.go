package dynamicprogramming

/*
Find Total no of possible partitions of a array what becomes S1 - S2 = D
- Base Condition: If we reach our target then we return 1 return 0
- Explore Paths: Choose every element with alternative signs +ve or -ve
- Cache: key will be idx, curr as this is what our previous computations compute
*/
