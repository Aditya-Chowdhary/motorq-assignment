## Questions
1. What all values shud we store in the vehicles table - up to us
## Notes



## Todo 
- [x] Function to call nhsta(vin)
- [x] Field for if its been set by child or parent?
- [ ] Created At + Updated At
- [ ] Bonus: Pagination for get orgs
- [ ] Version for block conflicts
- [ ] Bonus: Specific Organisation
- [x] Check if updating something that doesnt exist??
- [ ] Bonus: Add authentication to your APIs to ensure no one outside More Torque is using this.
- [ ] Bonus: Rate limit all API’s so that they aren’t being misused by a rogue script!
- [ ] Bonus: Implement caching in relevant places to save time and improve latency.
- [ ] In get request - check if the speed/fuel is inherited from parent
- [ ] Rate limit calling the NHSTA thing



## Endpoints
1. GET /vehicles/decode/:vin - call nhtsa
2. POST /vehicles
   1. vin - decode from nhtsa
   2. org
3. GET /vehicles/:vin - get from db
4. [x] POST /Orgs
   1. Name
   2. Accnt
   3. Website
   4. Fuel
   5. Speed
   6. Parent
5. Patch /orgs
   1. Accnt
   2. Website
   3. Fuel
   4. Speed
   5. Parent
6. [x] Get /orgs
