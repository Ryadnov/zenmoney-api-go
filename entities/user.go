package entities

import (
	"fmt"

	"github.com/satori/go.uuid"
)

type (
	AccountId ObjectId
	MerchantId ObjectId
	TagId ObjectId
	ReminderId ObjectId
	ReminderMarkerId ObjectId
	TransactionId ObjectId

//
//	ObjectClassAccount        ObjectClass = "Account"
//ObjectClassTag            ObjectClass = "Tag"
//ObjectClassMerchant       ObjectClass = "Merchant"
//ObjectClassBudget         ObjectClass = "Budget"
//ObjectClassReminder       ObjectClass = "Reminder"
//ObjectClassReminderMarker ObjectClass = "ReminderMarker"
//ObjectClassTransaction    ObjectClass = "Transaction"

	// format 'yyyy-MM-dd'
	Date string
	// >= -90  && <= 90
	Latitude float64
	// >= -180 && <= 180
	Longitude float64
	// >= 0
	Value float64

	AccountType string
	ReminderMarkerState string

	DateInterval string
	AccountEndDateOffset DateInterval
	AccountPayoffInterval DateInterval
	ReminderInterval DateInterval

	UserEntity interface {
		TouchChanged()
	}

	Account struct {
		Id         AccountId
		Changed    Timestamp
		User       UserId
		Role       *UserId
		Instrument *InstrumentId
		Company    *CompanyId
		Type       AccountType
		Title      string
		SyncID     []string

		Balance      *float64
		StartBalance *float64
		CreditLimit  *float64

		InBalance        bool
		Savings          *bool
		EnableCorrection bool
		EnableSMS        bool
		Archive          bool

		// Для счетов с типом отличных от 'loan' и 'deposit' в  этих полях можно ставить null
		Capitalization        *bool
		Percent               *float64
		StartDate             *Date
		EndDateOffset         *int
		EndDateOffsetInterval *AccountEndDateOffset
		PayoffStep            *int
		PayoffInterval        *AccountPayoffInterval
	}

	Tag struct {
		Id      TagId
		Changed Timestamp
		User    UserId

		Title   string
		Parent  *TagId
		Icon    *string
		Picture *string
		Color   *int

		ShowIncome    bool
		ShowOutcome   bool
		BudgetIncome  bool
		BudgetOutcome bool
		Required      *bool
	}

	Merchant struct {
		Id      MerchantId
		Changed Timestamp
		User    UserId
		Title   string
	}

	Reminder struct {
		Id      ReminderId
		Changed Timestamp
		User    UserId

		IncomeInstrument  InstrumentId
		IncomeAccount     AccountId
		Income            Value
		OutcomeInstrument InstrumentId
		OutcomeAccount    AccountId
		Outcome           Value

		Tag      *[]TagId
		Merchant *MerchantId
		Payee    *string
		Comment  *string

		Interval  *ReminderInterval
		Step      *int
		Points    []int
		StartDate Date
		EndDate   *Date
		Notify    bool
	}

	ReminderMarker struct {
		Id      ReminderMarkerId
		Changed Timestamp
		User    UserId

		IncomeInstrument  InstrumentId
		IncomeAccount     AccountId
		Income            Value
		OutcomeInstrument InstrumentId
		OutcomeAccount    AccountId
		Outcome           Value

		Tag      *[]TagId
		Merchant *MerchantId
		Payee    *string
		Comment  *string

		Date Date

		Reminder ReminderId
		State    ReminderMarkerState

		Notify bool
	}

	Transaction struct {
		Id      TransactionId
		Changed Timestamp
		Created Timestamp
		User    UserId
		Deleted bool
		Hold    *bool

		IncomeInstrument  InstrumentId
		IncomeAccount     AccountId
		Income            Value
		OutcomeInstrument InstrumentId
		OutcomeAccount    AccountId
		Outcome           Value

		Tag           *[]TagId
		Merchant      *MerchantId
		Payee         *string
		OriginalPayee *string
		Comment       *string

		Date Date

		Mcc *int

		QrCode *string //  qrCode

		ReminderMarker *ReminderMarkerId

		OpIncome            *Value
		OpIncomeInstrument  *InstrumentId
		OpOutcome           *Value
		OpOutcomeInstrument *InstrumentId

		Latitude  *Latitude
		Longitude *Longitude
	}

	Budget struct {
		Id      ReminderMarkerId
		Changed Timestamp
		User    UserId

		Tag *TagId

		Date Date

		Income      Value
		IncomeLock  bool
		Outcome     Value
		OutcomeLock bool
	}
)

const (
	ReminderMarkerStatePlanned   ReminderMarkerState = "planned"
	ReminderMarkerStateProcessed ReminderMarkerState = "processed"
	ReminderMarkerStateDeleted   ReminderMarkerState = "deleted"

	AccountTypeCash     AccountType = "cash"
	AccountTypeCcard    AccountType = "ccard"
	AccountTypeChecking AccountType = "checking"
	AccountTypeLoan     AccountType = "loan"
	AccountTypeDeposit  AccountType = "deposit"
	AccountTypeEmoney   AccountType = "emoney"
	AccountTypeDebt     AccountType = "debt"

	DateIntervalDay   DateInterval = "day"
	DateIntervalWeek  DateInterval = "week"
	DateIntervalMonth DateInterval = "month"
	DateIntervalYear  DateInterval = "year"

	AccountEndDateOffsetDay   AccountEndDateOffset = AccountEndDateOffset(DateIntervalDay)
	AccountEndDateOffsetWeek  AccountEndDateOffset = AccountEndDateOffset(DateIntervalWeek)
	AccountEndDateOffsetMonth AccountEndDateOffset = AccountEndDateOffset(DateIntervalMonth)
	AccountEndDateOffsetYear  AccountEndDateOffset = AccountEndDateOffset(DateIntervalYear)

	AccountPayoffIntervalMonth AccountPayoffInterval = AccountPayoffInterval(DateIntervalMonth)
	AccountPayoffIntervalYear  AccountPayoffInterval = AccountPayoffInterval(DateIntervalYear)

	ReminderIntervalDay   ReminderInterval = ReminderInterval(DateIntervalDay)
	ReminderIntervalWeek  ReminderInterval = ReminderInterval(DateIntervalWeek)
	ReminderIntervalMonth ReminderInterval = ReminderInterval(DateIntervalMonth)
	ReminderIntervalYear  ReminderInterval = ReminderInterval(DateIntervalYear)

	BudgetMonthTag = "00000000-0000-0000-0000-000000000000"
)

func generateUuid() string {
	return fmt.Sprintf("%s", uuid.NewV4())
}

func (e *Account) TouchChanged() () {
	e.Changed = GetCurrentTimestamp()
}

func (e *Tag) TouchChanged() () {
	e.Changed = GetCurrentTimestamp()
}

func (e *Merchant) TouchChanged() () {
	e.Changed = GetCurrentTimestamp()
}

func (e *Reminder) TouchChanged() () {
	e.Changed = GetCurrentTimestamp()
}

func (e *ReminderMarker) TouchChanged() () {
	e.Changed = GetCurrentTimestamp()
}

func (e *Transaction) TouchChanged() () {
	e.Changed = GetCurrentTimestamp()
}

func (e *Budget) TouchChanged() () {
	e.Changed = GetCurrentTimestamp()
}
