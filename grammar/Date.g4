// Date.g4
// $antlr-format on
grammar Date;

import CommonLexer;

// Relative Dates
DAY: D A Y S? | D;
WEEK: W E E K S? | W;
MONTH: M O N T H S? | M;
YEAR: Y E A R S? | Y;
HOUR: H O U R S? | H;
MINUTE: M I N U T E S? | M I N;
SECOND: S E C O N D S? | S E C;

// Date Literals
YESTERDAY: Y E S T E R D A Y;
NOW: N O W;
TOMORROW: T O M O R R O W;
TODAY: T O D A Y;

JANUARY: J A N U A R Y | J A N;
FEBRUARY: F E B R U A R Y | F E B;
MARCH: M A R C H | M A R;
APRIL: A P R I L | A P R;
MAY: M A Y | M A Y;
JUNE: J U N E | J U N;
JULY: J U L Y | J U L;
AUGUST: A U G U S T | A U G;
SEPTEMBER: S E P T E M B E R | S E P;
OCTOBER: O C T O B E R | O C T;
NOVEMBER: N O V E M B E R | N O V;
DECEMBER: D E C E M B E R | D E C;

date
  : natural_date
  | formal_date
  ;

formal_date_time
  : formal_date .? time
  ;

formal_date
  : year_month_day
  // | month_day_year
  // | day_month_year
  // | day_month
  // | month_year
  ;

year_month_day
  : year date_separator? month_day
  ;

// month_day_year
//   : month_day date_separator? year
//   ;

// day_month_year
//   : day_month date_separator? year
//   ;

// day_month
//   : day_of_month date_separator? month
//   ;

month_day
  : month date_separator? day_of_month
  ;

// month_year
//   : month date_separator? year
//   ;

year
  : UNSIGNED_INT+
  ;

month
  : vocal_month
  | numeric_month
  ;

vocal_month
  : JANUARY | FEBRUARY | MARCH | APRIL | MAY | JUNE | JULY | AUGUST 
  | SEPTEMBER | OCTOBER | NOVEMBER | DECEMBER
  ;

numeric_month
  : UNSIGNED_INT+
  ;

day_of_month
  : UNSIGNED_INT+
  ;

date_separator
  : COMMA
  | DOT
  | COLON
  | SLASH
  | DASH
  ;

time
  :
  ;
// time
//   : hour_minute am_pm timezone
//   | hour_minute am_pm
//   | hour_minute timezone
//   | hour_minute
//   ;

natural_date
  : YESTERDAY
  | NOW
  | TOMORROW
  | TODAY
  ;