// Code generated from Calscript.g4 by ANTLR 4.7.2. DO NOT EDIT.

package calscript_lang // Calscript
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCalscriptListener is a complete listener for a parse tree produced by CalscriptParser.
type BaseCalscriptListener struct{}

var _ CalscriptListener = &BaseCalscriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalscriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalscriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalscriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalscriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCalscript is called when production calscript is entered.
func (s *BaseCalscriptListener) EnterCalscript(ctx *CalscriptContext) {}

// ExitCalscript is called when production calscript is exited.
func (s *BaseCalscriptListener) ExitCalscript(ctx *CalscriptContext) {}

// EnterBegin is called when production begin is entered.
func (s *BaseCalscriptListener) EnterBegin(ctx *BeginContext) {}

// ExitBegin is called when production begin is exited.
func (s *BaseCalscriptListener) ExitBegin(ctx *BeginContext) {}

// EnterScript_name is called when production script_name is entered.
func (s *BaseCalscriptListener) EnterScript_name(ctx *Script_nameContext) {}

// ExitScript_name is called when production script_name is exited.
func (s *BaseCalscriptListener) ExitScript_name(ctx *Script_nameContext) {}

// EnterName is called when production name is entered.
func (s *BaseCalscriptListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseCalscriptListener) ExitName(ctx *NameContext) {}

// EnterEvent is called when production event is entered.
func (s *BaseCalscriptListener) EnterEvent(ctx *EventContext) {}

// ExitEvent is called when production event is exited.
func (s *BaseCalscriptListener) ExitEvent(ctx *EventContext) {}

// EnterStart_end is called when production start_end is entered.
func (s *BaseCalscriptListener) EnterStart_end(ctx *Start_endContext) {}

// ExitStart_end is called when production start_end is exited.
func (s *BaseCalscriptListener) ExitStart_end(ctx *Start_endContext) {}

// EnterDate_period is called when production date_period is entered.
func (s *BaseCalscriptListener) EnterDate_period(ctx *Date_periodContext) {}

// ExitDate_period is called when production date_period is exited.
func (s *BaseCalscriptListener) ExitDate_period(ctx *Date_periodContext) {}

// EnterDate_period_start is called when production date_period_start is entered.
func (s *BaseCalscriptListener) EnterDate_period_start(ctx *Date_period_startContext) {}

// ExitDate_period_start is called when production date_period_start is exited.
func (s *BaseCalscriptListener) ExitDate_period_start(ctx *Date_period_startContext) {}

// EnterDate_period_end is called when production date_period_end is entered.
func (s *BaseCalscriptListener) EnterDate_period_end(ctx *Date_period_endContext) {}

// ExitDate_period_end is called when production date_period_end is exited.
func (s *BaseCalscriptListener) ExitDate_period_end(ctx *Date_period_endContext) {}

// EnterRecurrence is called when production recurrence is entered.
func (s *BaseCalscriptListener) EnterRecurrence(ctx *RecurrenceContext) {}

// ExitRecurrence is called when production recurrence is exited.
func (s *BaseCalscriptListener) ExitRecurrence(ctx *RecurrenceContext) {}

// EnterReminders is called when production reminders is entered.
func (s *BaseCalscriptListener) EnterReminders(ctx *RemindersContext) {}

// ExitReminders is called when production reminders is exited.
func (s *BaseCalscriptListener) ExitReminders(ctx *RemindersContext) {}

// EnterRecurrence_value is called when production recurrence_value is entered.
func (s *BaseCalscriptListener) EnterRecurrence_value(ctx *Recurrence_valueContext) {}

// ExitRecurrence_value is called when production recurrence_value is exited.
func (s *BaseCalscriptListener) ExitRecurrence_value(ctx *Recurrence_valueContext) {}

// EnterRecurrence_value_unit is called when production recurrence_value_unit is entered.
func (s *BaseCalscriptListener) EnterRecurrence_value_unit(ctx *Recurrence_value_unitContext) {}

// ExitRecurrence_value_unit is called when production recurrence_value_unit is exited.
func (s *BaseCalscriptListener) ExitRecurrence_value_unit(ctx *Recurrence_value_unitContext) {}

// EnterDescription is called when production description is entered.
func (s *BaseCalscriptListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BaseCalscriptListener) ExitDescription(ctx *DescriptionContext) {}

// EnterAttendees is called when production attendees is entered.
func (s *BaseCalscriptListener) EnterAttendees(ctx *AttendeesContext) {}

// ExitAttendees is called when production attendees is exited.
func (s *BaseCalscriptListener) ExitAttendees(ctx *AttendeesContext) {}

