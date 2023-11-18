
curl  -k http://127.0.0.1:7001/v1/book/1449311601
Unauthorized

curl  -k http://127.0.0.1:7001/v1/book/1449311601 -u scott:tiger
Author: Ryan Boyd

curl  -k http://127.0.0.1:7001/v1/auth/token -u scott:tiger
token: b11d1e98-5d2e-4d56-89f8-415b68fdbd74

curl  -k http://127.0.0.1:7001/v1/book/1449311601 -H "Authorization: Bearer 11d1e98-5d2e-4d56-89f8-415b68fdbd74"
Author: Ryan Boyd