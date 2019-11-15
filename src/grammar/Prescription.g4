// Prescription.g4
// $antlr-format on
grammar Prescription;

import Date;

// Global
WHITESPACE: [ \t\n\r]+ -> skip;
COMMENT: '/*' .*? '*/' -> skip;
LINE_COMMENT: '//' ~[\r\n]* -> skip;

// Keywords
STARTING: S T A R T I N G;
UNTIL: U N T I L;
THEN: T H E N;
FOR: F O R;
EVERY: E V E R Y;

prescription
  : prescription_definition+
  ;

prescription_definition
  : STARTING date_period recipe 
  | THEN date_period recipe
  ;

date_period
  : date_period_start? date_period_end
  ;

date_period_start
  : date
  ;

date_period_end
  : FOR date_recurrence
  | UNTIL date
  ;

date_recurrence
  : recurrence_value DATE_TIME_UNIT alert_repeater?
  ;

recurrence_value
  : UNSIGNED_INT
  ;

alert_repeater
  : EVERY recurrence_value DATE_TIME_UNIT
  ;

recipe
  : IDENTIFIER+
  ;