// EnterAttendee is called when production attendee is entered.
func (s *BaseCalscriptListener) EnterAttendee(ctx *AttendeeContext) {}

// ExitAttendee is called when production attendee is exited.
func (s *BaseCalscriptListener) ExitAttendee(ctx *AttendeeContext) {}

// EnterLocation is called when production location is entered.
func (s *BaseCalscriptListener) EnterLocation(ctx *LocationContext) {}

// ExitLocation is called when production location is exited.
func (s *BaseCalscriptListener) ExitLocation(ctx *LocationContext) {}

// EnterLocation_string is called when production location_string is entered.
func (s *BaseCalscriptListener) EnterLocation_string(ctx *Location_stringContext) {}

// ExitLocation_string is called when production location_string is exited.
func (s *BaseCalscriptListener) ExitLocation_string(ctx *Location_stringContext) {}

// EnterDate is called when production date is entered.
func (s *BaseCalscriptListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *BaseCalscriptListener) ExitDate(ctx *DateContext) {}

// EnterFormal_date_time is called when production formal_date_time is entered.
func (s *BaseCalscriptListener) EnterFormal_date_time(ctx *Formal_date_timeContext) {}

// ExitFormal_date_time is called when production formal_date_time is exited.
func (s *BaseCalscriptListener) ExitFormal_date_time(ctx *Formal_date_timeContext) {}

// EnterFormal_date is called when production formal_date is entered.
func (s *BaseCalscriptListener) EnterFormal_date(ctx *Formal_dateContext) {}

// ExitFormal_date is called when production formal_date is exited.
func (s *BaseCalscriptListener) ExitFormal_date(ctx *Formal_dateContext) {}

// EnterYear_month_day is called when production year_month_day is entered.
func (s *BaseCalscriptListener) EnterYear_month_day(ctx *Year_month_dayContext) {}

// ExitYear_month_day is called when production year_month_day is exited.
func (s *BaseCalscriptListener) ExitYear_month_day(ctx *Year_month_dayContext) {}

// EnterMonth_day is called when production month_day is entered.
func (s *BaseCalscriptListener) EnterMonth_day(ctx *Month_dayContext) {}

// ExitMonth_day is called when production month_day is exited.
func (s *BaseCalscriptListener) ExitMonth_day(ctx *Month_dayContext) {}

// EnterYear is called when production year is entered.
func (s *BaseCalscriptListener) EnterYear(ctx *YearContext) {}

// ExitYear is called when production year is exited.
func (s *BaseCalscriptListener) ExitYear(ctx *YearContext) {}

// EnterMonth is called when production month is entered.
func (s *BaseCalscriptListener) EnterMonth(ctx *MonthContext) {}

// ExitMonth is called when production month is exited.
func (s *BaseCalscriptListener) ExitMonth(ctx *MonthContext) {}

// EnterVocal_month is called when production vocal_month is entered.
func (s *BaseCalscriptListener) EnterVocal_month(ctx *Vocal_monthContext) {}

// ExitVocal_month is called when production vocal_month is exited.
func (s *BaseCalscriptListener) ExitVocal_month(ctx *Vocal_monthContext) {}

// EnterNumeric_month is called when production numeric_month is entered.
func (s *BaseCalscriptListener) EnterNumeric_month(ctx *Numeric_monthContext) {}

// ExitNumeric_month is called when production numeric_month is exited.
func (s *BaseCalscriptListener) ExitNumeric_month(ctx *Numeric_monthContext) {}

// EnterDay_of_month is called when production day_of_month is entered.
func (s *BaseCalscriptListener) EnterDay_of_month(ctx *Day_of_monthContext) {}

// ExitDay_of_month is called when production day_of_month is exited.
func (s *BaseCalscriptListener) ExitDay_of_month(ctx *Day_of_monthContext) {}

// EnterDate_separator is called when production date_separator is entered.
func (s *BaseCalscriptListener) EnterDate_separator(ctx *Date_separatorContext) {}

// ExitDate_separator is called when production date_separator is exited.
func (s *BaseCalscriptListener) ExitDate_separator(ctx *Date_separatorContext) {}

// EnterTime is called when production time is entered.
func (s *BaseCalscriptListener) EnterTime(ctx *TimeContext) {}

// ExitTime is called when production time is exited.
func (s *BaseCalscriptListener) ExitTime(ctx *TimeContext) {}

// EnterNatural_date is called when production natural_date is entered.
func (s *BaseCalscriptListener) EnterNatural_date(ctx *Natural_dateContext) {}

// ExitNatural_date is called when production natural_date is exited.
func (s *BaseCalscriptListener) ExitNatural_date(ctx *Natural_dateContext) {}
