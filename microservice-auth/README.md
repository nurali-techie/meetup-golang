## Run - all four services (api-gateway, auth-service, book-service, review-service)

```
cd $GOPATH/src/github.com/nurali-techie/meetup-golang/microservice-auth/api-gateway/
go run api_gateway.go
INFO[0000] api gateway running at:localhost:8080        

cd $GOPATH/src/github.com/nurali-techie/meetup-golang/microservice-auth/auth-service
go run auth_srvc.go
INFO[0000] auth service running at:localhost:8081       

cd $GOPATH/src/github.com/nurali-techie/meetup-golang/microservice-auth/book-service
go run book_srvc.go
INFO[0000] book service running at:localhost:8082       

cd $GOPATH/src/github.com/nurali-techie/meetup-golang/microservice-auth/review-service
go run review_sevc.go
INFO[0000] review service running at:localhost:8083     
```

## Call - different API (Login, Get Book, Get Book Reviews)

**Login call:**
```
curl -X POST http://localhost:8080/api/login --data '{"username":"ali", "password":"abcd1234"}'
```
output:
```
"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsaUBnbWFpbC5jb20iLCJleHAiOjE1NTU0MzU2NzYsInJvbGUiOiJjdXN0b21lciIsInN1YiI6IjEiLCJ0eXAiOiJCZWFyZXIifQ.hXTu9VWMySF0774sW1rrA9_zvBLlPHlmc24szQjMmp6sE7tvOnVcG-z5TXUylgsbMQ03c2BGJalI2Ax88pqtPEJH5QmewonKMfCfCeZ5vfeYXwPuqt_sZWBDL7MNT4ysP341_0JoRcZjH4Mqrz05TxiYuGBxrcJpGjF80bZ0jeDx5P_bo5B1YVmIpdi7U039hinX77ZVOkgxLIaxSHwAmF6loMO8AMDicLB4fK0bZ46ANCqfrdfqKggZlUm8FN9Ppg5BgwGYsoZLIFVVQbYPysxYM_hddoYr85ZIf3WQ4TE_OqKJDedfjsPzCvh03QABFr_ebATjOYDh9GDO_owdWQ","token_type":"Bearer","expiry":"2019-04-16T22:57:56+05:30"}
```

export token to variable:
```
export token=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsaUBnbWFpbC5jb20iLCJleHAiOjE1NTU0MzU2NzYsInJvbGUiOiJjdXN0b21lciIsInN1YiI6IjEiLCJ0eXAiOiJCZWFyZXIifQ.hXTu9VWMySF0774sW1rrA9_zvBLlPHlmc24szQjMmp6sE7tvOnVcG-z5TXUylgsbMQ03c2BGJalI2Ax88pqtPEJH5QmewonKMfCfCeZ5vfeYXwPuqt_sZWBDL7MNT4ysP341_0JoRcZjH4Mqrz05TxiYuGBxrcJpGjF80bZ0jeDx5P_bo5B1YVmIpdi7U039hinX77ZVOkgxLIaxSHwAmF6loMO8AMDicLB4fK0bZ46ANCqfrdfqKggZlUm8FN9Ppg5BgwGYsoZLIFVVQbYPysxYM_hddoYr85ZIf3WQ4TE_OqKJDedfjsPzCvh03QABFr_ebATjOYDh9GDO_owdWQ
```

**Get book details call:**
```
curl http://localhost:8080/api/books?id=1 -H "Authorization: Bearer $token"
```
output:
```
{"ID":"1","Name":"Microservices Patterns","Author":"Chris Richardson","Price":50,"Reviews":[{"Reviewer":"Hemal","Rating":5,"Comment":"Great book with real life example"},{"Reviewer":"Mani","Rating":5,"Comment":"Nice content and explaination"}]}
```

**Get book reviews call:**
```
curl http://localhost:8080/api/reviews?bookId=1 -H "Authorization: Bearer $token"
```
output:
```
[{"Reviewer":"Hemal","Rating":5,"Comment":"Great book with real life example"},{"Reviewer":"Mani","Rating":5,"Comment":"Nice content and explaination"},{"Reviewer":"John","Rating":4,"Comment":"Good book but some repetation"},{"Reviewer":"Lacy","Rating":5,"Comment":"Love this book"}]
```

## Slides - Microservice Authentication with Golang by Nurali Virani:

1. Basic design with Book and Review service:

![microservice_basic_design](https://raw.githubusercontent.com/nurali-techie/meetup-golang/blob/master/microservice-auth/slides/slide01_ms_basic_design.png)

2. Auth service added for authentication:

![microservice_auth_service](https://raw.githubusercontent.com/nurali-techie/meetup-golang/blob/master/microservice-auth/slides/slide02_ms_with_auth_service.png)

3. Auth service with JWT token:

![auth_service_jwt_token](https://raw.githubusercontent.com/nurali-techie/meetup-golang/blob/master/microservice-auth/slides/slide03_auth_service_with_jwt.png)

4. Final design with API gateway:

![final_design_api_gateway](https://raw.githubusercontent.com/nurali-techie/meetup-golang/blob/master/microservice-auth/slides/slide04_final_ms_design_with_api_gateway.png)
