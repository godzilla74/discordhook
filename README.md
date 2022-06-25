# discordhook
This package provides a super simple interface to send discord messages through webhooks in golang.

## Installation
```
go get github.com/jonathantyar/discordhook
```
## Example
Below is some example that you can use, and try on example.go right away.

### Simple Message
```
	message := discordhook.Message{
		Username: &username,
		Content:  &message,
	}

	err := discordhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}
```
Generate simple message like this.
![image](https://user-images.githubusercontent.com/19704585/175761705-aa3fb66d-1509-4ebb-875f-56b43e3bebdc.png)

### Message with Embeds
```
	var (
		ccy       = "BTC"
		value     = "0.0002"
		inline    = true
		typeName  = "type"
		valueType = "join"
	)

	fields := make([]discordhook.Field, 0)
	fields = append(fields, discordhook.Field{
		Name:   &ccy,
		Value:  &value,
		Inline: &inline,
	})
	fields = append(fields, discordhook.Field{
		Name:   &typeName,
		Value:  &valueType,
		Inline: &inline,
	})
	fields = append(fields, discordhook.Field{
		Name:  &ccy,
		Value: &value,
	})

	embeds := make([]discordhook.Embed, 0)
	embeds = append(embeds, discordhook.Embed{
		Title:       &message,
		Description: &description,
		Fields:      &fields,
	})

	message := discordhook.Message{
		Username: &username,
		Content:  &message,
		Embeds:   &embeds,
	}

	err := discordhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}

```
![image](https://user-images.githubusercontent.com/19704585/175761732-fcdd037a-a560-433a-a316-42f5869f44f7.png)
Generate some embed message like that. Please note that if you want the field to be inline, pass the Field type to true.
