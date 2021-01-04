# go-projects
Repository for Go projects

## Source
The basics and examples are taken from the [Go Tour](https://tour.golang.org/list)

The [Go Documentation](https://golang.org/doc/) is also useful

More resources can be found at the end of the tour, [here](https://tour.golang.org/concurrency/11)

## How to run

You run a Go file like this:
```
go run FILEPATH
```

## How to use local packages

Inside the directory, run: 
```
go mod init MOD_NAME
```

Then you have to build the package: 
```
go build
```

## Projects
### HTTP SHA-256 Encryption Server
To run the server: 
```
go run http/server.go &
```

To test the **single message encryption**:
```
curl -G localhost:8090/encrypt-single --data-urlencode "msg=hello world"
``` 

To test the **multiple message encryption**:
```
curl -XPOST -d "msgs=abc'de&msgs=hello there&msgs=abc" localhost:8090/encrypt-multiple
```

To test the **multiple message encryption** with a very big request:
```
curl -XPOST -d "msgs=lets do this&msgs=this is the real test&msgs=it is gonna be real big this request&msgs=I can make sure we are using at least 10 items&msgs=I dont really know if this is gonna break it, but it should at least take longer&msgs=Love For All, Hatred For None. – Khalifatul Masih III&msgs=Change the world by being yourself. – Amy Poehler&msgs=Every moment is a fresh beginning. – T.S Eliot&msgs=Never regret anything that made you smile. – Mark Twain&msgs=Die with memories, not dreams. – Unknown&msgs=Aspire to inspire before we expire. – Unknown&msgs=Everything you can imagine is real. – Pablo Picasso&msgs=Simplicity is the ultimate sophistication. – Leonardo da Vinci&msgs=Whatever you do, do it well. – Walt Disney&msgs=What we think, we become. – Buddha&msgs=All limitations are self-imposed. – Oliver Wendell Holmes&msgs=Tough times never last but tough people do. – Robert H. Schiuller&msgs=Problems are not stop signs, they are guidelines. – Robert H. Schiuller&msgs=One day the people that don’t even believe in you will tell everyone how they met you. – Johnny Depp&msgs=If I’m gonna tell a real story, I’m gonna start with my name. – Kendrick Lamar&msgs=If you tell the truth you don’t have to remember anything. – Mark Twain&msgs=Have enough courage to start and enough heart to finish. – Jessica N. S. Yourko&msgs=Hate comes from intimidation, love comes from appreciation. – Tyga&msgs=I could agree with you but then we’d both be wrong. – Harvey Specter&msgs=Oh, the things you can find, if you don’t stay behind. – Dr. Seuss&msgs=Determine your priorities and focus on them. – Eileen McDargh&msgs=Be so good they can’t ignore you. – Steve Martin&msgs=Dream as if you’ll live forever, live as if you’ll die today. – James Dean&msgs=Yesterday you said tomorrow. Just do it. – Nike&msgs=I don’t need it to be easy, I need it to be worth it. – Lil Wayne&msgs=Never let your emotions overpower your intelligence. – Drake&msgs=Nothing lasts forever but at least we got these memories. – J. Cole&msgs=Don’t you know your imperfections is a blessing? – Kendrick Lamar&msgs=Reality is wrong, dreams are for real. – Tupac&msgs=To live will be an awfully big adventure. – Peter Pan&msgs=Try to be a rainbow in someone’s cloud. – Maya Angelou&msgs=There is no substitute for hard work. – Thomas Edison&msgs=What consumes your mind controls your life- Unknown&msgs=Strive for greatness. – Lebron James&msgs=Wanting to be someone else is a waste of who you are. – Kurt Cobain&msgs=And still, I rise. – Maya Angelou&msgs=The time is always right to do what is right. – Martin Luther King Jr.&msgs=Let the beauty of what you love be what you do. – Rumi&msgs=May your choices reflect your hopes, not your fears. – Nelson Mandela&msgs=A happy soul is the best shield for a cruel world. – Atticus&msgs=White is not always light and black is not always dark. – Habeeb Akande&msgs=Life becomes easier when you learn to accept the apology you never got. – R. Brault&msgs=Happiness depends upon ourselves. – Aristotle&msgs=Turn your wounds into wisdom. – Oprah Winfrey&msgs=Change the game, don’t let the game change you. – Macklemore&msgs=It hurt because it mattered. – John Green&msgs=If the world was blind how many people would you impress? – Boonaa Mohammed&msgs=I will remember and recover, not forgive and forget. – Unknown&msgs=The meaning of life is to give life meaning. – Ken Hudgins&msgs=The true meaning of life is to plant trees, under whose shade you do not expect to sit. – Nelson Henderson&msgs=When words fail, music speaks. – Shakespeare&msgs=Embrace the glorious mess that you are. – Elizabeth Gilbert&msgs=Normality is a paved road: it’s comfortable to walk but no flowers grow. – Vincent van Gogh&msgs=I have nothing to lose but something to gain. – Eminem" localhost:8090/encrypt-multiple
``` 

To test the **multiple parallel message encryption**:
```
curl -XPOST -d "msgs=abc'de&msgs=hello there&msgs=abc" localhost:8090/encrypt-multiple-parallel
```

To test the **multiple parallel message encryption** with a very big request:
```
curl -XPOST -d "msgs=lets do this&msgs=this is the real test&msgs=it is gonna be real big this request&msgs=I can make sure we are using at least 10 items&msgs=I dont really know if this is gonna break it, but it should at least take longer&msgs=Love For All, Hatred For None. – Khalifatul Masih III&msgs=Change the world by being yourself. – Amy Poehler&msgs=Every moment is a fresh beginning. – T.S Eliot&msgs=Never regret anything that made you smile. – Mark Twain&msgs=Die with memories, not dreams. – Unknown&msgs=Aspire to inspire before we expire. – Unknown&msgs=Everything you can imagine is real. – Pablo Picasso&msgs=Simplicity is the ultimate sophistication. – Leonardo da Vinci&msgs=Whatever you do, do it well. – Walt Disney&msgs=What we think, we become. – Buddha&msgs=All limitations are self-imposed. – Oliver Wendell Holmes&msgs=Tough times never last but tough people do. – Robert H. Schiuller&msgs=Problems are not stop signs, they are guidelines. – Robert H. Schiuller&msgs=One day the people that don’t even believe in you will tell everyone how they met you. – Johnny Depp&msgs=If I’m gonna tell a real story, I’m gonna start with my name. – Kendrick Lamar&msgs=If you tell the truth you don’t have to remember anything. – Mark Twain&msgs=Have enough courage to start and enough heart to finish. – Jessica N. S. Yourko&msgs=Hate comes from intimidation, love comes from appreciation. – Tyga&msgs=I could agree with you but then we’d both be wrong. – Harvey Specter&msgs=Oh, the things you can find, if you don’t stay behind. – Dr. Seuss&msgs=Determine your priorities and focus on them. – Eileen McDargh&msgs=Be so good they can’t ignore you. – Steve Martin&msgs=Dream as if you’ll live forever, live as if you’ll die today. – James Dean&msgs=Yesterday you said tomorrow. Just do it. – Nike&msgs=I don’t need it to be easy, I need it to be worth it. – Lil Wayne&msgs=Never let your emotions overpower your intelligence. – Drake&msgs=Nothing lasts forever but at least we got these memories. – J. Cole&msgs=Don’t you know your imperfections is a blessing? – Kendrick Lamar&msgs=Reality is wrong, dreams are for real. – Tupac&msgs=To live will be an awfully big adventure. – Peter Pan&msgs=Try to be a rainbow in someone’s cloud. – Maya Angelou&msgs=There is no substitute for hard work. – Thomas Edison&msgs=What consumes your mind controls your life- Unknown&msgs=Strive for greatness. – Lebron James&msgs=Wanting to be someone else is a waste of who you are. – Kurt Cobain&msgs=And still, I rise. – Maya Angelou&msgs=The time is always right to do what is right. – Martin Luther King Jr.&msgs=Let the beauty of what you love be what you do. – Rumi&msgs=May your choices reflect your hopes, not your fears. – Nelson Mandela&msgs=A happy soul is the best shield for a cruel world. – Atticus&msgs=White is not always light and black is not always dark. – Habeeb Akande&msgs=Life becomes easier when you learn to accept the apology you never got. – R. Brault&msgs=Happiness depends upon ourselves. – Aristotle&msgs=Turn your wounds into wisdom. – Oprah Winfrey&msgs=Change the game, don’t let the game change you. – Macklemore&msgs=It hurt because it mattered. – John Green&msgs=If the world was blind how many people would you impress? – Boonaa Mohammed&msgs=I will remember and recover, not forgive and forget. – Unknown&msgs=The meaning of life is to give life meaning. – Ken Hudgins&msgs=The true meaning of life is to plant trees, under whose shade you do not expect to sit. – Nelson Henderson&msgs=When words fail, music speaks. – Shakespeare&msgs=Embrace the glorious mess that you are. – Elizabeth Gilbert&msgs=Normality is a paved road: it’s comfortable to walk but no flowers grow. – Vincent van Gogh&msgs=I have nothing to lose but something to gain. – Eminem" localhost:8090/encrypt-multiple-parallel
``` 

### Blockchain Server
To run the server:
```
go run http/blockchainServer.go &
```

To **get chain**:
```
curl -G localhost:8090/chain
```

To **add a transaction**:
```
curl -XPOST -d "sender=SENDER&receiver=RECEIVER&amount=10" localhost:8090/transactions
```

To **mine the block**:
```
curl -XPOST localhost:8090/mine
```