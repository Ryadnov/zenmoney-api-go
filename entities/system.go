package entities

type (
	Instrument struct {
		Id         InstrumentId
		Changed    Timestamp
		Title      string
		ShortTitle string
		Symbol     string
		Rate       float64
	}

	Company struct {
		Id        CompanyId
		Changed   Timestamp
		Title     string
		FullTitle string
		Www       string
		Country   string
	}

	User struct {
		Id       UserId
		Changed  Timestamp
		Login    *string
		Currency InstrumentId
		Parent   *UserId
	}

	//Deletion struct {
	//	Id       ObjectId
	//	Changed  Timestamp
	//	Login    *string
	//	Currency InstrumentId
	//	Parent   *UserId
	//}

	//	deletion: [
	//{
	//	id:     String -> Object.id
	//	object: String -> Object.class
	//	stamp:  Int
	//	user:   Int
	//}
	//]?
)
