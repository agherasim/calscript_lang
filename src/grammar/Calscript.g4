// CalScript.g4
// $antlr-format on
grammar Calscript;

import Date;

// Global
WHITESPACE: [ \t\n\r]+ -> skip;
COMMENT: '/*' .*? '*/' -> skip;
LINE_COMMENT: '//' ~[\r\n]* -> skip;
LIST_SEPARATOR: COMMA;

// Keywords
ON: O N;
STARTING: S T A R T I N G;
UNTIL: U N T I L;
THEN: T H E N;
FOR: F O R;
EVERY: E V E R Y;
WITH: W I T H;
AT: A T;

calscript
  : begin event+
  ;

begin
  : 'calscript' script_name
  ;

script_name
  : IDENTIFIER+
  ;

name
  : IDENTIFIER
  ;

event
  : start_end description attendees?
  ;

start_end
  : ON date
  | STARTING date_period
  | THEN date_period
  ;

date_period
  : date_period_start? date_period_end
  ;

date_period_start
  : date
  ;

date_period_end
  : FOR recurrence
  | UNTIL recurrence
  ;

recurrence
  : recurrence_value reminders?
  | date reminders?
  ;

reminders
  : EVERY recurrence_value
  ;

recurrence_value
  : UNSIGNED_INT recurrence_value_unit
  ;

recurrence_value_unit
  : DAY
  | WEEK
  | MONTH
  | YEAR
  | HOUR
  | MINUTE
  | SECOND
  ;

description
  : STRING
  ;

attendees
  : WITH ':' attendee (LIST_SEPARATOR attendee)+
  ;

attendee
  : IDENTIFIER
  ;

// @todo fix location, causes unresponsivess
location
  : AT ':' location_string
  ;

location_string
  : IDENTIFIER
  ;