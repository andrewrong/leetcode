# 🚀 LeetCode Solutions in Go

This repository contains my personal solutions to LeetCode problems, implemented in Go.
The project is structured to be modular, testable, and easy to navigate.

## Project Structure

- `kit/`: Contains common data structures like `ListNode` and `TreeNode` to avoid code duplication.
- `problems/`: Each subdirectory corresponds to a single LeetCode problem, named with its ID and title (e.g., `lc0001_two_sum`).
  - Each problem directory has a `solution.go` for the implementation and a `solution_test.go` for unit tests.

## How to Use

### Add a New Problem

1. Copy the template directory:
   ```bash
   cp -r problems/lcxxxx_template/ problems/lc0001_two_sum
   ```
2. Write your solution in `solution.go`.
3. Add test cases to `solution_test.go`.

### Run Tests

- **Run all tests:**
  ```bash
  go test ./...
  ```
- **Test a specific problem:**
  ```bash
  go test -v ./problems/lc0001_two_sum
  ```

## My Progress

*(Optional: You can create a table here to track your progress)*

| #    | Title                | Solution                                       | Difficulty |
|------|----------------------|------------------------------------------------|------------|
| 1    | Two Sum              | [Go](./problems/lc0001_two_sum/solution.go)    | Easy       |

