# bar-mail.sh
Shell script that creates a weekly e-mail for bar shifts, but requiring a
golang build to build the mail body from the c-base calendar.

Run `make install`

Copy c-base.cron to /etc/cron.d/c-base
