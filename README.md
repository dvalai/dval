# dval

## Goal

The ultimate goal of dval is to provide a simple, yet powerful, way to validate any data. Once the goal is reached it should be possible to send any data to dval and get a response back if the data is valid or not. The response should also return which validator are applicable.

An example of what should be possible:

Input:
```
data="johndoe@gmail.com"
```

Response:
```
valid=true
type=email
```

As you can see, we're not asking dval to validate if the data is an email address, we will let dval figure it out. As such we should be able to send any data like telephone numbers, dates, credit card numbers, addresses, etc. and get a response back.

## Current state
