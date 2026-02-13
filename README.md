### URL Shortener 

**Problem Statement:** create a services that shortens a long url to a simple short one 


### Non-Functional Requirement 

1. should support 100 million url generated request per day 
2. shorten url as short as possible 
3. Should be highly avaliable 

### Cost Estimation 
- 100,000,000 / (24/(60 * 60)) = 1160 writes per second 
- read can be in 10:1 ration so 1160 * 10 = 11600 reads per second 
- assuming url shortener will run for 5 years  = 100,000,000 * 365 * 10 = 365 billion records 
- average url length is  100 bytes 
- storage required will be 365 billion *  100 bytes = 36.5 TB 

