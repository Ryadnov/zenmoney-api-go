package entities

type (
	ObjectClass string

	Diff struct {
		CurrentClientTimestamp Timestamp `json:"currentClientTimestamp"`
		ServerTimestamp        Timestamp `json:"serverTimestamp"`

		ForceFetch []ObjectClass `json:"forceFetch"`

		Instrument     []Instrument     `json:"instrument"`
		Company        []Company        `json:"company"`
		User           []User           `json:"user"`
		Account        []Account        `json:"account"`
		Tag            []Tag            `json:"tag"`
		Merchant       []Merchant       `json:"merchant"`
		Budget         []Instrument     `json:"budget"`
		Reminder       []Reminder       `json:"reminder"`
		ReminderMarker []ReminderMarker `json:"reminderMarker"`
		Transaction    []Transaction    `json:"instrument"`

		//
		//	deletion: [
		//{
		//	id:     String -> Object.id
		//	object: String -> Object.class
		//	stamp:  Int
		//	user:   Int
		//}
		//]?
	}
)

const (
	//ObjectClassInstrument     ObjectClass = "Instrument"
	//ObjectClassCompany        ObjectClass = "Company"
	//ObjectClassUser           ObjectClass = "User"
	ObjectClassAccount        ObjectClass = "Account"
	ObjectClassTag            ObjectClass = "Tag"
	ObjectClassMerchant       ObjectClass = "Merchant"
	ObjectClassBudget         ObjectClass = "Budget"
	ObjectClassReminder       ObjectClass = "Reminder"
	ObjectClassReminderMarker ObjectClass = "ReminderMarker"
	ObjectClassTransaction    ObjectClass = "Transaction"
)
