package _402

import (
	"fmt"
	"sort"
)

type room struct {
	roomNumber int
	endTime    int
}

func insertSortedRoom(rooms []room, e room) []room {
	i := sort.Search(len(rooms), func(i int) bool { return rooms[i].endTime > e.endTime })
	rooms = append(rooms, room{})
	copy(rooms[i+1:], rooms[i:])
	rooms[i] = e
	return rooms
}

func getEnded(rooms []room, endTime int) ([]room, []room) {
	var roomsWithEndedMeeting int
	for _, r := range rooms {
		if r.endTime > endTime {
			break
		}
		roomsWithEndedMeeting++
	}

	freeRooms := make([]room, roomsWithEndedMeeting)
	copy(freeRooms, rooms[:roomsWithEndedMeeting])

	copy(rooms[:len(rooms)-roomsWithEndedMeeting], rooms[roomsWithEndedMeeting:])
	return rooms[:len(rooms)-roomsWithEndedMeeting], freeRooms
}

func popRoom(s []room) ([]room, room) {
	v := s[0]
	copy(s[:len(s)-1], s[1:])
	return s[:len(s)-1], v
}

func insertSorted(s []int, e int) []int {
	i := sort.Search(len(s), func(i int) bool { return s[i] > e })
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = e
	return s
}

func pop(s []int) ([]int, int) {
	if len(s) <= 0 {
		return s, -1
	}
	v := s[0]
	copy(s[:len(s)-1], s[1:])
	return s[:len(s)-1], v
}

type meetings [][]int

func (m meetings) Len() int {
	return len(m)
}

func (m meetings) Less(i, j int) bool {
	return m[i][0] < m[j][0]
}

func (m meetings) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func mostBooked(n int, m [][]int) int {
	sort.Sort(meetings(m))
	fmt.Println(m)

	roomUsage := make([]int, n)
	freeRooms := make([]int, n)
	for i := 0; i < n; i++ {
		freeRooms[i] = i
	}
	bookedRooms := make([]room, 0, n)

	var roomNumber int
	var endedMeetings []room
	var freeRoom room

	for _, meeting := range m {
		bookedRooms, endedMeetings = getEnded(bookedRooms, meeting[0])
		for _, r := range endedMeetings {
			freeRooms = insertSorted(freeRooms, r.roomNumber)
		}

		freeRooms, roomNumber = pop(freeRooms)
		if roomNumber == -1 {
			bookedRooms, freeRoom = popRoom(bookedRooms)
			var newEndTime int
			if meeting[0] < freeRoom.endTime {
				newEndTime = meeting[1] - meeting[0] + freeRoom.endTime
			} else {
				newEndTime = meeting[1]
			}
			bookedRooms = insertSortedRoom(bookedRooms, room{roomNumber: freeRoom.roomNumber, endTime: newEndTime})
			roomUsage[freeRoom.roomNumber]++
			fmt.Println(bookedRooms)
			fmt.Println(freeRooms)
			fmt.Println()

			continue
		}
		bookedRooms = insertSortedRoom(bookedRooms, room{roomNumber: roomNumber, endTime: meeting[1]})
		roomUsage[roomNumber]++

		fmt.Println(bookedRooms)
		fmt.Println(freeRooms)
		fmt.Println()
	}

	var mostBookedRoomNumber int
	var mostBookedNumber int
	for rn, bn := range roomUsage {
		if mostBookedNumber < bn {
			mostBookedRoomNumber = rn
			mostBookedNumber = bn
		}
	}

	return mostBookedRoomNumber
}
