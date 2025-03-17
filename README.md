# pickmeshki-lab-2
---

## Description

This project is Lab #2 for the Software Architecture course. The main goal is to implement Continuous Integration (CI) practices and reinforce the use of dependency injection to facilitate testing. The lab involves working with prefix and postfix mathematical expressions and their conversion.

## Continuous Integration Aspects

CI is achieved through three main aspects:

-Version control of the code (using Git)

-Build automation

-Test automation

In this lab, you will implement and test functions for evaluating or converting prefix/postfix expressions while setting up CI/CD pipelines using GitHub Actions.

## Installation

Clone the repository:
```
git clone https://github.com/yur-ochka/pickmeshki-lab-2.git
```

Navigate to the project directory:
```
cd pickmeshki-lab-2

```

## Usage

Running the Program

You can run the application using the following commands:
```
go run ./cmd/example -e "5 5 +"

```
```
go run ./cmd/example -f input.txt

```
Use the -o flag to specify an output file:

go run ./cmd/example -e "42 1 -" -o result.txt

## Testing

### Run tests using:
```
go test
```

Ensure that your implementation supports both simple and complex expressions, along with error handling for invalid input.

## CI/CD Pipeline

This project includes GitHub Actions for automated testing and building. The CI pipeline runs tests on each commit and marks the build as failed if any tests fail.

### Example CI Build Links

[Successful Build Example](https://github.com/yur-ochka/pickmeshki-lab-2/pull/1)

[Failed Build Example](https://github.com/yur-ochka/pickmeshki-lab-2/pull/2)
