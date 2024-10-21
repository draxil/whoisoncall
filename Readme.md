# who is on call?

Do you:
 
a) Use pager duty

b) Have to be on call 

c) Sometimes wonder if it's you today, and are too lazy to open your calendar?

Wow, remarkably you have come to the right place!

Install this thing (you need golang):

```sh
go install github.com/draxil/whoisoncall@latest
```

Now do:

```sh
whoisoncall <the ics URL for your team>
```

As a fellow lazy person you probably want to wrap this up as a shell script so you don't need that URL every time!

# limitations & notes

This is currently in a "done in five minutes to scratch an itch" state, so:

1. Assumes this is the team cal, and someone is always on call.
2. Assumes every day *someone* needs to be on call.

I may amend for the above so it would work with a personal calendar (am I on call?), or a team that only has limited windows of responsibility. If you somehow are reading this and care about these things.. Please pop me an issue!

This refers to pager duty, but really any http available ics file which only has one event per day with a useful summary would work.
