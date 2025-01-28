package weeklycompetition

import (
	"sort"
	"strconv"
	"strings"
)

func countMentions(numberOfUsers int, events [][]string) []int {
	mentions := make([]int, numberOfUsers)
	offlineEnd := make([]int, numberOfUsers)
	for i := range offlineEnd {
		offlineEnd[i] = -1 // Initially, all users are online
	}

	type eventWithIndex struct {
		timestamp     int
		eventType     int
		originalIndex int
		event         []string
	}

	eventList := make([]eventWithIndex, len(events))
	for i, e := range events {
		timestamp, _ := strconv.Atoi(e[1])
		eventType := 1
		if e[0] == "OFFLINE" {
			eventType = 0
		}
		eventList[i] = eventWithIndex{
			timestamp:     timestamp,
			eventType:     eventType,
			originalIndex: i,
			event:         e,
		}
	}

	sort.Slice(eventList, func(i, j int) bool {
		if eventList[i].timestamp != eventList[j].timestamp {
			return eventList[i].timestamp < eventList[j].timestamp
		}
		if eventList[i].eventType != eventList[j].eventType {
			return eventList[i].eventType < eventList[j].eventType
		}
		return eventList[i].originalIndex < eventList[j].originalIndex
	})

	for _, e := range eventList {
		if e.event[0] == "OFFLINE" {
			userID, _ := strconv.Atoi(e.event[2])
			timestamp, _ := strconv.Atoi(e.event[1])
			newEnd := timestamp + 60
			if newEnd > offlineEnd[userID] {
				offlineEnd[userID] = newEnd
			}
		} else {
			mentionsStr := e.event[2]
			timestamp, _ := strconv.Atoi(e.event[1])
			switch mentionsStr {
			case "ALL":
				for i := 0; i < numberOfUsers; i++ {
					mentions[i]++
				}
			case "HERE":
				for i := 0; i < numberOfUsers; i++ {
					if timestamp >= offlineEnd[i] {
						mentions[i]++
					}
				}
			default:
				parts := strings.Split(mentionsStr, " ")
				for _, part := range parts {
					if len(part) >= 2 && part[:2] == "id" {
						numStr := part[2:]
						if userID, err := strconv.Atoi(numStr); err == nil {
							if userID >= 0 && userID < numberOfUsers {
								mentions[userID]++
							}
						}
					}
				}
			}
		}
	}

	return mentions
}

func countMentions1(numberOfUsers int, events [][]string) []int {
	type Event struct {
		timestamp int
		eventType string
		data      string
	}

	eventList := make([]Event, len(events))
	for i, e := range events {
		timestamp, _ := strconv.Atoi(e[1])
		eventList[i] = Event{
			timestamp: timestamp,
			eventType: e[0],
			data:      e[2],
		}
	}

	sort.Slice(eventList, func(i, j int) bool {
		if eventList[i].timestamp != eventList[j].timestamp {
			return eventList[i].timestamp < eventList[j].timestamp
		}
		if eventList[i].eventType != eventList[j].eventType {
			return eventList[i].eventType == "OFFLINE"
		}
		return i < j
	})

	mentions := make([]int, numberOfUsers)
	offlineUntil := make([]int, numberOfUsers)

	for _, event := range eventList {
		timestamp := event.timestamp

		if event.eventType == "OFFLINE" {
			userID, _ := strconv.Atoi(event.data)
			offlineUntil[userID] = timestamp + 60
		} else if event.eventType == "MESSAGE" {
			tokens := strings.Split(event.data, " ")
			for _, token := range tokens {
				if token == "ALL" {
					for i := 0; i < numberOfUsers; i++ {
						mentions[i]++
					}
				} else if token == "HERE" {
					for i := 0; i < numberOfUsers; i++ {
						if offlineUntil[i] <= timestamp {
							mentions[i]++
						}
					}
				} else if strings.HasPrefix(token, "id") {
					userIDStr := token[2:]
					userID, err := strconv.Atoi(userIDStr)
					if err == nil && userID >= 0 && userID < numberOfUsers {
						mentions[userID]++
					}
				}
			}
		}
	}

	return mentions
}
