Exercise 3: Text Counter (URL Variables & Methods)
Goal: Build a server with a /count route. If a user sends a GET request, return the text "Send a POST request with text to count words". If they send a POST request, read the text body and return the number of characters.
Key Tasks:
Differentiate between GET and POST methods using r.Method.
Read the entire request body using io.ReadAll(r.Body).
Return the character length as a string.
