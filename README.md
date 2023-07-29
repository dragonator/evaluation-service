# evaluation-service
An interview assignment project: REST API providing endpoints for validating and evaluating expressions.

The service exposes the following endpoints for working with math expressions in verbal format:

`POST /evaluate` - evaluates a single math expression.

    {
        "expression": "<a math expression>"
    }

`POST /validate` - validates a single math expression.

    {
        "expression": "<a math expression>"
    }

`GET /errors` - returns errors occurred during evaluating or validating expressions.

The math expression should be in the following format: `What is <number> <operation> <number>?`

The supported operations are:
* `plus`
* `minus`
* `multiplied by`
* `divided by`

For a number one could provide both whole numbers and decimals. Examples:
* `What is 5.2 plus 13?`
* `What is 7 minus 5.5?`
* `What is 6.2 multiplied by 4?`
* `What is 25 divided by 5.5?`

The service can also handle a set of operations evaluating them from left to right. Here are some such examples and the result that they'll yield:
* `What is 3 plus 2 multiplied by 3?`: 15 (i.e. not 9)
* `What is 5 multiplied by 12 plus 3 divided by 3?`: 21 (i.e. not 61)


## Run the service

#### Init environment file:

    make init

#### Run the service:

    make server-start