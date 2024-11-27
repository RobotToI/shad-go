//go:build !solution

package hotelbusiness

type Guest struct {
	CheckInDate  int
	CheckOutDate int // 0 means today is out date
}

type Load struct {
	StartDate  int
	GuestCount int
}

/*
Need to know when all guests goes out, and new one gets in
only when number of guests is changed

	guests: []Guest{{1, 2}},
	result: []Load{{1, 1}, {2, 0}},
*/
func ComputeLoad(guests []Guest) []Load {
	hotelLog := make(map[int]int)
	maxDays := 0
	for _, guest := range guests {
		for i := guest.CheckInDate; i <= guest.CheckOutDate; i++ {
			hotelLog[i] += 1
			if i > maxDays {
				maxDays = i
			}
		}
		hotelLog[guest.CheckOutDate] -= 1
	}

	resultLoad := []Load{}
	prevGuestsNumber := 0
	for day := 0; day <= maxDays; day++ {
		if hotelLog[day] != prevGuestsNumber {
			resultLoad = append(resultLoad, Load{StartDate: day, GuestCount: hotelLog[day]})
		}
		prevGuestsNumber = hotelLog[day]
	}
	return resultLoad
}
