Exercise 5: User-Agent Echo (Reading Headers)
Goal: Create an /agent route that reads the incoming browser or client header details and echoes it back in plain text: "You are visiting us using: [User-Agent Info]".
Key Tasks:
Inspect request headers using r.Header.Get("User-Agent").
Handle instances where the header might be missing or empty.
