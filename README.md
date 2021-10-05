# Auth mock service

Simple auth service for testing purposes.
Based on https://github.com/stefanprodan/podinfo

-------

Credentials:
```json
{
  "username": "username-role",
  "password": "any password"
}
```

Example:
```json
{
  "username": "john-admin",
  "password": "any password"
}
```

Response:
```json
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4iLCJyb2xlIjoiQURNSU4iLCJleHAiOjE2MzM0NjY1ODJ9.biI6wExoQJ9vIGOrlgrZ1GoZAk55TVJJs0gsIRLxBzA","username":"john","role":"ADMIN"}
```

Decoded JWT:
```json
{
  "alg": "HS256",
  "typ": "JWT"
}
{
  "username": "john",
  "role": "ADMIN",
  "exp": 1633466582
}
```