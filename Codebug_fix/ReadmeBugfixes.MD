Server Code Fixes
Let's break down the improvements and bug fixes made to the original code based on the updated version provided:

1. Fixed Initialization of ride Slice
Original Code:

rc := rollercoaster{
	ride: make([]*rider, numberOfCars*carCapacity),
}

Updated Code:

rc := rollercoaster{
	ride: make([]*rider, 0, numberOfCars*carCapacity),
}


Fixes:

The ride slice is initialized with zero capacity in the updated code. In the original code, it was initialized with a capacity equal to the maximum number of riders. This was likely intended to be a dynamic list that starts empty, so initializing it with 0 capacity is more appropriate.






2. Changed json.NewDecoder(r.Body).Decode(d) to Handle Errors
Original Code:

json.NewDecoder(r.Body).Decode(d)

Updated Code:
if err := json.NewDecoder(r.Body).Decode(d); err != nil {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}


Fixes:

The updated code adds error handling to the JSON decoding process. This ensures that if decoding fails, a 400 Bad Request response is returned to the client, improving robustness.


3. Revised rideQueue Management for VIP and Non-VIP Riders
Original Code:

if rider.vipStatus {
	rc.rideQueue = slices.Insert(rc.rideQueue, 0, rider)
} else {
	rc.rideQueue = append(rc.rideQueue, rider)
}


Updated Code:

if newRider.vipStatus {
	rc.rideQueue = append([]*rider{newRider}, rc.rideQueue...)
} else {
	rc.rideQueue = append(rc.rideQueue, newRider)
}


Fixes:

Changed from using slices.Insert to using append with a slice literal for VIP riders. This is more idiomatic and straightforward for prepending an element to a slice.


4. Improved rideQueue Processing in start Method
Original Code:


i := 0
for ; i <= len(rc.rideQueue)-1; i++ {
	if i == carCapacity*numberOfCars {
		break
	}
	rc.rideQueueMu.Lock()
	r := rc.rideQueue[0]
	car := int(math.Abs(float64((i)/2))) + 1
	carSeat := i % carCapacity
	rc.seatRider(r, car, carSeat)
	rc.rideQueue = rc.rideQueue[1:]
	rc.rideQueueMu.Unlock()
}

Updated Code:


if len(rc.rideQueue) > 0 {
	for i := 0; i < numberOfCars*carCapacity && i < len(rc.rideQueue); i++ {
		r := rc.rideQueue[i]
		rc.ride = append(rc.ride, r)
		car := i / carCapacity
		carSeat := i % carCapacity
		rc.rideQueueMu.Unlock()

		rc.rideMu.Lock()
		rc.seatRider(r, car, carSeat)
		rc.rideMu.Unlock()

		rc.rideQueueMu.Lock()
	}
	rc.rideQueue = rc.rideQueue[min(numberOfCars*carCapacity, len(rc.rideQueue)):]
}



Fixes:

The updated code ensures proper locking and unlocking of rideQueueMu and rideMu.
Correctly calculates the car number using integer division (i / carCapacity), whereas the original code used math.Abs, which was incorrect.
Ensures that the ride queue is trimmed based on the number of cars and capacity, improving the management of ride queue processing.


5. Adjusted Logging Messages for Seat Numbering
Original Code:

log.Printf("Ride: %s entering the car %d in seat %d\n", r.name, car, carSeat)

Updated Code:

log.Printf("Ride: %s entering the car %d in seat %d\n", r.name, car+1, carSeat+1)


Fixes:

Adjusted car and seat numbering to start from 1 for better readability. The original code used zero-based indexing which was less intuitive.


6. Added min Function
Original Code:
No min function was defined in the original code.
Updated Code:


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


Fixes:

Added a min function to handle the logic of limiting the length of the rideQueue. This function is essential for ensuring that the slice operations do not exceed its capacity.










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

