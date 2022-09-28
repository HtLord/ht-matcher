Ht-Matcher
===

Goal
---

Simulating stock exchange dealing orders match up behavior which sold or bough with the 
same price. Display how deal been matched with each other.

### E.g.

### Inputs(order type will not be implemented)

| timestamp | seq | action | quantity | orderType  | price |
|---|---|--------|----------|------------|-------|
|2022-01-01T00:00:00Z| 1| buy    | 10       | same price | 5.1   | 
|2022-01-01T00:01:01Z| 2| buy    | 10       | same price | 5.2   | 
|2022-01-01T00:02:02Z| 3| buy    | 50       | same price | 5.3   | 
|2022-01-01T00:03:03Z| 4| sell   | 10       | same price | 5.1   | 
|2022-01-01T00:04:04Z| 5| buy    | 10       | same price | 5.2   | 
|2022-01-01T00:05:05Z| 6| sell   | 70       | same price | 5.3   | 


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
| sell seq | buy seq | quantity | price |
|---|----|---|---|
| 4 | 1 | 10 | 5.1|


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
- [ ] Create test case for basic logic
- [ ] Create main logic which details are in 2nd question of question section
- [ ] Test scenarios
- [ ] Unit test 
- [ ] Stubs
- [ ] Create data model
- [ ] Create an input file each line stand for an order
- [ ] Build as an executable command


System Enhancement
---


Advance
---
- Containerize
- Deployment
- Github CI/CD
- Monitoring
- Load test