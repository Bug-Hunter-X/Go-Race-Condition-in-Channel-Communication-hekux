# Go Race Condition Example

This repository demonstrates a race condition in a Go program that uses channels for inter-goroutine communication. The program has multiple goroutines writing to a channel without proper synchronization, leading to potential data loss or unexpected program termination.

## How to Reproduce

1. Clone the repository.
2. Run the `bug.go` file.
3. Observe the output. The order of received messages will likely vary between executions, indicating that the race condition is present.

## Solution

The `bugSolution.go` file provides a corrected version of the program with a solution to prevent the race condition.  The solution includes proper synchronization mechanisms to ensure consistent and predictable behavior.