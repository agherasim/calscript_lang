// Code generated from Calscript.g4 by ANTLR 4.7.2. DO NOT EDIT.

package calscript_lang // Calscript
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalscriptListener is a complete listener for a parse tree produced by CalscriptParser.
type CalscriptListener interface {
	antlr.ParseTreeListener

	// EnterCalscript is called when entering the calscript production.
	EnterCalscript(c *CalscriptContext)

	// EnterBegin is called when entering the begin production.
	EnterBegin(c *BeginContext)

	// EnterScript_name is called when entering the script_name production.
	EnterScript_name(c *Script_nameContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterEvent is called when entering the event production.
	EnterEvent(c *EventContext)

	// EnterStart_end is called when entering the start_end production.
	EnterStart_end(c *Start_endContext)

	// EnterDate_period is called when entering the date_period production.
	EnterDate_period(c *Date_periodContext)

	// EnterDate_period_start is called when entering the date_period_start production.
	EnterDate_period_start(c *Date_period_startContext)

	// EnterDate_period_end is called when entering the date_period_end production.
	EnterDate_period_end(c *Date_period_endContext)

	// EnterRecurrence is called when entering the recurrence production.
	EnterRecurrence(c *RecurrenceContext)

	// EnterReminders is called when entering the reminders production.
	EnterReminders(c *RemindersContext)

	// EnterRecurrence_value is called when entering the recurrence_value production.
	EnterRecurrence_value(c *Recurrence_valueContext)

	// EnterRecurrence_value_unit is called when entering the recurrence_value_unit production.
	EnterRecurrence_value_unit(c *Recurrence_value_unitContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// EnterAttendees is called when entering the attendees production.
	EnterAttendees(c *AttendeesContext)

	// EnterAttendee is called when entering the attendee production.
	EnterAttendee(c *AttendeeContext)

	// EnterLocation is called when entering the location production.
	EnterLocation(c *LocationContext)

	// EnterLocation_string is called when entering the location_string production.
	EnterLocation_string(c *Location_stringContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// EnterFormal_date_time is called when entering the formal_date_time production.
	EnterFormal_date_time(c *Formal_date_timeContext)

	// EnterFormal_date is called when entering the formal_date production.
	EnterFormal_date(c *Formal_dateContext)

	// EnterYear_month_day is called when entering the year_month_day production.
	EnterYear_month_day(c *Year_month_dayContext)

	// EnterMonth_day is called when entering the month_day production.
	EnterMonth_day(c *Month_dayContext)

	// EnterYear is called when entering the year production.
	EnterYear(c *YearContext)

	// EnterMonth is called when entering the month production.
	EnterMonth(c *MonthContext)

	// EnterVocal_month is called when entering the vocal_month production.
	EnterVocal_month(c *Vocal_monthContext)

	// EnterNumeric_month is called when entering the numeric_month production.
	EnterNumeric_month(c *Numeric_monthContext)

	// EnterDay_of_month is called when entering the day_of_month production.
	EnterDay_of_month(c *Day_of_monthContext)

	// EnterDate_separator is called when entering the date_separator production.
	EnterDate_separator(c *Date_separatorContext)

	// EnterTime is called when entering the time production.
	EnterTime(c *TimeContext)

	// EnterNatural_date is called when entering the natural_date production.
	EnterNatural_date(c *Natural_dateContext)

	// ExitCalscript is called when exiting the calscript production.
	ExitCalscript(c *CalscriptContext)

	// ExitBegin is called when exiting the begin production.
	ExitBegin(c *BeginContext)

	// ExitScript_name is called when exiting the script_name production.
	ExitScript_name(c *Script_nameContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitEvent is called when exiting the event production.
	ExitEvent(c *EventContext)

	// ExitStart_end is called when exiting the start_end production.
	ExitStart_end(c *Start_endContext)

	// ExitDate_period is called when exiting the date_period production.
	ExitDate_period(c *Date_periodContext)

	// ExitDate_period_start is called when exiting the date_period_start production.
	ExitDate_period_start(c *Date_period_startContext)

	// ExitDate_period_end is called when exiting the date_period_end production.
	ExitDate_period_end(c *Date_period_endContext)

	// ExitRecurrence is called when exiting the recurrence production.
	ExitRecurrence(c *RecurrenceContext)

	// ExitReminders is called when exiting the reminders production.
	ExitReminders(c *RemindersContext)

	// ExitRecurrence_value is called when exiting the recurrence_value production.
	ExitRecurrence_value(c *Recurrence_valueContext)

	// ExitRecurrence_value_unit is called when exiting the recurrence_value_unit production.
	ExitRecurrence_value_unit(c *Recurrence_value_unitContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)

	// ExitAttendees is called when exiting the attendees production.
	ExitAttendees(c *AttendeesContext)

	// ExitAttendee is called when exiting the attendee production.
	ExitAttendee(c *AttendeeContext)

	// ExitLocation is called when exiting the location production.
	ExitLocation(c *LocationContext)

	// ExitLocation_string is called when exiting the location_string production.
	ExitLocation_string(c *Location_stringContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)

	// ExitFormal_date_time is called when exiting the formal_date_time production.
	ExitFormal_date_time(c *Formal_date_timeContext)

	// ExitFormal_date is called when exiting the formal_date production.
	ExitFormal_date(c *Formal_dateContext)

	// ExitYear_month_day is called when exiting the year_month_day production.
	ExitYear_month_day(c *Year_month_dayContext)

	// ExitMonth_day is called when exiting the month_day production.
	ExitMonth_day(c *Month_dayContext)

	// ExitYear is called when exiting the year production.
	ExitYear(c *YearContext)

	// ExitMonth is called when exiting the month production.
	ExitMonth(c *MonthContext)

	// ExitVocal_month is called when exiting the vocal_month production.
	ExitVocal_month(c *Vocal_monthContext)

	// ExitNumeric_month is called when exiting the numeric_month production.
	ExitNumeric_month(c *Numeric_monthContext)

	// ExitDay_of_month is called when exiting the day_of_month production.
	ExitDay_of_month(c *Day_of_monthContext)

	// ExitDate_separator is called when exiting the date_separator production.
	ExitDate_separator(c *Date_separatorContext)

	// ExitTime is called when exiting the time production.
	ExitTime(c *TimeContext)

	// ExitNatural_date is called when exiting the natural_date production.
	ExitNatural_date(c *Natural_dateContext)
}
