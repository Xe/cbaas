import dice

def handle(event, context):
    return [{
        "to": event["to"],
        "body": [
            {"body": dice.diceWrapper(event["bodyString"].split())},
        ],
    }]
