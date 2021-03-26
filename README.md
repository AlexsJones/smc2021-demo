## Traffic splitting demo 😅

This demonstrates Linkerd's traffic splitting with real code.

- This code deploys 2 Apps with 1 client
- `go-open-api-v2` will start rejecting the client creation calls as it has a change in it's spec.


![images](images/splitting.png)


### Setup 🙋


- `kind create cluster --image rossgeorgiev/kind-node-arm64:v1.20.0`
- `linkerd install | kubectl apply -f -`
- `linkerd viz install | kubectl apply -f -`
- `kubectl create ns apps`
- `kubectl annotate ns apps "linkerd.io/inject=enabled"`

```
cd demo
./deploy_with_helm.sh
```

```
❯ k get pods -n apps
NAME                                 READY   STATUS    RESTARTS   AGE
go-openapi-client-668b94dbf4-cwbsp   2/2     Running   0          35s
go-openapi-v1-549c9cb676-spzsb       2/2     Running   0          35s
go-openapi-v2-7bb874cf5b-c7mg2       2/2     Running   0          35s
```


Apply the traffic split `kubectl apply -f demo/traffic-split.yaml` 💅

```
❯ linkerd -n apps viz routes deploy/go-openapi-v2
ROUTE             SERVICE   SUCCESS      RPS   LATENCY_P50   LATENCY_P95   LATENCY_P99
[DEFAULT]   go-openapi-v1   100.00%   0.0rps           1ms           9ms          10ms
[DEFAULT]   go-openapi-v2         -        -             -             -             -`
```


#### Code differences with go-openapi-v2

The V2 API decided all users need to state their food preferences for the international buffet. 🥝🫑🌶🥒🥬

```
#go-openapi-v2
definitions:
  User:
    type: "object"
    required:
      - foodPreference
    properties:
      id:
        type: "integer"
        format: "int64"
      username:
        type: "string"
      firstName:
        type: "string"
      lastName:
        type: "string"
      foodPreference:
        type: "string"
```

Let's post to test that new API directly...

```
curl -X POST "http://localhost:8080/v2/user" -H  "accept: application/xml" -H  "Content-Type: application/json" -d "{  \"email\": \"string\",  \"firstName\": \"string\",  \"id\": 0,  \"lastName\": \"string\",  \"password\": \"string\",  \"phone\": \"string\",  \"userStatus\": 0,  \"username\": \"alex\"}"
{"code":602,"message":"foodPreference in body is required"}%
```

Bummer! 🙅🏽

See that a bunch of stuff is erroring out due to missing field on the User object yay, it works! 👍


```
│ go-openapi-client &{madisonwhite662@example.org Matthew Harris <nil> 0 Wilson Cougarfir +243 56 11472 4389 0 Spiritshell}                                                           │
│ go-openapi-client Error:  [POST /user][422] createUser default                                                                                                                      │
│ go-openapi-client &{charlottejones332@example.net Addison Robinson <nil> 0 Thomas Painterrust +371 750 5343 155 2 0 Fishersleet}                                                    │
│ go-openapi-client Error:  [POST /user][422] createUser default                                                                                                                      │
│ go-openapi-client &{aidenwilson270@example.net Olivia White <nil> 0 Thompson Whipwheat +241 665 6 0395302 0 Thiefdark}                                                              │
│ go-openapi-client &{williamgarcia341@example.com Madison Miller <nil> 0 Davis Touchplanet +972 037 3996359 5 0 Tradergravel}                                                        │
│ go-openapi-client Error:  [POST /user][422] createUser default                                                                                                                      │
│ go-openapi-client &{madisongarcia667@example.org William Williams <nil> 0 Garcia Binderplume +880 304 77801444 0 Foxrowan}                                                          │
│ go-openapi-client Error:  [POST /user][422] createUser default                                                                                                                      │
│ go-openapi-client &{madisonthompson427@example.org Olivia Wilson <nil> 0 Williams Salmonoasis +252 4 1543558433 0 Ladysplit}                                                        │
│ go-openapi-client Error:  [POST /user][422] createUser default                                                                                                                      │
│ go-openapi-client &{joshuamiller826@example.com Matthew Miller <nil> 0 White Volescratch +506 201 753 611 52 0 Whipchrome}                                                          │
│ go-openapi-client Error:  [POST /user][422] createUser default                                                                                                                      │
│ go-openapi-client &{noahmartin562@example.org James Anderson <nil> 0 Taylor Butterflysugar +45 33 817 26608 55 0 Cowlbrick}                                                         │
│ go-openapi-client &{addisonbrown520@test.org Avery White <nil> 0 Thomas Napeshadow +33 73 9182632411 0 Servantharvest}                                                              │
│ go-openapi-client &{aidenwhite622@example.org Natalie Jones <nil> 0 Robinson Pixieglory +993 85 23550 92 3 2 0 Vulturetulip}                                                        │
│ go-openapi-client &{sofiajones246@example.net Benjamin Thompson <nil> 0 White Slicerweak +963 717 1 6 124055 0 Shirttranslucent}                                                    │
│ go-openapi-client &{zoeywilson061@example.com Mia Williams <nil> 0 Miller Whaleolive +218 53 81547153 6 0 Princessbog}
```
