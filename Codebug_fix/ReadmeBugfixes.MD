Server Code Fixes

Changes Made

1. Imports
 Removed the 'slices' package: The original code used 'slices.Insert', which is not a part of the standard Go library. It seems to be an error or a placeholder for functionality to insert elements into a slice.
Replaced 'slices.Insert' with a custom implementation: The insertion of VIP riders at the start of the queue is now handled by creating a new slice and appending the VIP rider at the start.

2. Queue Handling
Added locks to ensure thread-safe access to 'rideQueue':
     Lock acquisition before modifying the queue: The 'rideQueueMu' mutex is locked before modifying the 'rideQueue' to ensure that concurrent   access does not corrupt the state of the queue.
     Queue size logging: The size of the queue is logged after adding a new rider.

3. Seating Riders
   Fixed seating logic to avoid concurrency issues:
        Properly locking 'rideQueue' during seating: Riders are moved from 'rideQueue' to ride in a thread-safe manner by locking 'rideQueueMu' and 'rideMu' appropriately.
        Proper car and seat assignment: The car and seat calculations were adjusted to ensure riders are seated correctly.

4. Ride Logic
   Fixed the start method to handle the ride correctly:
      Loop to process riders: The loop in the 'start' method now properly iterates over rideQueue and transfers riders to ride.
   	Clearing the queue after processing: The 'rideQueue' is truncated to remove riders who have been moved to the 'ride'.
    	Log messages for ride start and finish: Added 'logging' to indicate when the ride starts and finishes.
    	Clearing the ride array: After the ride duration, the 'ride' array is cleared to simulate the end of the ride.

By making these changes, the code now correctly handles the rollercoaster queue, seats riders properly, and ensures that the ride logic works without concurrency issues










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

