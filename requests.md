# Фактический прогон curl

```text
$ curl -sS -i --max-time 15 -X POST http://127.0.0.1:18010/api/v1/login -H Content-Type: application/json -d {"email": "admin@example.com", "password": "secret123"}
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 31 Oct 2026 08:44:18 GMT
Content-Length: 1008

{"access":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwicm9sZSI6ImFkbWluIiwia2luZCI6ImFjY2VzcyIsImlzcyI6InByMTAiLCJzdWIiOiIxIiwiYXVkIjpbInByMTAtY2xpZW50cyJdLCJleHAiOjE3OTM0MzcxNTgsImlhdCI6MTc5MzQzNjI1OCwianRpIjoiYTk0ZmMyYjZjNmI3NGQxNmE2NWM2NGVhNTBhNTNiMzAifQ.v8DKWVXHZPzA5zgUod5swskS8EpkAQkN3KRJdqyT8Sw","refresh":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwicm9sZSI6ImFkbWluIiwia2luZCI6InJlZnJlc2giLCJpc3MiOiJwcjEwIiwic3ViIjoiMSIsImF1ZCI6WyJwcjEwLWNsaWVudHMiXSwiZXhwIjoxNzk0MDQxMDU4LCJpYXQiOjE3OTM0MzYyNTgsImp0aSI6Ijg3OTczYWNiNzFiZTlhNGM5NDRiMDZmOWYwM2Y0N2M1In0.EqHX7mvceNbnrDFGMwIge9gJueqvpZhvabUcrUooBoM","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwicm9sZSI6ImFkbWluIiwia2luZCI6ImFjY2VzcyIsImlzcyI6InByMTAiLCJzdWIiOiIxIiwiYXVkIjpbInByMTAtY2xpZW50cyJdLCJleHAiOjE3OTM0MzcxNTgsImlhdCI6MTc5MzQzNjI1OCwianRpIjoiYTk0ZmMyYjZjNmI3NGQxNmE2NWM2NGVhNTBhNTNiMzAifQ.v8DKWVXHZPzA5zgUod5swskS8EpkAQkN3KRJdqyT8Sw"}


$ curl -sS -i --max-time 15 -X POST http://127.0.0.1:18010/api/v1/login -H Content-Type: application/json -d {"email": "user@example.com", "password": "secret123"}
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 31 Oct 2026 08:44:18 GMT
Content-Length: 999

{"access":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJyb2xlIjoidXNlciIsImtpbmQiOiJhY2Nlc3MiLCJpc3MiOiJwcjEwIiwic3ViIjoiMiIsImF1ZCI6WyJwcjEwLWNsaWVudHMiXSwiZXhwIjoxNzkzNDM3MTU4LCJpYXQiOjE3OTM0MzYyNTgsImp0aSI6Ijg1YTU4OTczOTIzYjI2NDBjYWE1ODc0ODg5ZDFlZjBkIn0.-JqGf881PLzip5Dmu7LufZGrU1AcyhXiyPQ1tsY4hiA","refresh":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJyb2xlIjoidXNlciIsImtpbmQiOiJyZWZyZXNoIiwiaXNzIjoicHIxMCIsInN1YiI6IjIiLCJhdWQiOlsicHIxMC1jbGllbnRzIl0sImV4cCI6MTc5NDA0MTA1OCwiaWF0IjoxNzkzNDM2MjU4LCJqdGkiOiI3NTlmZDM4MmRkZDIyNmE0ZjQ3ZmVlNzI2NjQ5NTIzZiJ9.keXBYrNgUo32wwk2e8W9rlBYvkmqKOKybTOPJbQ3fzM","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJyb2xlIjoidXNlciIsImtpbmQiOiJhY2Nlc3MiLCJpc3MiOiJwcjEwIiwic3ViIjoiMiIsImF1ZCI6WyJwcjEwLWNsaWVudHMiXSwiZXhwIjoxNzkzNDM3MTU4LCJpYXQiOjE3OTM0MzYyNTgsImp0aSI6Ijg1YTU4OTczOTIzYjI2NDBjYWE1ODc0ODg5ZDFlZjBkIn0.-JqGf881PLzip5Dmu7LufZGrU1AcyhXiyPQ1tsY4hiA"}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18010/api/v1/me -H Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwicm9sZSI6ImFkbWluIiwia2luZCI6ImFjY2VzcyIsImlzcyI6InByMTAiLCJzdWIiOiIxIiwiYXVkIjpbInByMTAtY2xpZW50cyJdLCJleHAiOjE3OTM0MzcxNTgsImlhdCI6MTc5MzQzNjI1OCwianRpIjoiYTk0ZmMyYjZjNmI3NGQxNmE2NWM2NGVhNTBhNTNiMzAifQ.v8DKWVXHZPzA5zgUod5swskS8EpkAQkN3KRJdqyT8Sw
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 31 Oct 2026 08:44:18 GMT
Content-Length: 54

{"email":"admin@example.com","id":"1","role":"admin"}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18010/api/v1/admin/stats -H Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwicm9sZSI6ImFkbWluIiwia2luZCI6ImFjY2VzcyIsImlzcyI6InByMTAiLCJzdWIiOiIxIiwiYXVkIjpbInByMTAtY2xpZW50cyJdLCJleHAiOjE3OTM0MzcxNTgsImlhdCI6MTc5MzQzNjI1OCwianRpIjoiYTk0ZmMyYjZjNmI3NGQxNmE2NWM2NGVhNTBhNTNiMzAifQ.v8DKWVXHZPzA5zgUod5swskS8EpkAQkN3KRJdqyT8Sw
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 31 Oct 2026 08:44:18 GMT
Content-Length: 12

{"users":2}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18010/api/v1/admin/stats -H Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJyb2xlIjoidXNlciIsImtpbmQiOiJhY2Nlc3MiLCJpc3MiOiJwcjEwIiwic3ViIjoiMiIsImF1ZCI6WyJwcjEwLWNsaWVudHMiXSwiZXhwIjoxNzkzNDM3MTU4LCJpYXQiOjE3OTM0MzYyNTgsImp0aSI6Ijg1YTU4OTczOTIzYjI2NDBjYWE1ODc0ODg5ZDFlZjBkIn0.-JqGf881PLzip5Dmu7LufZGrU1AcyhXiyPQ1tsY4hiA
HTTP/1.1 403 Forbidden
Content-Type: application/json
Date: Sat, 31 Oct 2026 08:44:18 GMT
Content-Length: 22

{"error":"forbidden"}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18010/api/v1/users/2 -H Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJyb2xlIjoidXNlciIsImtpbmQiOiJhY2Nlc3MiLCJpc3MiOiJwcjEwIiwic3ViIjoiMiIsImF1ZCI6WyJwcjEwLWNsaWVudHMiXSwiZXhwIjoxNzkzNDM3MTU4LCJpYXQiOjE3OTM0MzYyNTgsImp0aSI6Ijg1YTU4OTczOTIzYjI2NDBjYWE1ODc0ODg5ZDFlZjBkIn0.-JqGf881PLzip5Dmu7LufZGrU1AcyhXiyPQ1tsY4hiA
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 31 Oct 2026 08:44:18 GMT
Content-Length: 50

{"id":2,"email":"user@example.com","role":"user"}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18010/api/v1/users/1 -H Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJyb2xlIjoidXNlciIsImtpbmQiOiJhY2Nlc3MiLCJpc3MiOiJwcjEwIiwic3ViIjoiMiIsImF1ZCI6WyJwcjEwLWNsaWVudHMiXSwiZXhwIjoxNzkzNDM3MTU4LCJpYXQiOjE3OTM0MzYyNTgsImp0aSI6Ijg1YTU4OTczOTIzYjI2NDBjYWE1ODc0ODg5ZDFlZjBkIn0.-JqGf881PLzip5Dmu7LufZGrU1AcyhXiyPQ1tsY4hiA
HTTP/1.1 403 Forbidden
Content-Type: application/json
Date: Sat, 31 Oct 2026 08:44:18 GMT
Content-Length: 22

{"error":"forbidden"}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18010/api/v1/me
HTTP/1.1 401 Unauthorized
Content-Type: application/json
Date: Sat, 31 Oct 2026 08:44:19 GMT
Content-Length: 25

{"error":"unauthorized"}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18010/api/v1/me -H Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwicm9sZSI6ImFkbWluIiwia2luZCI6InJlZnJlc2giLCJpc3MiOiJwcjEwIiwic3ViIjoiMSIsImF1ZCI6WyJwcjEwLWNsaWVudHMiXSwiZXhwIjoxNzk0MDQxMDU4LCJpYXQiOjE3OTM0MzYyNTgsImp0aSI6Ijg3OTczYWNiNzFiZTlhNGM5NDRiMDZmOWYwM2Y0N2M1In0.EqHX7mvceNbnrDFGMwIge9gJueqvpZhvabUcrUooBoM
HTTP/1.1 401 Unauthorized
Content-Type: application/json
Date: Sat, 31 Oct 2026 08:44:19 GMT
Content-Length: 25

{"error":"unauthorized"}


$ curl -sS -i --max-time 15 -X POST http://127.0.0.1:18010/api/v1/refresh -H Content-Type: application/json -d {"refresh": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwicm9sZSI6ImFkbWluIiwia2luZCI6InJlZnJlc2giLCJpc3MiOiJwcjEwIiwic3ViIjoiMSIsImF1ZCI6WyJwcjEwLWNsaWVudHMiXSwiZXhwIjoxNzk0MDQxMDU4LCJpYXQiOjE3OTM0MzYyNTgsImp0aSI6Ijg3OTczYWNiNzFiZTlhNGM5NDRiMDZmOWYwM2Y0N2M1In0.EqHX7mvceNbnrDFGMwIge9gJueqvpZhvabUcrUooBoM"}
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 31 Oct 2026 08:44:19 GMT
Content-Length: 1008

{"access":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwicm9sZSI6ImFkbWluIiwia2luZCI6ImFjY2VzcyIsImlzcyI6InByMTAiLCJzdWIiOiIxIiwiYXVkIjpbInByMTAtY2xpZW50cyJdLCJleHAiOjE3OTM0MzcxNTksImlhdCI6MTc5MzQzNjI1OSwianRpIjoiYzQyNmI1NmExNzMyZmQxODM2MDIzOGY4NTgxZWI0ZjIifQ.3sEmFCVg7qXDcvdGF-3KxrN2uhStRK6aizk1rdNtb8w","refresh":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwicm9sZSI6ImFkbWluIiwia2luZCI6InJlZnJlc2giLCJpc3MiOiJwcjEwIiwic3ViIjoiMSIsImF1ZCI6WyJwcjEwLWNsaWVudHMiXSwiZXhwIjoxNzk0MDQxMDU5LCJpYXQiOjE3OTM0MzYyNTksImp0aSI6Ijk4OTU2MjcyZjlmMjYzODgyN2NlNDYyMmE5NzBjYzU4In0.jR96V74fa1rXTAkxW3S_XVdDU0207odVFrjT2XANJt8","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwicm9sZSI6ImFkbWluIiwia2luZCI6ImFjY2VzcyIsImlzcyI6InByMTAiLCJzdWIiOiIxIiwiYXVkIjpbInByMTAtY2xpZW50cyJdLCJleHAiOjE3OTM0MzcxNTksImlhdCI6MTc5MzQzNjI1OSwianRpIjoiYzQyNmI1NmExNzMyZmQxODM2MDIzOGY4NTgxZWI0ZjIifQ.3sEmFCVg7qXDcvdGF-3KxrN2uhStRK6aizk1rdNtb8w"}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18010/api/v1/me -H Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwicm9sZSI6ImFkbWluIiwia2luZCI6ImFjY2VzcyIsImlzcyI6InByMTAiLCJzdWIiOiIxIiwiYXVkIjpbInByMTAtY2xpZW50cyJdLCJleHAiOjE3OTM0MzcxNTksImlhdCI6MTc5MzQzNjI1OSwianRpIjoiYzQyNmI1NmExNzMyZmQxODM2MDIzOGY4NTgxZWI0ZjIifQ.3sEmFCVg7qXDcvdGF-3KxrN2uhStRK6aizk1rdNtb8w
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 31 Oct 2026 08:44:19 GMT
Content-Length: 54

{"email":"admin@example.com","id":"1","role":"admin"}


$ curl -sS -i --max-time 15 -X POST http://127.0.0.1:18010/api/v1/refresh -H Content-Type: application/json -d {"refresh": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwicm9sZSI6ImFkbWluIiwia2luZCI6InJlZnJlc2giLCJpc3MiOiJwcjEwIiwic3ViIjoiMSIsImF1ZCI6WyJwcjEwLWNsaWVudHMiXSwiZXhwIjoxNzk0MDQxMDU4LCJpYXQiOjE3OTM0MzYyNTgsImp0aSI6Ijg3OTczYWNiNzFiZTlhNGM5NDRiMDZmOWYwM2Y0N2M1In0.EqHX7mvceNbnrDFGMwIge9gJueqvpZhvabUcrUooBoM"}
HTTP/1.1 401 Unauthorized
Content-Type: application/json
Date: Sat, 31 Oct 2026 08:44:19 GMT
Content-Length: 28

{"error":"refresh_revoked"}


$ curl -sS -i --max-time 15 -X POST http://127.0.0.1:18010/api/v1/refresh -H Content-Type: application/json -d {"refresh": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwicm9sZSI6ImFkbWluIiwia2luZCI6ImFjY2VzcyIsImlzcyI6InByMTAiLCJzdWIiOiIxIiwiYXVkIjpbInByMTAtY2xpZW50cyJdLCJleHAiOjE3OTM0MzcxNTksImlhdCI6MTc5MzQzNjI1OSwianRpIjoiYzQyNmI1NmExNzMyZmQxODM2MDIzOGY4NTgxZWI0ZjIifQ.3sEmFCVg7qXDcvdGF-3KxrN2uhStRK6aizk1rdNtb8w"}
HTTP/1.1 401 Unauthorized
Content-Type: application/json
Date: Sat, 31 Oct 2026 08:44:19 GMT
Content-Length: 28

{"error":"invalid_refresh"}

```
