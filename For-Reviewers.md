# Notes

## Todo

This project is still missing many pieces, mostly beyond the scope of the assignment:
- Additional tests
- Pagination
- Risk updates and deletion
- Persistent storage


## Implementation notes 

I've made assumptions around the POST endpoint. Mainly, that this is strictly to add a new Risk. This also assumed that the risk ID and status are not user-assigned at this point. IDs are generated at creation, and initial status is assumed to be `open`. 

Creating a new risk requires POSTing a Risk fragment (title and description only):

```
curl \
    --header "Content-Type: application/json" \
    --request POST \
    --data '{"title":"risk title","description":"risk body"}' \
    http://localhost:8080/v1/risks
```

A postman config is available `postman.json` for the basic requests. 



