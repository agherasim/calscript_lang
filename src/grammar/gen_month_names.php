<?php

$month_names = [
    'JANUARY',
    'FEBRUARY',
    'MARCH',
    'APRIL',
    'MAY',
    'JUNE',
    'JULY',
    'AUGUST',
    'SEPTEMBER',
    'OCTOBER',
    'NOVEMBER',
    'DECEMBER',
];

foreach ($month_names as $month_name) {
    printf(
        "%s: %s | %s;\n", 
        $month_name, 
        implode(' ', str_split($month_name)),
        implode(' ', str_split(substr($month_name, 0, 3)))
    );
}