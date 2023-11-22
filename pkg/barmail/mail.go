package barmail

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"

	"github.com/jinzhu/now"
)

const (
	mailBody = `Im folgenden der Barplan fuer naechste Woche. Wer Schichten uebernehmen
moechte, moege sich bitte einfach eintragen und den Plan als Antwort
wieder auf die Liste schicken. Die angegebenen Zeiten dienen der
Orientierung, ihr koennt sie euch gerne anpassen.

Wenn dann jeder nur auf die letzte Mail im Thread antwortet, sollte das
eigentlich ganz gut funktionieren...
%s
Dies ist eine automatisch gesendete Nachricht, sie ist
ohne Unterschrift gueltig.
`

	day = `
%s, %s:%s
15:00 -> 20:00:
19:00 -> 00:30:
`

	event = `
Event: %s, %s`
)

var days = map[int]string{
	0: "Mo",
	1: "Di",
	2: "Mi",
	3: "Do",
	4: "Fr",
	5: "Sa",
	6: "So",
}

type receivedEvents struct {
	Events   []Event `json:"c_base_events"`
	Regulars []Event `json:"c_base_regulars"`
}

type Event struct {
	Description string    `json:"description"`
	Title       string    `json:"title"`
	Start       time.Time `json:"start"`
}

func (re *receivedEvents) events() []Event {
	return append(re.Events, re.Regulars...)
}

func GetBarMail() error {
	resp, err := http.Get("http://c-base.org/calendar/exported/events.json")
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var cBaseEvents receivedEvents
	if err := json.Unmarshal(body, &cBaseEvents); err != nil {
		return err
	}

	then := time.Now().AddDate(0, 0, 7)
	now.WeekStartDay = time.Monday
	start := now.With(then).BeginningOfWeek()
	end := now.With(then).EndOfWeek()

	var events []Event
	for _, event := range cBaseEvents.events() {
		if event.Start.After(start) && event.Start.Before(end) {
			events = append(events, event)
		}
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].Start.Before(events[j].Start)
	})

	mapDays := make(map[int][]Event)
	for _, event := range events {
		weekday := int(event.Start.Weekday())
		if mapDays[weekday] == nil {
			mapDays[weekday] = []Event{}
		}
		mapDays[weekday] = append(mapDays[weekday], event)
	}

	var dayString string
	for i := 0; i < 7; i++ {
		if mapDays[i] == nil {
			dayString += fmt.Sprintf(day, days[i], start.AddDate(0, 0, i).Format("02.01.2006"), "")
			continue
		}
		var eventStrings string
		for _, eventData := range mapDays[i] {
			eventStrings += fmt.Sprintf(event, eventData.Title, eventData.Start.Format("15:04"))
		}
		dayString += fmt.Sprintf(day, days[i], start.AddDate(0, 0, i).Format("02.01.2006"), eventStrings)
	}

	fmt.Println(fmt.Sprintf(mailBody, dayString))
	return nil
}
