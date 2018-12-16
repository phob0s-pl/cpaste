# cpaste

cpaste is command line client for pastebin.com

To obtain the tool type in the console:
```bash
go install github.com/phob0s-pl/cpaste/cmd/cpaste
```

Before first usage user must have pro account on pastebin and obtain user key with cpaste tool.
User key allows user to create pastes as logged user. To obtain such a key:
```bash
cpaste session --devkey development_key --user my_mysername --pass password
```

cpaste stores only dev and user keys in file located under **~/.config/cpaste.json**

#### Posting pastes
To post paste from stdin:
```bash
phob0s@pc ~ $ echo "hi I am paste" | cpaste publish -e 1H
https://pastebin.com/id1234
```

User can also post a file:
```bash
phob0s@pc ~ $ cpaste file --path file.go -e 1H
https://pastebin.com/id1235
```

#### Listing pastes
To show own pastes:
```bash
phob0s@pc ~ $ cpaste list
----------------------------------------------------------------------------------
|           Paste URL          |  time left   | scope |  format  |  title        |
----------------------------------------------------------------------------------
https://pastebin.com/id1235     1h0m0s       public   [go]      "file.go"
https://pastebin.com/id1234     1h0m0s       public   [text]    ""
```

#### Deleting pastes
To delete paste:
```bash
phob0s@pc ~ $ cpaste delete --paste-key id1235
Paste Removed
```





