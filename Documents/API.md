### User API
1. **Register User**
    - Endpoint: `/api/users/register`
    - Method: POST
    - Body: `{ "accountNumber": string, "password": string }`
    - Response: `{ "success": boolean, "message": string, "userId": string }`

2. **Login User**
    - Endpoint: `/api/users/login`
    - Method: POST
    - Body: `{ "accountNumber": string, "password": string }`
    - Response: `{ "success": boolean, "message": string, "token": string }`

3. **Get User Information**
    - Endpoint: `/api/users/{userId}`
    - Method: GET
    - Headers: `{ "Authorization": "Bearer <token>" }`
    - Response: `{ "accountNumber": string, "torTorCoinBalance": float }`

### Transaction API
1. **Create Transaction**
    - Endpoint: `/api/transactions`
    - Method: POST
    - Headers: `{ "Authorization": "Bearer <token>" }`
    - Body: `{ "receiverAccountNumber": string, "amount": float }`
    - Response: `{ "success": boolean, "message": string, "transactionId": string }`

2. **Get Transaction History**
    - Endpoint: `/api/transactions/history/{userId}`
    - Method: GET
    - Headers: `{ "Authorization": "Bearer <token>" }`
    - Response: `[ { "sender": string, "receiver": string, "amount": float, "timestamp": datetime } ]`

### Friendship API
1. **Add Friend**
    - Endpoint: `/api/friends/add`
    - Method: POST
    - Headers: `{ "Authorization": "Bearer <token>" }`
    - Body: `{ "friendAccountNumber": string }`
    - Response: `{ "success": boolean, "message": string }`

2. **Get Friends List**
    - Endpoint: `/api/friends/list/{userId}`
    - Method: GET
    - Headers: `{ "Authorization": "Bearer <token>" }`
    - Response: `[ { "accountNumber": string, "friendSince": datetime } ]`

### Request API
1. **Create Assistance Request**
    - Endpoint: `/api/requests`
    - Method: POST
    - Headers: `{ "Authorization": "Bearer <token>" }`
    - Body: `{ "friendAccountNumber": string, "description": string, "torTorCoinCost": float }`
    - Response: `{ "success": boolean, "message": string, "requestId": string }`

2. **Complete Assistance Request**
    - Endpoint: `/api/requests/complete/{requestId}`
    - Method: PUT
    - Headers: `{ "Authorization": "Bearer <token>" }`
    - Response: `{ "success": boolean, "message": string }`