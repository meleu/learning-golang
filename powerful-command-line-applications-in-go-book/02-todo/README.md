# ToDo CLI

Interesting things I've learned from this chapter:

- usual way to organize files and directories in a CLI project (p. 12)
- reading/writing JSON files with `json.Marshal()` and `json.Unmarshal()` (p. 16-17)
- to test only the public API of your package, just name your test package as `PKG_test` (p. 17)
- basic CLI integration tests (p. 24)
- read command line flags (p. 27-30)
- Stringer interface (p. 34)
- reading env vars (p.36-37)
- reading input from stdin (p. 39)

## Recreating

Steps I suggest (to myself) in order to recreate this project (using TDD, of course).

### TodoList API

- test with `package todo_test` (as we want to test only the public API)
- `TodoList` starts empty
  - implement it as a slice
- `TodoList.Add`
  - "add a task to the list"
  - starts with `Done` status as false
  - task starts with `CreatedAt` as `time.Now()` (I used a workaround to check this - beforeTime & afterTime)
- `TodoList.Complete`
  - mark the task as `Done`
  - `Complete` an non-existent task returns error
  - assign a timestamp to the `CompletedAt` task attribute
- `TodoList.Delete`
  - deletes a task from the list
- `TodoList.Save` & `TodoList.Get` (I think I would name `Load` instead of `Get`)
  - the book creates a test with two lists, one list1 saves a task list, and list2 gets the task list from  the file.

### todo CLI

- `TestMain` to build the binary, run the tests, and then clean up.
- add a task
  - via command line argument
  - via stdin
- list tasks
- complete a task
- delete a task
