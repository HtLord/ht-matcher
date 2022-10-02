Ht-Matcher
===

Goal
---

Simulating stock exchange dealing orders match up behavior which sold or bough with the 
same price. Display how deal been matched with each other.

### E.g.

### Inputs(order type will not be implemented)

| price type | price | order type | quantity | sequence  | status | time |
|------------|---|--------|----------|------------|-------|---|
| same       |5.1|buy|10|1|neutral|*|
| same       |5.2|buy|10|2|neutral|*|
| same       |5.3|buy|10|3|neutral|*|
| same       |5.1|sell|10|4|neutral|*|
| same       |5.1|buy|10|5|neutral|*|
| same       |5.1|sell|70|6|neutral|*|

### Steps

1. seq 1 may killed
2. seq 2 may killed
3. seq 3 may killed
4. seq 4 will filled with seq 1
5. seq 5 may killed
6. seq 6 may killed
7. Mark none filled orders to killed
8. Export filled order


### Outputs
| price type | price | order type | quantity | sequence  | status  | time |
|------------|---|--------|----------|------------|---------|---|
| same       |5.1|buy|10|1| filled  |*|
| same       |5.2|buy|10|2| killed |*|
| same       |5.3|buy|10|3| killed |*|
| same       |5.1|sell|10|4| filled |*|
| same       |5.1|buy|10|5| killed |*|
| same       |5.1|sell|70|6| killed |*|

Think
---

1. Define a window as maximum system contained order for simulating paused queue
2. Running matching logic
3. Export matched sell records
4. Concurrent will take place at input and output lock the data status from database or somewhere record the order
status


Question
---

> Q: What if there is only `so` or `po` in queue
> <br />
> A: I think according to FOK the whole transaction will all killed
   
> Q: What is the matching logic? What if there encounter partial filled need rollback or what?
> <br />
> A: There can be lots of strategy. But it is not being defined clearly in spec. Then I will assume make up my own. <br/>
> 1. Get an order from window will be called, check order. <br/>
> 2. Try to find possible combination from remained window orders if there is fit price and status are not `filled` then
> all of them will be mark as filled then move on to next order as check order from remained window till orders in 
> window are all filled or killed. <br />
> 3. Mark rest of unfilled orders to killed


Todos
---

- [x] Design 1st version of solution
- [x] Create test case for basic logic
- [x] Create main logic which details are in 2nd question of question section
- [x] Test scenarios
- [x] Unit test
- ~~[ ] Stubs~~
- [x] Create data model
- [x] Create an input file each line stand for an order
- [x] Build as an executable command


System Enhancement
---


Advance
---
- Containerize
- Deployment
- Github CI/CD
- Monitoring
- Load test