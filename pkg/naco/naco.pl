#   /usr/bin/env perl
use warnings;
use strict;

use Text::Normalize::NACO qw( naco_normalize );
use MARC::Field::Normalize::NACO qw(naco_from_string);

my @o = ("ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷øùúûüýþÿ",
	"⁰¹²³⁴⁵⁶⁷⁸⁹₀₁₂₃₄₅₆₇₈₉",
	"ÆŒ",

 );





 foreach my $original (@o) {
#my $normalized = naco_normalize( $original );
my $normalized = naco_from_string( $original );

print $normalized . "\n";

}
