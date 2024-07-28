Server Code Fixes

  1. Proper Mutex Usage: I Ensured proper locking and unlocking mechanisms to prevent data races.

     I Added proper use of 'rideQueueMu' and 'rideMu' to avoid race conditions when accessing 'rideQueue' and 'ride'.

 2. Check for Ride Capacity: I Ensured that the ride queue does not exceed the capacity of the ride.

     Added logic to ensure that only the required number of riders are moved from the queue to the ride.
    
 3. Decoding JSON Request: I Ensured the JSON request body is properly decoded and the body is closed after reading.

       Added 'defer r.Body.Close()' to close the request body after reading.


 4. Append Riders Properly: Used slices.Insert and append correctly to maintain the queue.

     Used 'append([]*rider{rider}, rc.rideQueue...)' for inserting VIP riders at the front.






Client Code Fixes

1. Proper Error Handling: Added error handling for HTTP requests.

     I Added checks to handle errors returned by 'http.Post'.

2. Proper Random Number Generation: Ensured 'rand' is seeded properly for random sleep durations.

    I Used 'rand.Seed(time.Now().UnixNano())' to seed the random number generator.


