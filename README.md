# go-exercises
Go exercises from calhoun.io

### List of execises

- [x] Quiz Game
- [x] Url Shortener
- [ ] Choose Your Own Adventure <---- WIP
- [ ] HTML Link Parser
- [ ] Sitemap Builder
- [ ] Hacker Rank Problem
- [ ] CLI Task Manager
- [ ] Phone Number Normalizer
- [ ] Deck of Cards
- [ ] Blackjack
- [ ] Blackjack AI
- [ ] File Renaming Tool
- [ ] Quiet HN
- [ ] Recover Middleware
- [ ] Recover Middleware w/ Source Code
- [ ] Twitter Retweet Contest CLI
- [ ] Secrets API and CLI
- [ ] Image Transform Service
- [ ] Building Images (png & svg)
- [ ] Building PDFs

### Quiz Game

Problems reading from standard input and turning into string several characters like end of line.
I learnd how to use channels and concurrent threds

### Url Shortener

I learnd how to use http.Handler and http.HandlerFunc
I had problems with yaml unmarshall due to tabs into []byte objects (delete any tab char from that objects in this unmarshall actions)
I had little isuee understanding correct behaviour of defined struct till I saw I should pass an array of my defined struct into unmarshall action

- [ ] read from yaml file instead of variable in yaml format
- [ ] read from json
- [ ] read from database


### Choose Your Own Adventure

When some object is created for some template you have to create public params, otherwise parse method could not parse that values.
If you want to create some array in execution cycle it has to be with make method otherwise you could throw some outofbounds exception when you use that array in future