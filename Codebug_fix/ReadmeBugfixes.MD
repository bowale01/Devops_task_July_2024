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

1. Fixed the rand import:

      Changed math/rand/v2 to math/rand.


2. Seeding the random number generator:

      Added rand.Seed(time.Now().UnixNano()) inside the runEntrance function.

3. Corrected the range for rand.IntN:

     Changed rand.IntN(maxWaitTime-minWaitTime)+minWaitTime to rand.IntN(maxWaitTime-minWaitTime+1)+minWaitTime.

4. Added proper error handling for JSON marshalling:

      Added a check for errors when calling json.Marshal.
      Used log.Printf to log any errors encountered.

5. Added proper error handling for HTTP POST requests:

      Added a check for errors when calling http.Post.
      Used log.Printf to log any errors encountered.
      Closed the response body with resp.Body.Close().

6. Logging for shutdown:

      Added a log message to indicate when a shutdown signal is received.

These changes ensure that the program runs correctly, handles errors gracefully, and logs relevant information for debugging.

