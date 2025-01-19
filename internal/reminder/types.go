package reminder

const birthdayType = "birthday"
const deadlineType = "deadline"
const dateType = "date"
const untilType = "until"

type OnyxExpression struct {
	Type    string
	Content string
}
