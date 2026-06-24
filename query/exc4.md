Exercise 4: Basic Math API (Multiple Query Parameters)
Goal: Create a /calculate route that accepts three query parameters: op (operation), a, and b. For example, /calculate?op=add&a=10&b=5 should respond with Result: 15.
Key Tasks:
Parse string query variables using r.URL.Query().Get().
Convert string inputs to integers using strconv.Atoi().
Support add, subtract, and multiply. Return an HTTP 400 Bad Request if the operation is unknown or parsing fails.
