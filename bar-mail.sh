#!/bin/bash

# COPY THIS FILE TO: /usr/local/bin/bar-mail.sh

# week2date 2023 14 1 ergibt das Datum (DD.MM.) vom Montag der KW 14,
# week2date 2023 14 7 liefert das Datum des Sonntags der KW14 zurueck.
# siehe https://stackoverflow.com/questions/15606567/unix-date-how-to-convert-week-number-date-w-to-a-date-range-mon-sun
function week2date () {
	local year=$1
	local week=$2
	local dayofweek=$3
	date -d "$year-01-01 +$(( $week * 7 + 1 - $(date -d "$year-01-04" +%u ) - 3 )) days -2 days + $dayofweek days" +"%d.%m."
}

# Aktuelle Kalenderwoche ermitteln und dann plus 1 drauf rechnen.
KW=$((`date +%V | sed 's/^0//'`+1))
YEAR=`date +%Y`

# Montag - Sonntag
MON=`week2date $YEAR $KW 1`
DIE=`week2date $YEAR $KW 2`
MIT=`week2date $YEAR $KW 3`
DON=`week2date $YEAR $KW 4`
FRE=`week2date $YEAR $KW 5`
SAM=`week2date $YEAR $KW 6`
SON=`week2date $YEAR $KW 7`

SUB="Barplan KW${KW} ${MON}-${SON}${YEAR}"

# UNCOMMENT HERE TO USE THE CORRECT EMAIL ADDR.
# TO=xxx-xxxx@c-base.org
TO=xxx@c-base.org

# UNCOMMENT FOR TESTING WITHOUT SENDING AN EMAIL:
# echo $SUB
# exit 0

cat | mail -s "$SUB" -a "From: vorstand@c-base.org" ${TO} << EOF
Im folgenden der Barplan fuer naechste Woche. Wer Schichten uebernehmen
moechte, moege sich bitte einfach eintragen und den Plan als Antwort
wieder auf die Liste schicken. Die angegebenen Zeiten dienen der
Orientierung, ihr koennt sie euch gerne anpassen.

Wenn dann jeder nur auf die letzte Mail im Thread antwortet, sollte das
eigentlich ganz gut funktionieren...

Mo, $MON:
15:00 -> 20:00:
19:00 -> 00:30:

Di, $DIE:
15:00 -> 20:00:
19:00 -> 00:30:

Mi, $MIT:
15:00 -> 20:00:
19:00 -> 00:30:

Do, $DON:
15:00 -> 20:00:
19:00 -> 00:30:

Fr, $FRE:
15:00 -> 20:00:
19:00 -> 00:30:

Sa, $SAM:
14:00 -> 20:00:
19:00 -> 00:30:

So, $SON:
15:00 -> 20:00:
19:00 -> 00:30:


Dies ist eine automatisch gesendete Nachricht, sie ist
ohne Unterschrift gueltig.
EOF

